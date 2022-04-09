package Delivery

import (
	"Revisi_WBH/manager"
	"Revisi_WBH/util"
	"fmt"
	"strings"
)

func ShowListOfMenu(useCase manager.UseCaseManager) {
	fmt.Println(util.MenuList)
	fmt.Println(strings.Repeat("*", 50))
	fmt.Printf(util.MenuListFormat, util.No, util.NamaMakanan, util.HargaMakanan)
	for idx, v := range useCase.CallMenuList().ShowMenuList() {
		fmt.Printf(util.MenuListDbFormat, idx+1, v.ProductName, v.ProductPrice)
	}
}
