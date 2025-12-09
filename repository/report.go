package repository

import (
	"context"
	"session-14/database"
	"session-14/model"
)

type RepositoryReportInterface interface {
	GetReportMonthly(status string) ([]*model.Report, error)
	GetTopCustomerPerMonth() ([]*model.CustomerReport, error)
	GetTopAreaByType(areaType string) ([]*model.AreaReport, error)
	GetOrdersByHour() ([]*model.HourReport, error)
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

func (repo *RepositoryReport) GetTopCustomerPerMonth() ([]*model.CustomerReport, error) {
	query := `WITH monthly_customer_orders AS (
		SELECT 
			TO_CHAR(o.requested_at, 'Month YYYY') AS periode,
			u.name AS customer_name,
			COUNT(*) AS total_order
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		JOIN users u ON c.user_id = u.id
		GROUP BY TO_CHAR(o.requested_at, 'Month YYYY'), u.name
	),
	ranked_customers AS (
		SELECT 
			periode,
			customer_name,
			total_order,
			RANK() OVER (PARTITION BY periode ORDER BY total_order DESC) AS ranking
		FROM monthly_customer_orders
	)
	SELECT 
		periode,
		customer_name,
		total_order
	FROM ranked_customers
	WHERE ranking = 1
	ORDER BY periode DESC`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customerReports []*model.CustomerReport

	for rows.Next() {
		var r model.CustomerReport
		err := rows.Scan(&r.Month, &r.CustomerName, &r.TotalOrder)
		if err != nil {
			return nil, err
		}
		customerReports = append(customerReports, &r)
	}

	return customerReports, nil
}

func (repo *RepositoryReport) GetTopAreaByType(areaType string) ([]*model.AreaReport, error) {
	var query string

	if areaType == "pickup" {
		query = `SELECT 
			a.name AS area_name,
			a.city,
			COUNT(*) AS total_order
		FROM orders o
		JOIN areas a ON o.pickup_area_id = a.id
		GROUP BY a.id, a.name, a.city
		ORDER BY total_order DESC
		LIMIT 10`
	} else if areaType == "dropoff" {
		query = `SELECT 
			a.name AS area_name,
			a.city,
			COUNT(*) AS total_order
		FROM orders o
		JOIN areas a ON o.dropoff_area_id = a.id
		GROUP BY a.id, a.name, a.city
		ORDER BY total_order DESC
		LIMIT 10`
	} else {
		query = `SELECT 
			a.name AS area_name,
			a.city,
			COALESCE(pickup.total, 0) + COALESCE(dropoff.total, 0) AS total_aktivitas
		FROM areas a
		LEFT JOIN (
			SELECT pickup_area_id, COUNT(*) AS total
			FROM orders
			GROUP BY pickup_area_id
		) pickup ON a.id = pickup.pickup_area_id
		LEFT JOIN (
			SELECT dropoff_area_id, COUNT(*) AS total
			FROM orders
			GROUP BY dropoff_area_id
		) dropoff ON a.id = dropoff.dropoff_area_id
		WHERE COALESCE(pickup.total, 0) + COALESCE(dropoff.total, 0) > 0
		ORDER BY total_aktivitas DESC
		LIMIT 10`
	}

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areaReports []*model.AreaReport

	for rows.Next() {
		var r model.AreaReport
		err := rows.Scan(&r.AreaName, &r.City, &r.TotalOrder)
		if err != nil {
			return nil, err
		}
		areaReports = append(areaReports, &r)
	}

	return areaReports, nil
}

func (repo *RepositoryReport) GetOrdersByHour() ([]*model.HourReport, error) {
	query := `WITH hourly_orders AS (
		SELECT 
			EXTRACT(HOUR FROM requested_at)::INT AS jam,
			COUNT(*) AS total_order
		FROM orders
		GROUP BY EXTRACT(HOUR FROM requested_at)
	),
	stats AS (
		SELECT 
			AVG(total_order) AS rata_rata
		FROM hourly_orders
	)
	SELECT 
		ho.jam,
		LPAD(ho.jam::TEXT, 2, '0') || ':00 - ' || LPAD(ho.jam::TEXT, 2, '0') || ':59' AS rentang_waktu,
		ho.total_order,
		CASE 
			WHEN ho.total_order >= s.rata_rata THEN 'RAMAI'
			ELSE 'SEPI'
		END AS kategori
	FROM hourly_orders ho
	CROSS JOIN stats s
	ORDER BY ho.jam`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hourReports []*model.HourReport

	for rows.Next() {
		var r model.HourReport
		err := rows.Scan(&r.Hour, &r.TimeRange, &r.TotalOrder, &r.Category)
		if err != nil {
			return nil, err
		}
		hourReports = append(hourReports, &r)
	}

	return hourReports, nil
}
