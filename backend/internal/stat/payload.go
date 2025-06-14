package stat

import "time"

type StatGetByDate struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	By   string    `json:"by"`
}

type GetGroupStatParamsRequest struct {
	From string `schema:"from"`
	To   string `schema:"to"`
	By   string `schema:"by"`
}
type GetGroupStatResponse struct {
	Period string `json:"period"`
	Count  string `json:"count"`
}
