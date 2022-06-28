package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// https://www.etix.com/ticket//online/timedEntrySearch.do?method=getMonthSummaryData&event_id=1002345&year=2022&month=7
	result, err := http.Get("https://www.etix.com/ticket//online/timedEntrySearch.do?method=getMonthSummaryData&event_id=1002345&year=2022&month=7")
	if err != nil {
		log.Fatal(err)

	}

	defer result.Body.Close()

	dec := json.NewDecoder(result.Body)

	m := map[string]interface{}{}
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	for k := range m {
		fmt.Println(k)
		fmt.Println(m[k])
	}
	fmt.Println(m["A-#fafafa"])
}
