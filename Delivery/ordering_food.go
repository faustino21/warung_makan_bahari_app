package Delivery

import (
	"Revisi_WBH/manager"
	"Revisi_WBH/util"
	"fmt"
)

func OrderingFood(usecase manager.UseCaseManager) {
	var confirmation string
	var customer string
	var table int
	var listMenuChoosen []int
	var menuChoosen int


	for {
		fmt.Println(util.AskName)
		fmt.Scanln(&customer)
		ShowTable(usecase)
		fmt.Println(util.TableChoosen)
		fmt.Scanln(&table)
		fmt.Println(util.ConfirmationForNextOption)
		fmt.Scanln(&confirmation)
		if confirmation != "y" {
			break
		}
		ShowListOfMenu(usecase)
		fmt.Println(util.ChooseMenu)
		for {
			fmt.Scanln(&menuChoosen)
			listMenuChoosen = append(listMenuChoosen, menuChoosen)
			fmt.Print(util.ChooseMenu2)
			fmt.Scanln(&confirmation)
			if confirmation == "y" {
				continue
			} else {
				fmt.Println(util.ConfirmationForNextOption)
				fmt.Scanln(&confirmation)
				if confirmation == "y" {
					break
				} else {
					listMenuChoosen = nil
					continue
				}
			}

		}
		fmt.Print(util.ConfirmationForNextOption)
		fmt.Scanln(&confirmation)
		if confirmation == "y" {
			var orderId int
			payment := util.PaymentStatus
			usecase.CallRegisterCustomer().RegisterCustomer()
			for _,v := range listMenuChoosen {
				orderId = usecase.CallOrderFood().CountingOrder()
				usecase.CallOrderFood().OrderingFood(*orderId+1, *customer, *)

			}
		}
	}

	MainMenu()

}
