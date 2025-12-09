package repository

import (
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func TestTotalOrderPerMount(t *testing.T) {
	// Mock DB and sqlmock
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"month", "total_order"}).
		AddRow("January", 10).
		AddRow("December", 20)

	mock.ExpectQuery("SELECT (.+) FROM orders WHERE status = \\$1 GROUP BY month ORDER BY MIN").
		WithArgs("completed").
		WillReturnRows(rows)

	repo := NewRepoReport(mock)

	result, err := repo.GetReportMonthly("completed")

	assert.NoError(t, err)
	assert.Len(t, result, 2)

	assert.Equal(t, "January", result[0].Month)
	assert.Equal(t, 10, result[0].TotalOrder)

	assert.Equal(t, "December", result[1].Month)
	assert.Equal(t, 20, result[1].TotalOrder)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetTopCustomerPerMonth(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"periode", "customer_name", "total_order"}).
		AddRow("January 2025", "Alvin Saputra", 5).
		AddRow("December 2024", "Rizky Pratama", 8)

	mock.ExpectQuery("WITH monthly_customer_orders AS").
		WillReturnRows(rows)

	repo := NewRepoReport(mock)
	result, err := repo.GetTopCustomerPerMonth()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Alvin Saputra", result[0].CustomerName)
	assert.Equal(t, 5, result[0].TotalOrder)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetTopAreaByTypePickup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"area_name", "city", "total_order"}).
		AddRow("Tegalsari", "Surabaya", 15).
		AddRow("Wonokromo", "Surabaya", 12)

	mock.ExpectQuery("SELECT (.+) FROM orders o JOIN areas a ON o.pickup_area_id").
		WillReturnRows(rows)

	repo := NewRepoReport(mock)
	result, err := repo.GetTopAreaByType("pickup")

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Tegalsari", result[0].AreaName)
	assert.Equal(t, "Surabaya", result[0].City)
	assert.Equal(t, 15, result[0].TotalOrder)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetOrdersByHour(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"jam", "rentang_waktu", "total_order", "kategori"}).
		AddRow(8, "08:00 - 08:59", 25, "RAMAI").
		AddRow(14, "14:00 - 14:59", 15, "SEPI")

	mock.ExpectQuery("WITH hourly_orders AS").
		WillReturnRows(rows)

	repo := NewRepoReport(mock)
	result, err := repo.GetOrdersByHour()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 8, result[0].Hour)
	assert.Equal(t, "08:00 - 08:59", result[0].TimeRange)
	assert.Equal(t, "RAMAI", result[0].Category)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
