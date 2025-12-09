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
	if  err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("Order Per Month:")
	for _, report := range reports {
		fmt.Printf("Month: %s, Total Order: %d\n", report.Month, report.TotalOrder)
	}
}