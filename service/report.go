package service

import (
	"session-14/model"
	"session-14/repository"
)

type ServiceReportInterface interface {
	GetReportMonthly(status string) ([]*model.Report, error)
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
