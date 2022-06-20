package main

import (
	"log"
	"net/http"

	helpers "csv-processor/helpers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload-csv", func(c *gin.Context) {
		file_ptr, err := c.FormFile("file")
		c.Request.ParseForm()
		department := c.Request.Form.Get("department")
		project := c.Request.Form.Get("project")
		user := c.Request.Form.Get("user")
		email := c.Request.Form.Get("email")

		if err != nil {
			log.Println(err.Error())
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		log.Println(file_ptr.Filename)
		file, err := file_ptr.Open()
		if err != nil {
			log.Println(err.Error())
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		defer file.Close()
		rows := helpers.CSVToMap(file)
		clockify_entries := make([]helpers.Clockify, len(rows))

		for i, row := range rows {
			duration := helpers.ConvertTimeToDuration(row["Duration"])
			clockify_entries[i] = helpers.Clockify{
				Project:         project,
				Department:      department,
				Description:     row["Description"],
				Task:            row["Task"],
				User:            user,
				Email:           email,
				Tags:            row["Tags"],
				Billable:        "Yes",
				StartDate:       row["Start date"],
				StartTime:       row["Start time"],
				EndDate:         row["End date"],
				EndTime:         row["End time"],
				DurationHours:   row["Duration"],
				DurationDecimal: float64(int(duration.Hours()*100)) / 100,
				BillableRate:    0,
				BillableAmount:  0,
			}
		}
		b := helpers.MapToCSV(clockify_entries)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename=clockify-entries.csv")
		c.Data(http.StatusOK, "text/csv", b.Bytes())
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
