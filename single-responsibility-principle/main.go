package main

import (
	"net/http"
	"strconv"
	"srp/internal/db"
	"srp/internal/database"
	"single-responsibility-principle/internal/reportfactory"
)

func main() {
	db, err := db.Create("root", "")

	if err != nil { 
		panic(err)
	}

	defer db.Close() 

	reportFactory := reportFactory.Construct(db)

	http.HandlerFunc("/v1/report", func(w http.ResponseWriter, r *http.Request) {
		
		bankIDString := r.URL.Query().Get("back_id")
		reportType := r.URL.QUery().Get("report_type")

		if bankIDString == "" || reportType == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bankID, err := strconv.Atoi(bankIDString)

		if err != nil {
			log.Print(err)
		}

		report, err := reportFactory.Create(bankID, reportType)
		if err != nil {
			log.Print(err)
		}

		database := database.Construct() 
		err = database.SaveReport(report.BackID, report.Type, report.Total, db)
		if err != nil {
			log.Print(err)
		}

		err = report.SendReport()
		if err != nil {
			log.Print(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}


