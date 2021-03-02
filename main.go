package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func getdtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format(time.RFC3339))
}

func startServer(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/tests/get-dt", getdtHandler)
	http.Handle("/", router)

	fmt.Printf("Server is listening at: http://%s ...", addr)

	http.ListenAndServe(addr, nil)
}

func main() {
	var portPtr = flag.Int("port", 8181, "TCP/IP port")
	flag.Parse()

	addr := "0.0.0.0:" + strconv.Itoa(*portPtr)

	startServer(addr)
}
