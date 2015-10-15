package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gocraft/dbr"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/searchTxComment", txCommentSearch)
	fmt.Println("Listening on port 5831...")
	log.Fatal(http.ListenAndServe(":5831", nil))
}

func setHeaders(w http.ResponseWriter, r *http.Request, method string) http.ResponseWriter {
	endpoint := r.URL.Path[1:]
	rv := w
	rv.Header().Set("Access-Control-Allow-Origin", "*")
	rv.Header().Add("Access-Control-Allow-Methods", method)
	rv.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	fmt.Printf("%v %v - %v sent %v to %v (", r.URL, r.Method, r.RemoteAddr, r.Method, endpoint)
	return rv
}

func txCommentSearch(w http.ResponseWriter, r *http.Request) {
	// read POST body
	setHeaders(w, r, "POST")

	type APIRequest struct {
		SearchTerm     string `json:"search"`
		Page           int    `json:"page"`
		ResultsPerPage int    `json:"results-per-page"`
	}
	var body APIRequest
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: reading POST body failed\n%v\n", err)
		return
	}
	if bodyBytes == nil || len(bodyBytes) < 1 {
		log.Printf("bodyBytes is nil or couldn't be parsed\n")
		return
	}
	fmt.Printf("%v bytes)\n", len(bodyBytes))

	// parse POST json
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Printf("ERROR: json unmarshal failed\n%v", err)
		return
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")

	db, _ := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/floblockexplorer")
	connection := dbr.NewConnection(db, nil)

	// Create a session for each business unit of execution (e.g. a web request or goworkers job)
	dbrSess := connection.NewSession(nil)

	// Get a record
	type Result struct {
		Hash    dbr.NullString `db: "hash", json: "hash"`
		Message dbr.NullString `db: "message", json: "message"`
	}

	var results []*Result
	page := uint64(body.Page)
	resultsPerPage := uint64(body.ResultsPerPage)
	if resultsPerPage > 30 {
		resultsPerPage = 30
	}

	builder := dbrSess.Select("hash, message").From("tx").Where("message LIKE ?", "%"+body.SearchTerm+"%").Offset(page * resultsPerPage).Limit(resultsPerPage)

	_, err = builder.LoadStructs(&results)
	if err != nil {
		log.Printf("ERROR: database query failure\n%v\n", (err.Error()))
		return
	} else {
		//fmt.Printf("results: %v\n", results)
		json, err := json.Marshal(results)
		if err != nil {
			fmt.Printf("ERROR marshaling json\n%v\n", err)
			return
		}
		//fmt.Printf("%v\n", string(json[:len(json)]))
		fmt.Fprintf(w, "%v", string(json))
	}
}
