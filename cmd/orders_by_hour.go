package cmd

import (
	"fmt"
	"os"
	"session-14/handler"
)

func OrdersByHour(handlerReport handler.ReportHandler) {
	var choice string

	handlerReport.OrdersByHour()

	fmt.Print("\nApakah kamu ingin melanjutkan ke halaman lain? (ya/tidak): ")
	fmt.Scanln(&choice)

	switch choice {
	case "ya":
		ClearScreen()
	case "tidak":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid, kembali ke menu utama.")
		ClearScreen()
		return
	}
}
