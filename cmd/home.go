package cmd

import (
	"fmt"
	"os"
	"session-14/handler"
)

func HomePage(handlerReport handler.ReportHandler) {
	for {
		fmt.Println("\n=== HOME MENU ===")
		fmt.Println("1. View Report Monthly")
		fmt.Println("2. View Top Customer Per Month")
		fmt.Println("3. View Top Area")
		fmt.Println("4. View Orders By Hour")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			ClearScreen()
			ReportMonthly(handlerReport)
		case 2:
			ClearScreen()
			TopCustomer(handlerReport)
		case 3:
			ClearScreen()
			TopArea(handlerReport)
		case 4:
			ClearScreen()
			OrdersByHour(handlerReport)
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose between 1 - 5.")
		}
	}

}
