package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "master.shopee_lcs.mysql.cloud.test.shopee.io"
	dbPort = "6606"
	dbUser = "sz_sc_test2"
	dbPass = "6vJaoy5HHiPP0ms_sUnO"
	dbName = "entry_task"

	// Batch size for database inserts
	batchSize = 100

	// Start from row number (set to 0 to start from beginning, or set to specific row to resume)
	// Row 1 is the header, so startRow = 2 means skip header and start from first data row
	// startRow = 835482 means skip until row 835482
	startRow = 1467311
)

func main() {
	csvPath := "/Users/IDSP34746/Documents/2025-11-21-16-23-05-b3331af2-50b7-4c89-b08c-a9bb61243268.csv"

	// Connect to database
	db, err := connectDB()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	// Process CSV row by row to avoid loading entire file into memory
	err = processCSVStreaming(csvPath, db)
	if err != nil {
		fmt.Printf("Error processing CSV: %v\n", err)
		return
	}
}

// processCSVStreaming processes CSV file row by row to minimize memory usage
func processCSVStreaming(csvPath string, db *sql.DB) error {
	// Open the CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Batch for inserts
	batch := make([]logData, 0, batchSize)
	rowCount := 0
	insertCount := 0

	// Skip to startRow if specified
	if startRow > 1 {
		fmt.Printf("Skipping to row %d...\n", startRow)
		for rowCount < startRow {
			_, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					return fmt.Errorf("reached end of file before reaching row %d", startRow)
				}
				return fmt.Errorf("failed to read CSV row while skipping: %w", err)
			}
			rowCount++
			if rowCount%100000 == 0 {
				fmt.Printf("Skipped %d rows...\n", rowCount)
			}
		}
		fmt.Printf("Reached row %d, starting processing...\n", rowCount)
	}

	// Read and process row by row
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read CSV row: %w", err)
		}

		rowCount++

		// Skip header row (only if we started from beginning)
		if rowCount == 1 && startRow <= 1 {
			continue
		}

		// CSV format: @az,@id,@message,@sequence,@timestamp
		if len(record) < 3 {
			continue
		}

		message := record[2]

		// Extract request payload from message
		requestPayload := extractRequestPayload(message)
		if requestPayload == "" {
			continue
		}

		// Extract resource_ids from the payload
		resourceIDs := extractResourceIDs(requestPayload)
		// Extract request ID from message
		requestID := extractRequestID(message)
		// Extract tracking number from payload
		trackingNumber := extractTrackingNumber(requestPayload)

		// Parse timestamp from CSV (last column, index 4)
		var timestamp int64
		if len(record) >= 5 {
			if ts, err := strconv.ParseInt(record[4], 10, 64); err == nil {
				timestamp = ts
			}
		}

		// Add to batch
		batch = append(batch, logData{
			id:             0,
			payload:        requestPayload,
			rawMessage:     message,
			requestID:      requestID,
			timestamps:     timestamp,
			resourceID:     resourceIDs,
			trackingNumber: trackingNumber,
		})

		// Insert batch when it reaches batchSize
		if len(batch) >= batchSize {
			if err := insertBatch(db, batch); err != nil {
				return fmt.Errorf("failed to insert batch: %w", err)
			}
			insertCount += len(batch)
			fmt.Printf("Processed %d rows, inserted %d records\n", rowCount, insertCount)
			// Clear batch to free memory
			batch = batch[:0]
		}
	}

	// Insert remaining records
	if len(batch) > 0 {
		if err := insertBatch(db, batch); err != nil {
			return fmt.Errorf("failed to insert final batch: %w", err)
		}
		insertCount += len(batch)
	}

	fmt.Printf("Processing complete. Total rows: %d, Total inserted: %d\n", rowCount, insertCount)
	return nil
}

// logData represents a single log entry
type logData struct {
	id             int
	payload        string
	rawMessage     string
	requestID      string
	timestamps     int64
	resourceID     string
	trackingNumber string
}

// extractRequestPayload extracts the request payload from a message string
// It looks for content after "request body:" or "http raw request body:" (case insensitive)
func extractRequestPayload(message string) string {
	// Convert to lowercase for case-insensitive search
	lowerMessage := strings.ToLower(message)

	// Look for various markers that indicate request body
	markers := []string{
		"http raw request body:",
		"http raw request body :",
		"request body:",
		"request body :",
		"requestbody:",
		"requestbody :",
	}

	var startIdx int = -1
	for _, marker := range markers {
		idx := strings.Index(lowerMessage, strings.ToLower(marker))
		if idx != -1 {
			startIdx = idx + len(marker)
			break
		}
	}

	if startIdx == -1 {
		// If no marker found, return empty string
		return ""
	}

	// Extract the payload (everything after the marker)
	payload := message[startIdx:]

	// Trim whitespace
	payload = strings.TrimSpace(payload)

	return payload
}

