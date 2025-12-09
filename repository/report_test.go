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