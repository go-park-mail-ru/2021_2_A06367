package models

type SearchResult struct {
	Profiles []Profile  `json:"profiles" example:"{'Достать ножи', 'Казино рояль', 'Девушка с татуировкой дракона'}"`
	Films   []Film    `json:"films"    example:"{'Достать ножи', 'Казино рояль', 'Девушка с татуировкой дракона'}"`
}
