package commands

import (
	"encoding/json"
	"github.com/russellcxl/dinner-booker/http_"
	"github.com/russellcxl/dinner-booker/models"
	"io/ioutil"
)

func GetCurrent(bot *http_.DinnerBot) (*models.CurrentMenuResponse, error) {
	res, err := bot.Get("/api/current", nil)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var menu models.CurrentMenuResponse
	if err = json.Unmarshal(b, &menu); err != nil {
		return nil, err
	}
	return &menu, nil
}