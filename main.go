package main

import (
	"PersonalIncomeTaxCalculation/service"
	"PersonalIncomeTaxCalculation/utils"
	"fmt"
	"log"
	"os"
)

// Mac下编译Linux, Windows平台的64位可执行程序：
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

func main() {
	fmt.Println("请选择计税方式：")
	fmt.Println("[1]. 年终奖单独计税")
	fmt.Println("[2]. 年终奖综合计税")
	inputHandle()
}

func inputHandle() {
	fmt.Println("请输入对应的序号, 按回车键确认(输入exit退出): ")
	inputType, err := utils.GetCommandString()
	if err != nil {
		log.Fatal("输入错误")
	}
	fmt.Println("请输入月平均工资")
	salary, err := utils.GetCommandFloat()
	fmt.Println("请输入年终奖")
	yearEndAwards, err := utils.GetCommandFloat()
	fmt.Println("五险一金比例，默认18.5%")
	fmt.Println("专项扣除，默认1500")
	if inputType == "1" {
		service.SeparateTax(salary, yearEndAwards)
	}
	if inputType == "2" {
		service.ComprehensiveTax(salary, yearEndAwards)
	}
	fmt.Println("输入exit退出, 按回车键继续计算")
	inputType, err = utils.GetCommandString()
	if inputType == "exit" {
		os.Exit(0)
	} else {
		inputHandle()
	}
}
