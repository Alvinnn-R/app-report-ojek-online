package repository

import (
	"context"
	"session-14/database"
	"session-14/model"
)

type RepositoryReportInterface interface {
	GetReportMonthly(status string) ([]*model.Report, error)
}

// type RepositoryReport struct {
// 	DB *pgx.Conn
// }

type RepositoryReport struct {
	DB database.PgxIface
}

func NewRepoReport(db database.PgxIface) RepositoryReport {
	return RepositoryReport{
		DB: db,
	}
}

func (repo *RepositoryReport) GetReportMonthly(status string) ([]*model.Report, error) {
	query := `SELECT 
    TO_CHAR(DATE_TRUNC('month', requested_at), 'Month') AS month,
    COUNT(*) AS total_order
FROM orders
WHERE status = $1
GROUP BY month
ORDER BY MIN(DATE_TRUNC('month', requested_at))`

	// now := time.Now()
	rows, err := repo.DB.Query(context.Background(), query, status)

	if err != nil {
		return nil, err
	}

	var reportsMonthly []*model.Report
	var r model.Report

	for rows.Next() {
		err := rows.Scan(&r.Month, &r.TotalOrder)
		if err != nil {
			return nil, err
		}

		reportsMonthly = append(reportsMonthly, &r)
	}

	return reportsMonthly, nil
}
