package service

import (
	"session-14/model"
	"session-14/repository"
)

type ServiceReportInterface interface {
	GetReportMonthly(status string) ([]*model.Report, error)
	GetTopCustomerPerMonth() ([]*model.CustomerReport, error)
	GetTopAreaByType(areaType string) ([]*model.AreaReport, error)
	GetOrdersByHour() ([]*model.HourReport, error)
}

type ServiceReport struct {
	RepoReport repository.RepositoryReportInterface
}

func NewServiceReport(repoReport repository.RepositoryReportInterface) ServiceReport {
	return ServiceReport{
		RepoReport: repoReport,
	}
}

func (serviceReport *ServiceReport) GetReportMonthly(status string) ([]*model.Report, error) {
	return serviceReport.RepoReport.GetReportMonthly(status)
}

func (serviceReport *ServiceReport) GetTopCustomerPerMonth() ([]*model.CustomerReport, error) {
	return serviceReport.RepoReport.GetTopCustomerPerMonth()
}

func (serviceReport *ServiceReport) GetTopAreaByType(areaType string) ([]*model.AreaReport, error) {
	return serviceReport.RepoReport.GetTopAreaByType(areaType)
}

func (serviceReport *ServiceReport) GetOrdersByHour() ([]*model.HourReport, error) {
	return serviceReport.RepoReport.GetOrdersByHour()
}
