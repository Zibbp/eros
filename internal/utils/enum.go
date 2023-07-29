package utils

type ReportStatus string

const (
	ReportStatusSuccess ReportStatus = "success"
	ReportStatusFailed  ReportStatus = "failed"
)

func (ReportStatus) Values() (kinds []string) {
	for _, s := range []ReportStatus{ReportStatusSuccess, ReportStatusFailed} {
		kinds = append(kinds, string(s))
	}
	return
}
