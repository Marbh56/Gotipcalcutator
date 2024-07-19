package main

import (
	"fmt"
	"strconv"
)

func main() {
	var billAmount string
	var tip string

	fmt.Print("Bill Amount: ")
	_, err := fmt.Scanln(&billAmount)
	if err != nil {
		fmt.Println("Error reading bill", err)
		return
	}

	fmt.Print("Tip: ")
	_, err = fmt.Scanln(&tip)
	if err != nil {
		fmt.Println("Error reading tip", err)
		return
	}

	tipFloat, err := strconv.ParseFloat(tip, 64)
	if err != nil {
		fmt.Println("error converting tip from string")
		return
	}
	tipPercent := (tipFloat / 100)

	billFloat, err := strconv.ParseFloat(billAmount, 64)
	if err != nil {
		fmt.Println("error converting bill to float")
	}

	finalTip := billFloat * tipPercent
	totalAmount := billFloat + finalTip

	fmt.Printf("Tip: $%.2f\n", finalTip)
	fmt.Printf("Amount: $%.2f\n", totalAmount)

}
