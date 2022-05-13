package commands

import (
	"encoding/json"
	"fmt"
	"github.com/russellcxl/dinner-booker/http_"
	"github.com/russellcxl/dinner-booker/models"
	"io/ioutil"
)

func GetMenu(bot *http_.DinnerBot, menuID int) (*models.MenuResponse, error){
	res, err := bot.Get(fmt.Sprintf("/api/menu/%d", menuID), nil)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var menuResp models.MenuResponse
	if err = json.Unmarshal(b, &menuResp); err != nil {
		return nil, err
	}
	return &menuResp, err
}