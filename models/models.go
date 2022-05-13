package models

import "time"

type Time time.Time

type TimeRange struct {
	Start time.Time `json:"start" url:"start" layout:"2006-01-02"`
	End   time.Time `json:"end" url:"end" layout:"2006-01-02"`
}

type Menu struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Comment     string    `json:"comment"`
	PollStart   time.Time `json:"poll_start"`
	PollEnd     time.Time `json:"poll_end"`
	ServingTime time.Time `json:"serving_time"`
	Active      bool      `json:"active"`
	VenueID     int       `json:"venue_id"`
}

type ListMenuResponse struct {
	Status string `json:"status"`
	Menu   []Menu `json:"menu"`
	Error  string `json:"error"`
}

type CurrentMenuResponse struct {
	Status string `json:"status"`
	Menu   Menu   `json:"menu"`
	Error  string `json:"error"`
}

type Food struct {
	Code        string      `json:"code"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ImageURL    string      `json:"image_url"`
	Ordered     int         `json:"ordered"`
	Quota       interface{} `json:"quota"`
	Disabled    bool        `json:"disabled"`
	Remaining   int         `json:"remaining,omitempty"`
}

type MenuResponse struct {
	Status string `json:"status"`
	Food   []Food `json:"food"`
	Menu   Menu   `json:"menu"`
	Error  string `json:"error"`
}

type ListOrderedResponse struct {
	Status string `json:"status"`
	Food   *Food  `json:"food"`
	Error  string `json:"error"`
}

type GetOrderResponse struct {
	Food  *Food  `json:"food"`
	Error string `json:"error"`
}

type OrderReq struct {
	FoodID int `json:"food_id" url:"food_id"`
}

type OrderResp struct {
	Selected int    `json:"selected"`
	Error    string `json:"error"`
}
