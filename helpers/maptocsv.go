package helpers

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func MapToCSV(clockify_entries []Clockify) *bytes.Buffer {
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	// define column headers
	headers := []string{
		"Project",
		"Department",
		"Description",
		"Task",
		"User",
		"Email",
		"Tags",
		"Billable",
		"Start Date",
		"Start Time",
		"End Date",
		"End Time",
		"Duration (h)",
		"Duration (decimal)",
		"Billable Rate (USD)",
		"Billable Amount (USD)",
	}

	// write column headers
	writer.Write(headers)

	for _, entry := range clockify_entries {

		r := make([]string, 0, 1+len(headers)) // capacity of 4, 1 + the number of properties your struct has & the number of column headers you are passing

		// convert the Record.ID to a string in order to pass into []string

		r = append(
			r,
			entry.Project,
			entry.Department,
			entry.Description,
			entry.Task,
			entry.User,
			entry.Email,
			entry.Tags,
			entry.Billable,
			entry.StartDate,
			entry.StartTime,
			entry.EndDate,
			entry.EndTime,
			entry.DurationHours,
			fmt.Sprintf("%.2f", entry.DurationDecimal),
			fmt.Sprintf("%d", entry.BillableRate),
			fmt.Sprintf("%d", entry.BillableAmount),
		)

		writer.Write(r)
	}
	writer.Flush()

	return b
}
