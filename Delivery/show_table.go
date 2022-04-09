package Delivery

import (
	"Revisi_WBH/manager"
	"Revisi_WBH/util"
	"fmt"
	"strings"
)

func ShowTable(usecase manager.UseCaseManager) {
	fmt.Println(util.TableList)
	fmt.Println(strings.Repeat("*", 50))
	for idx, v := range usecase.CallTableList().SearchTable() {
		if idx != 0 && idx%5 == 0 {
			fmt.Print("\n")
		}
		fmt.Printf(util.TableListFormat, v.IdTable)
	}

}
