package helpers

type Clockify struct {
	Project         string
	Department      string
	Description     string
	Task            string
	User            string
	Email           string
	Tags            string
	Billable        string
	StartDate       string
	StartTime       string
	EndDate         string
	EndTime         string
	DurationHours   string
	DurationDecimal float64
	BillableRate    int
	BillableAmount  int
}
