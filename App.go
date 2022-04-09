package main

import (
	"Revisi_WBH/Config"
	"Revisi_WBH/Delivery"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jmoiron/sqlx"
)

func main() {
	appConfig := Config.NewConfig()

	Delivery.MainMenu()

	for {
		var confirmation string
		fmt.Scanln(&confirmation)

		switch confirmation {
		case "1":
			Delivery.ShowTable(appConfig.UseCaseManager)

		}
	}

	//Delivery.ShowTable(appConfig.UseCaseManager)
	//Delivery.ShowListOfMenu(appConfig.UseCaseManager)
}
