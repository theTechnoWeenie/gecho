package echoservice

import (
	"encoding/json"
	"fmt"
	"github.com/theTechnoWeenie/greg/server"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
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

func main() {
	StartServer()
}

func StartServer() {
	startTime = time.Now()
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", root)
	http.HandleFunc("/uptime", retrieveUptime)
	http.HandleFunc("/healthCheck", verifyHealth)
	http.ListenAndServe(":8080", connectionPrinter(http.DefaultServeMux))
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
	if err == nil {
		fmt.Printf("ERR: Malformed JSON: %s\n", err.Error())
	}
	writer.Write(responseJson)
}

func verifyHealth(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(fmt.Sprintf(`{"status":"%s"}`, server.STATUS_RUNNING)))
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
	region, gregAddress := parseEnv()
	t, err := template.ParseFiles("echoservice/templates/index.html")
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
	}
	context := &Template{Region: region, GregAddress: gregAddress}
	t.Execute(writer, context)
}

func parseEnv() (string, string) {
	region := os.Getenv("REGION")
	if region == "" {
		region = "Development"
	}
	gregAddress := os.Getenv("GREG_ADDRESS")
	if gregAddress == "" {
		gregAddress = "Undefined"
	}
	return region, gregAddress
}
