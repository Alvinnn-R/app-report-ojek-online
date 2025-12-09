package cmd

import (
	"fmt"
	// "os"
	"session-14/handler"
)

func ReportMonthly(handlerReport handler.ReportHandler) {
	var status, choice string

	fmt.Print("Status : ")
	fmt.Scanln(&status)

	handlerReport.ReportMonthly(status)

	fmt.Print("\nApakah kamu ingin melanjutkan ke halaman lain? (ya/tidak): ")
	fmt.Scanln(&choice)

	switch choice { 
	case "ya":
		ClearScreen()
	case "tidak":
		fmt.Println("Kembali ke menu utama...")
		ClearScreen()
		return
	default:
		fmt.Println("Pilihan tidak valid, kembali ke menu utama.")
		ClearScreen()
		return
	}
}
