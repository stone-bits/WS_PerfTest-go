package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func getdtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format(time.RFC3339))
}

func normalizePathBase(pathBase string) string {
	if pathBase == "" {
		return pathBase
	}

	if strings.HasPrefix(pathBase, "/") == false {
		pathBase = strings.TrimLeft(pathBase, "/")
	}

	if strings.HasSuffix(pathBase, "/") {
		pathBase = strings.TrimRight(pathBase, "/")
	}

	pathBase = "/" + pathBase
	return pathBase
}

func startServer(addr string, pathBase string) {
	router := mux.NewRouter()

	app := router.PathPrefix(pathBase).Subrouter()
	app.HandleFunc("/tests/get-dt", getdtHandler)

	fmt.Printf("Server is listening at: http://%s%s ...", addr, pathBase)

	http.ListenAndServe(addr, router)
}

func main() {
	var portPtr = flag.Int("port", 8181, "TCP/IP port")
	var pathBasePtr = flag.String("pathbase", "", "Path base")

	flag.Parse()

	var pathBase = normalizePathBase(*pathBasePtr)

	addr := "0.0.0.0:" + strconv.Itoa(*portPtr)

	startServer(addr, pathBase)
}
