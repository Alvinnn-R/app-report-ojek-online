package handler

import (
	"fmt"
	"session-14/service"
)

type ReportHandler struct {
	ServiceReport service.ServiceReportInterface
}

func NewReportHandler(serviceReport service.ServiceReportInterface) ReportHandler {
	return ReportHandler{
		ServiceReport: serviceReport,
	}
}

func (handler *ReportHandler) ReportMonthly(status string) {
	reports, err := handler.ServiceReport.GetReportMonthly(status)
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
		return
	}

	fmt.Println("\n=== Order Per Month ===")
	for _, report := range reports {
		fmt.Printf("Month: %s, Total Order: %d\n", report.Month, report.TotalOrder)
	}
}

func (handler *ReportHandler) TopCustomerPerMonth() {
	reports, err := handler.ServiceReport.GetTopCustomerPerMonth()
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
		return
	}

	fmt.Println("\n=== Top Customer Per Month ===")
	for _, report := range reports {
		fmt.Printf("%s: %s (%d orders)\n", report.Month, report.CustomerName, report.TotalOrder)
	}
}

func (handler *ReportHandler) TopAreaByType(areaType string) {
	reports, err := handler.ServiceReport.GetTopAreaByType(areaType)
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
		return
	}

	var typeLabel string
	if areaType == "pickup" {
		typeLabel = "Pickup"
	} else if areaType == "dropoff" {
		typeLabel = "Dropoff"
	} else {
		typeLabel = "Overall"
	}

	fmt.Printf("\n=== Top Area (%s) ===\n", typeLabel)
	for i, report := range reports {
		fmt.Printf("%d. %s, %s - %d orders\n", i+1, report.AreaName, report.City, report.TotalOrder)
	}
}

func (handler *ReportHandler) OrdersByHour() {
	reports, err := handler.ServiceReport.GetOrdersByHour()
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
		return
	}

	fmt.Println("\n=== Orders By Hour ===")
	fmt.Println("Time Range\t\tTotal Orders\tCategory")
	fmt.Println("================================================")
	for _, report := range reports {
		fmt.Printf("%s\t\t%d\t\t%s\n", report.TimeRange, report.TotalOrder, report.Category)
	}
}
