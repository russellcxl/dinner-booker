package commands

import (
	"encoding/json"
	"fmt"
	"github.com/russellcxl/dinner-booker/http_"
	"github.com/russellcxl/dinner-booker/models"
	"io/ioutil"
	"time"
)

func Order(bot *http_.DinnerBot, menuID int, orderReq models.OrderReq) (*models.OrderResp, error) {
	var orderResp models.OrderResp
	var tryCount int

	// retries up to 30 times; 5 min total
	for {
		res, err := bot.Post(fmt.Sprintf("/api/order/%d", menuID), orderReq)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(b, &orderResp); err != nil {
			return nil, err
		}

		if orderResp.Error == "" {
			break
		}
		if tryCount > 30 {
			break
		}
		tryCount++
		time.Sleep(10 * time.Second)
	}

	return &orderResp, nil
}
