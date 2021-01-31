package main

import (
    "fmt"
	"net/http"
	"os"
	"encoding/json"
	"net/url"
	"strconv"
	"log"
)

func startup() {
	_, username := os.LookupEnv("DB_USER")
	_, password := os.LookupEnv("DB_USER_PASS")
	_, db_name := os.LookupEnv("DB_NAME")
	_, endpoint := os.LookupEnv("DB_ENDPOINT")

	if !(username && password && db_name && endpoint) {
		log.Fatalln("Environment Variables not set")
	}

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func recommend(w http.ResponseWriter, req *http.Request) {
	filters := req.URL.RawQuery // Everything after the ?

	//Functionality to check valid topic

	if filters == "" {
		category := ""
		topics := get_topics(category)

		//Encode
		w.Header().Set("Content-Type", "application/json") 
		json.NewEncoder(w).Encode(topics)
	} else {
		filter_map, err :=  url.ParseQuery(filters)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(filter_map)

		//Check that queries are valid and have 0 
		//CheckQueryValidity()
		minutes_reading, err := strconv.Atoi(filter_map["minutes_reading"][0])
		if err != nil {
			fmt.Println(err)
			// DO A CHECK HERE SUCH THAT IF NO MINUTES SPECIFIED , MINUTES = 0
		}
		topic := filter_map["topic"][0]

		//OLD FEATURE
		//topic := "kubernetes" //FOR DEBUG
		//query_db_with_topic(topic)
		//results := query_db_with_topic(topic)

		results := query_db_with_topic_and_minutes(topic, minutes_reading)

		if len(results) == 0 {
			// If for whatever reason no results are returned, then go to the google search for the topic
			// This may be due to lack of communication with the db
			fmt.Fprintf(w, "https://www.google.com/search?q="+topic)
		} else {
			recommendation := results[recommend_randon_number(len(results))]

			//Encode
			w.Header().Set("Content-Type", "application/json") 
			json.NewEncoder(w).Encode(recommendation)
		}
	}
}

// VALID PATHS:
// - /api/recommend[?<topic>]
// - /hello
// - /headers

func main() {
	log.Println("Starting saga-recommender-api server...")
	log.Println("Checking for environment variables")
	startup()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/api/recommend", recommend)
	
	http.ListenAndServe(":8090", nil)
}