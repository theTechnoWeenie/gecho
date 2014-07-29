package echoservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var startTime time.Time

type UptimeFormat struct {
	Miliseconds      int64
	HourMinuteSecond string
}

func main() {
	StartServer()
}

func StartServer() {
	startTime = time.Now()
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", root)
	http.HandleFunc("/uptime", retrieveUptime)
	http.ListenAndServe(":8080", connectionPrinter(http.DefaultServeMux))
}

func elapsedDuration() time.Duration {
	return time.Since(startTime)
}

func retrieveUptime(writer http.ResponseWriter, r *http.Request) {
	d := elapsedDuration()
	hourMinuteSecond := fmt.Sprintf("%02d:%02d:%02d", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60)
	uptimeStruct := UptimeFormat{d.Nanoseconds() / 1000 / 1000, hourMinuteSecond}
	responseJson, _ := json.Marshal(uptimeStruct)
	writer.Write(responseJson)
}

func echo(writer http.ResponseWriter, r *http.Request) {
	method := r.Method
	//Echo the query params for a get, and the body for a post.
	if method == "GET" {
		jsonString, _ := json.Marshal(r.URL.Query())
		writer.Write(jsonString)
	}
	if method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		writer.Write(body)
	}
}

func root(writer http.ResponseWriter, req *http.Request) {
	region := os.Getenv("REGION")
	if region == "" {
		region = "Development"
	}
	fmt.Fprintf(writer, "<html><body><h1>Echo Service serving <i>%s</i></h1></body></html>", region)
}

func connectionPrinter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
