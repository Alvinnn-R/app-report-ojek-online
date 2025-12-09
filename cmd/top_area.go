package cmd

import (
	"fmt"
	"os"
	"session-14/handler"
)

func TopArea(handlerReport handler.ReportHandler) {
	var areaType, choice string

	fmt.Println("\nPilih tipe area:")
	fmt.Println("1. Pickup Area (Penjemputan)")
	fmt.Println("2. Dropoff Area (Tujuan)")
	fmt.Println("3. Overall (Gabungan)")
	fmt.Print("Pilihan (1/2/3): ")
	fmt.Scanln(&areaType)

	var typeInput string
	switch areaType {
	case "1":
		typeInput = "pickup"
	case "2":
		typeInput = "dropoff"
	case "3":
		typeInput = "overall"
	default:
		fmt.Println("Pilihan tidak valid, menggunakan overall.")
		typeInput = "overall"
	}

	handlerReport.TopAreaByType(typeInput)

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
