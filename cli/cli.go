package cli

import (
	"github.com/manifoldco/promptui"
	"github.com/russellcxl/dinner-booker/commands"
	"github.com/russellcxl/dinner-booker/http_"
	"github.com/russellcxl/dinner-booker/models"
	"github.com/russellcxl/dinner-booker/timer"

	"log"
)

func Execute(bot *http_.DinnerBot) {
	currentMenu, err := commands.GetCurrent(bot)
	if err != nil {
		log.Fatalf("failed to get /api/current: %s\n", err.Error())
	}

	foodMenu, err := commands.GetMenu(bot, currentMenu.Menu.ID)
	if err != nil {
		log.Fatalf("failed to get /api/menu/{menuID}: %s\n", err.Error())
	}

	var foodOptions []string
	foodMap := make(map[string]int, len(foodMenu.Food))

	for _, f := range foodMenu.Food {
		foodMap[f.Name] = f.ID
		foodOptions = append(foodOptions, f.Name)
	}

	prompt := promptui.Select{
		Label: "Select action",
		Items: foodOptions,
	}

	_, res, err := prompt.Run()
	if err != nil {
		log.Fatalf("prompt failed: %v\n", err)
		return
	}

	foodID, ok := foodMap[res]
	if !ok {
		log.Fatalf("failed to find food ID\n")
	}

	orderReq := models.OrderReq{FoodID: foodID}

	// blocks until 12:30pm
	timer.SetRudimentaryTimer()

	orderRes, err := commands.Order(bot, currentMenu.Menu.ID, orderReq)
	if err != nil {
		log.Fatalf("failed to make order: %s\n", err.Error())
	}

	if orderRes.Error != "" {
		log.Printf("failed to make order: %s\n", orderRes.Error)
	} else {
		log.Printf("order made for: %s (%d)\n", res, foodID)
	}
}
