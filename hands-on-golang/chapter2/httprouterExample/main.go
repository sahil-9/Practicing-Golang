package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/local/go/bin/go", "version")
	io.WriteString(w, response)
	return
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, getCommandOutput("bin/cat", params.ByName("name")))
	return
}