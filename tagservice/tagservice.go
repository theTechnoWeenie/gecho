package tagservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

type UptimeFormat struct {
	Miliseconds      int64
	HourMinuteSecond string
}

type Template struct {
	Region      string
	GregAddress string
}

func StartServer() {
	startTime = time.Now()
	http.HandleFunc("/", root)
	http.HandleFunc("/uptime", retrieveUptime)
	http.HandleFunc("/healthCheck", verifyHealth)
	err := http.ListenAndServe(":8080", connectionPrinter(http.DefaultServeMux))
	fmt.Printf("%s", err)
}

func connectionPrinter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func retrieveUptime(writer http.ResponseWriter, r *http.Request) {
	d := time.Since(startTime)
	hourMinuteSecond := fmt.Sprintf("%02d:%02d:%02d", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60)
	uptimeStruct := UptimeFormat{d.Nanoseconds() / 1000 / 1000, hourMinuteSecond}
	responseJson, err := json.Marshal(uptimeStruct)
	if err != nil {
		fmt.Printf("ERR: Malformed JSON: %s\n", err.Error())
	}
	writer.Write(responseJson)
}

func verifyHealth(writer http.ResponseWriter, req *http.Request) {
	// TODO impl
}

func root(writer http.ResponseWriter, req *http.Request) {
	// TODO impl
}

