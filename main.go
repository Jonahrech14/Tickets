package main

import (
	"log"
	"net/http"
)

func main() {
	// https://www.etix.com/ticket//online/timedEntrySearch.do?method=getMonthSummaryData&event_id=1002345&year=2022&month=3
	result, err := http.Get("https://www.etix.com/ticket//online/timedEntrySearch.do?method=getMonthSummaryData&event_id=1002345&year=2022&month=3")
	if err != nil {
		log.Fatal(err)
	}
	result.Body.Close()
}
