package vote

import (
	"net/http"
	//"time"
)

const staticPath string = "./static/"

/*type Votes struct {
	ID       int
	Song1    int
	Song2    int
	Song3    int
	Song4    int
	Votes1   int
	Votes2   int
	Votes3   int
	Votes4   int
	DateTime time.Time
}*/

func Get(w http.ResponseWriter, r *http.Request) {

	js := "{\"test\":\"Werks\"}"
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}

func Post(w http.ResponseWriter, r *http.Request) {

	js := "{\"test\":\"Werks\"}"
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
