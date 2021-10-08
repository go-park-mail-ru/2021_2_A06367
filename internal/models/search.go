package models

type SearchResult struct {
	Profiles []Profile `json:"profiles"`
	Films    []Film    `json:"films"`
}
