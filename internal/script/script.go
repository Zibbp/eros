package script

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eros/ent"
	entReport "github.com/zibbp/eros/ent/report"
	"github.com/zibbp/eros/ent/script"
	"github.com/zibbp/eros/internal/database"
	"github.com/zibbp/eros/internal/utils"
)

type Service struct {
	db *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{
		db: db,
	}
}

type Script struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Hostname  string    `json:"hostname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaginationResponse struct {
	Pagination utils.Pagination `json:"pagination"`
	Data       []*ent.Script    `json:"data"`
}

func (s *Service) CreateScript(c echo.Context, scriptDto Script) (*ent.Script, error) {

	// check if script already exists
	_, err := s.db.Client.Script.Query().Where(script.NameEQ(scriptDto.Name)).Where(script.HostnameEQ(scriptDto.Hostname)).First(c.Request().Context())
	if err == nil {
		log.Error().Err(err).Msg("script already exists")
		return nil, fmt.Errorf("script already exists")
	}

	script, err := s.db.Client.Script.Create().SetID(uuid.New()).SetName(scriptDto.Name).SetHostname(scriptDto.Hostname).Save(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to create script")
		return nil, fmt.Errorf("failed to create script")
	}

	return script, nil

}

func (s *Service) GetScripts(c echo.Context, limit int, offset int) (PaginationResponse, error) {
	var paginationResponse PaginationResponse

	// query builder
	query := s.db.Client.Script.Query()

	var total int
	scripts, err := query.Order(ent.Desc(script.FieldUpdatedAt)).Limit(limit).Offset(offset).All(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			log.Info().Msg("No scripts found")
			// handle here if no scripts are found
		} else {
			log.Error().Err(err).Msg("Failed to get scripts")
			return paginationResponse, err
		}
	} else {
		// if scripts are found
		// get total count
		total, err = s.db.Client.Script.Query().Count(c.Request().Context())
		if err != nil {
			log.Error().Err(err).Msg("Failed to get scripts count")
			return paginationResponse, err
		}
	}

	// fill reports for each script
	for _, script := range scripts {
		reports, err := script.QueryReports().Order(ent.Desc(entReport.FieldCreatedAt)).Limit(1).All(c.Request().Context())
		if err != nil {
			log.Error().Err(err).Msg("Failed to get script reports")
			return paginationResponse, err
		}
		if len(reports) > 0 {
			script.Edges.Reports = reports
		}
	}

	paginationResponse.Pagination = utils.Pagination{
		Offset: offset,
		Limit:  limit,
		Total:  total,
		Pages:  int(math.Ceil(float64(total) / float64(limit))),
	}

	paginationResponse.Data = scripts

	return paginationResponse, nil
}

func (s *Service) GetScript(c echo.Context, id uuid.UUID) (*ent.Script, error) {
	script, err := s.db.Client.Script.Query().Where(script.IDEQ(id)).First(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			log.Info().Msg("Script not found")
			// handle here if script is not found
		} else {
			log.Error().Err(err).Msg("Failed to get script")
			return nil, err
		}
	}

	return script, nil
}
