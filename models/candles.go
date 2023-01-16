package models

type Candle struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Picture       string `json:"Picture"`
	Title         string `json:"Title"`
	FullName      string `json:"FullName"`
	Description   string `json:"Description"`
	PriceMassive  string `json:"PriceMassive"`
	TypeFragrance string `json:"TypeFragrance"`
	TopNotes      string `json:"TopNotes"`
	MiddleNotes   string `json:"MiddleNotes"`
	BasicNotes    string `json:"BasicNotes"`
}
