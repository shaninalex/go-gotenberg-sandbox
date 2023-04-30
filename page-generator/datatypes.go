package main

type DataGathering struct {
	CompanyOverview *CompanyOverviewPayload
	Wallstreetbets  []WallstreetbetsItem
}

// WallstreetbetsResponse
type WallstreetbetsResponse struct {
	Data   []WallstreetbetsItem `json:"data"`
	Status uint16               `json:"status"`
}

type WallstreetbetsItem struct {
	Date      string  `json:"Date"`
	Mentions  uint    `json:"Mentions"`
	Rank      uint8   `json:"Rank"`
	Sentiment float32 `json:"Sentiment"`
	Ticker    string  `json:"Ticker"`
}

// COMPANY OVERVIEW
type CompanyOverviewPayload struct {
	Name   string
	Symbol string
}
