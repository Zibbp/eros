package report

import (
	"fmt"
	"io/ioutil"
	"math"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eros/ent"
	"github.com/zibbp/eros/ent/report"
	entReport "github.com/zibbp/eros/ent/report"
	entScript "github.com/zibbp/eros/ent/script"
	"github.com/zibbp/eros/internal/database"
	"github.com/zibbp/eros/internal/s3"
	"github.com/zibbp/eros/internal/script"
	"github.com/zibbp/eros/internal/utils"
)

type Service struct {
	db            *database.Database
	s3            *s3.S3Client
	scriptService *script.Service
}

func NewService(db *database.Database, s3 *s3.S3Client, script *script.Service) *Service {
	return &Service{
		db:            db,
		s3:            s3,
		scriptService: script,
	}
}

type Report struct {
	ID        uuid.UUID          `json:"id"`
	Name      string             `json:"name"`
	Status    utils.ReportStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Script    script.Script      `json:"script"`
	File      *multipart.FileHeader
	FileSrc   multipart.File
}

type CreateReportResponse struct {
	ID     uuid.UUID          `json:"id"`
	Name   string             `json:"name"`
	Status utils.ReportStatus `json:"status"`
	S3File string             `json:"s3_file"`
}

type ReportResponse struct {
	*ent.Report
	Log string `json:"log"`
}

type ScriptReportsPaginationResponse struct {
	Pagination utils.Pagination `json:"pagination"`
	Data       []*ent.Report    `json:"data"`
}

func (s *Service) CreateReport(c echo.Context, reportDto Report) (*CreateReportResponse, error) {

	// get script if exists else create it
	script, err := s.db.Client.Script.Query().Where(entScript.NameEQ(reportDto.Script.Name)).Where(entScript.HostnameEQ(reportDto.Script.Hostname)).First(c.Request().Context())
	if err != nil {
		log.Info().Msgf("script %s for %s does not exist, creating it", reportDto.Script.Name, reportDto.Script.Hostname)
		script, err = s.scriptService.CreateScript(c, reportDto.Script)
		if err != nil {
			log.Error().Err(err).Msg("failed to create script")
			return nil, fmt.Errorf("failed to create script")
		}
	}

	reportUUID := uuid.New()

	// upload file to s3
	// get file ext
	ext := filepath.Ext(reportDto.File.Filename)
	info, err := s.s3.PutObject(reportDto.FileSrc, reportDto.File.Size, fmt.Sprintf("%s%s", reportUUID.String(), ext), reportDto.File.Header.Get("Content-Type"))
	if err != nil {
		log.Error().Err(err).Msg("failed to upload file to s3")
		return nil, fmt.Errorf("failed to upload file to s3")
	}

	// create report
	report, err := s.db.Client.Report.Create().
		SetID(reportUUID).
		SetName(reportDto.Name).
		SetStatus(reportDto.Status).
		SetScript(script).
		SetS3File(info.Key).
		Save(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to create report")
		return nil, fmt.Errorf("failed to create report")
	}

	// update script last run
	_, err = s.db.Client.Script.UpdateOneID(script.ID).SetLastRun(time.Now()).Save(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to update script last run")
		return nil, fmt.Errorf("failed to update script last run")
	}

	// return report
	reportResponse := &CreateReportResponse{
		ID:     report.ID,
		Name:   report.Name,
		Status: report.Status,
		S3File: info.Key,
	}

	return reportResponse, nil
}

func (s *Service) GetReport(c echo.Context, id uuid.UUID) (*ReportResponse, error) {
	report, err := s.db.Client.Report.Query().Where(entReport.IDEQ(id)).WithScript().First(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get report")
		return nil, fmt.Errorf("failed to get report")
	}

	var file string

	if report.S3File != nil {
		// get log from s3
		object, err := s.s3.GetObject(*report.S3File)
		if err != nil {
			log.Error().Err(err).Msg("failed to get log from s3")
			return nil, fmt.Errorf("failed to get log from s3")
		}

		defer object.Close()

		bytes, err := ioutil.ReadAll(object)
		if err != nil {
			log.Error().Err(err).Msg("failed to read log from s3")
			return nil, fmt.Errorf("failed to read log from s3")
		}

		file = string(bytes)

	}

	reportResponse := &ReportResponse{
		Report: report,
		Log:    file,
	}

	return reportResponse, nil
}

func (s *Service) GetScriptReports(c echo.Context, scriptID uuid.UUID, limit int, offset int) (ScriptReportsPaginationResponse, error) {
	var paginationResponse ScriptReportsPaginationResponse

	_, err := s.db.Client.Script.Query().Where(entScript.IDEQ(scriptID)).First(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get script")
		return paginationResponse, fmt.Errorf("failed to get script")
	}

	reports, err := s.db.Client.Report.Query().Where(entReport.HasScriptWith(entScript.IDEQ(scriptID))).Limit(limit).Offset(offset).Order(ent.Desc(report.FieldCreatedAt)).All(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get reports")
		return paginationResponse, fmt.Errorf("failed to get reports")
	}

	count, err := s.db.Client.Report.Query().Where(entReport.HasScriptWith(entScript.IDEQ(scriptID))).Count(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get reports count")
		return paginationResponse, fmt.Errorf("failed to get reports count")
	}

	paginationResponse = ScriptReportsPaginationResponse{
		Pagination: utils.Pagination{
			Total:  count,
			Limit:  limit,
			Offset: offset,
			Pages:  int(math.Ceil(float64(count) / float64(limit))),
		},
		Data: reports,
	}

	return paginationResponse, nil
}
