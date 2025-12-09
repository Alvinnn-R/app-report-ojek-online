# App Report Ojek Online

CLI-based reporting application for ride-hailing (ojek online) service with PostgreSQL integration and interactive menu system.

## 📋 Features

1. **View Report Monthly** - Display total orders per month filtered by status
2. **View Top Customer Per Month** - Show customers with the most orders each month
3. **View Top Area** - Display areas with the most orders (pickup/dropoff/overall)
4. **View Orders By Hour** - Analyze peak and off-peak hours for orders

## 🏗️ Project Structure

```
app_report_ojek_online/
├── cmd/
│   ├── home.go              # Main menu
│   ├── report_monthly.go    # Monthly report feature
│   ├── top_customer.go      # Top customer feature
│   ├── top_area.go          # Top area feature
│   ├── orders_by_hour.go    # Orders by hour feature
│   └── clear_screen.go      # Screen utility
├── database/
│   └── db.go                # Database connection
├── handler/
│   └── report.go            # Handler layer
├── model/
│   └── report.go            # Data models
├── repository/
│   ├── report.go            # Repository layer
│   └── report_test.go       # Unit tests
├── service/
│   └── report.go            # Service layer
└── main.go                  # Entry point
```

## 🚀 Installation

### Prerequisites

- Go 1.21+
- PostgreSQL 13+
- Database `db_aplikasi_ojek_online` (from previous project)

### Setup

1. Clone repository

```bash
git clone <repo-url>
cd app_report_ojek_online
```

2. Install dependencies

```bash
go mod tidy
```

3. Configure database connection in `database/db.go`

```go
connStr := "user=postgres password=root dbname=db_aplikasi_ojek_online sslmode=disable host=localhost"
```

4. Run application

```bash
go run .
```

## 🧪 Running Tests

```bash
go test ./repository -v
```

## 📊 Features Detail

### 1. Report Monthly

Filter orders by status (completed/canceled/started) and group by month.

**Query:** Total orders per month with status filter

### 2. Top Customer Per Month

Find customer with most orders in each month using window function RANK().

**Query:** Customer ranking by total orders per month

### 3. Top Area

Analyze most popular areas:

- **Pickup**: Most frequent pickup locations
- **Dropoff**: Most frequent destination locations
- **Overall**: Combined pickup + dropoff activity

**Query:** Top 10 areas by order count

### 4. Orders By Hour

Categorize hours as RAMAI (busy) or SEPI (quiet) based on average.

**Query:** Hourly order distribution with category

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **Driver**: pgx/v5
- **Testing**: testify, pgxmock

## 👨‍💻 Author

**Alvin**  
GitHub: [@Alvinnn-R](https://github.com/Alvinnn-R)

---

**Challenge Project** - Golang Intermediate Daytime Class Bootcamp Lumoshive Academy
