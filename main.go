package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

type ColumnMapping struct {
	ExcelColumn string `json:"excel_column"`
	DBColumn    string `json:"db_column"`
}

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Table    string `json:"table"`
	} `json:"database"`
	SheetName      string          `json:"sheet_name"`
	ColumnMappings []ColumnMapping `json:"column_mappings"`
	StateUUID      string          `json:"state_uuid"`
}

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Open database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	defer db.Close()

	// Process Excel file
	xlsx, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		log.Fatal("Failed to open Excel file:", err)
	}

	rows, err := xlsx.GetRows(config.SheetName)
	if err != nil {
		log.Fatal("Failed to read rows:", err)
	}

	// Skip header row
	if len(rows) > 0 {
		rows = rows[1:]
	}

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}

	// Defer rollback in case anything fails
	// If commit is successful, this becomes a no-op
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf("Transaction rolled back due to panic: %v", r)
		}
	}()

	// For master_cities, we'll use hardcoded columns instead of config mappings
	// The column mappings are only used to locate the name in the Excel file

	// For master_cities, we need to hardcode some columns: uuid, state_uuid, created_at, updated_at
	// Only name comes from the Excel file
	allColumns := []string{"uuid", "state_uuid", "name", "created_at", "updated_at"}
	// Create placeholders with commas (?,?,?,?,?) instead of just ?????
	placeholders := make([]string, len(allColumns))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		config.Database.Table,
		strings.Join(allColumns, ","),
		strings.Join(placeholders, ",")))
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to prepare statement:", err)
	}

	defer stmt.Close()

	// Insert data
	insertErrors := false
	for rowIndex, row := range rows {
		// Skip header row
		if rowIndex == 0 {
			continue
		}

		// Initialize with 5 values for all columns: uuid, state_uuid, name, created_at, updated_at
		values := make([]interface{}, 5)

		// Generate UUID for the city
		values[0] = uuid.NewString()

		// Hardcoded state_uuid - replace with the appropriate value for your data
		values[1] = config.StateUUID // Replace this with actual state UUID

		// Set the city name from Excel
		values[2] = row[2]

		// Set created_at and updated_at timestamps
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		values[3] = currentTime // created_at
		values[4] = currentTime // updated_at

		fmt.Println("Inserting row:", values)

		_, err := stmt.Exec(values...)
		if err != nil {
			log.Printf("Failed to insert row %d: %v\n", rowIndex+1, err)
			insertErrors = true
			// Uncomment the following line if you want to rollback on any error
			// tx.Rollback()
			continue
		}
	}

	// Commit the transaction if no errors occurred
	if !insertErrors {
		if err := tx.Commit(); err != nil {
			tx.Rollback()
			log.Fatal("Failed to commit transaction:", err)
		}
		log.Println("Transaction committed successfully!")
	} else {
		tx.Rollback()
		log.Println("Transaction rolled back due to insert errors")
	}

	log.Println("Data import completed successfully!")
}

func loadConfig() (*Config, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	config := &Config{}
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
