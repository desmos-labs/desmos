package v0150

type GenesisState struct {
	Reports []Report `json:"reports"`
}

func (state *GenesisState) FindReportsForPostWithID(id string) []Report {
	var reports []Report
	for _, report := range state.Reports {
		if report.PostID == id {
			reports = append(reports, report)
		}
	}
	return reports
}

type Report struct {
	PostID  string `json:"post_id"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	User    string `json:"user,omitempty"`
}