// extractResourceIDs extracts resource_ids from a JSON payload string
// Returns empty string "" if not found
func extractResourceIDs(payload string) string {
	// Search for "resource_ids": in the payload (handle both escaped and unescaped quotes)
	markers := []string{
		`"resource_ids":`,
		`""resource_ids"":`,
		`"resource_ids" :`,
		`""resource_ids"" :`,
	}
	var startIdx int = -1
	for _, marker := range markers {
		idx := strings.Index(payload, marker)
		if idx != -1 {
			startIdx = idx + len(marker)
			break
		}
	}

	if startIdx == -1 {
		return ""
	}
	// Skip whitespace after the colon
	for startIdx < len(payload) && (payload[startIdx] == ' ' || payload[startIdx] == '\t') {
		startIdx++
	}
	if startIdx >= len(payload) {
		return ""
	}
	// Find the start of the array (should be '[')
	if payload[startIdx] != '[' {
		return ""
	}
	// Find the matching closing bracket ']'
	depth := 0
	endIdx := startIdx
	for endIdx < len(payload) {
		switch payload[endIdx] {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				// Extract the array including brackets
				arrayStr := payload[startIdx : endIdx+1]
				// Remove any pipes if present
				arrayStr = strings.ReplaceAll(arrayStr, "|", "")
				return arrayStr
			}
		}
		endIdx++
	}
	return ""
}

// extractRequestID extracts the request ID after "HandleStandardWebhookTracking" in the message
// Returns empty string "" if not found
func extractRequestID(message string) string {
	// Find "HandleStandardWebhookTracking" in the message
	idx := strings.Index(message, "HandleStandardWebhookTracking")
	if idx == -1 {
		return ""
	}
	// Start after "HandleStandardWebhookTracking"
	start := idx + len("HandleStandardWebhookTracking")
	// Skip the pipe character if present
	if start < len(message) && message[start] == '|' {
		start++
	}
	// Extract until the next pipe or end of string
	end := start
	for end < len(message) && message[end] != '|' {
		end++
	}
	requestID := message[start:end]
	// Remove any remaining pipes (shouldn't be any, but just in case)
	requestID = strings.ReplaceAll(requestID, "|", "")
	// Trim whitespace
	requestID = strings.TrimSpace(requestID)
	return requestID
}

// extractTrackingNumber extracts the tracking number from the payload
// Searches for multiple possible field names: waybillNo, deliveryID, airwayBillNumber,
// store_order_id, airwaybill_no, carrier_tn
// Returns empty string "" if not found
func extractTrackingNumber(payload string) string {
	// List of possible tracking number field names to search for
	fieldNames := []string{
		"waybillNo",
		"deliveryID",
		"airwayBillNumber",
		"store_order_id",
		"airwaybill_no",
		"carrier_tn", // Keep original as fallback
	}

	// Try each field name
	for _, fieldName := range fieldNames {
		if trackingNumber := extractFieldValue(payload, fieldName); trackingNumber != "" {
			return trackingNumber
		}
	}

	return ""
}

// extractFieldValue extracts the value for a given field name from the payload
func extractFieldValue(payload, fieldName string) string {
	// Search for field name with various quote patterns
	markers := []string{
		`"` + fieldName + `":`,
		`""` + fieldName + `"":`,
		`"` + fieldName + `" :`,
		`""` + fieldName + `"" :`,
	}

	var startIdx int = -1
	for _, marker := range markers {
		idx := strings.Index(payload, marker)
		if idx != -1 {
			startIdx = idx + len(marker)
			break
		}
	}

	if startIdx == -1 {
		return ""
	}

	// Skip whitespace after the colon
	for startIdx < len(payload) && (payload[startIdx] == ' ' || payload[startIdx] == '\t') {
		startIdx++
	}

	if startIdx >= len(payload) {
		return ""
	}

	// Skip opening quote(s) - could be " or ""
	for startIdx < len(payload) && payload[startIdx] == '"' {
		startIdx++
	}

	if startIdx >= len(payload) {
		return ""
	}

	// Find the end of the value (closing quote or comma)
	endIdx := startIdx
	for endIdx < len(payload) {
		if payload[endIdx] == '"' {
			// Found closing quote
			break
		}
		if payload[endIdx] == ',' || payload[endIdx] == '}' {
			// Found end of value
			break
		}
		endIdx++
	}

	if endIdx > startIdx {
		value := payload[startIdx:endIdx]
		// Remove any pipes if present
		value = strings.ReplaceAll(value, "|", "")
		// Trim whitespace
		value = strings.TrimSpace(value)
		return value
	}

	return ""
}

// connectDB establishes a connection to the MySQL database with optimized settings
func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Optimize connection pool for large batch processing
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0) // Reuse connections

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// insertBatch inserts a batch of log data using a transaction with bulk INSERT
func insertBatch(db *sql.DB, batch []logData) error {
	if len(batch) == 0 {
		return nil
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Build bulk INSERT query with multiple VALUES
	query := `INSERT INTO entry_task.log_data_raw 
		(id, payload, raw_message, request_id, timestamps, resource_id, tracking_number) 
		VALUES`

	// Build values placeholders and arguments
	values := make([]string, 0, len(batch))
	args := make([]interface{}, 0, len(batch)*7)

	for _, data := range batch {
		values = append(values, "(?, ?, ?, ?, ?, ?, ?)")
		args = append(args, data.id, data.payload, data.rawMessage, data.requestID, data.timestamps, data.resourceID, data.trackingNumber)
	}

	// Combine query with all values
	query += strings.Join(values, ", ")

	// Execute the bulk insert
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute bulk insert: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
