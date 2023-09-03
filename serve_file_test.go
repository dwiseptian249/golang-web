package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func Servefile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8181",
		Handler: http.HandlerFunc(Servefile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServefileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcesOk)
	} else {
		fmt.Fprint(writer, resourcesNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8181",
		Handler: http.HandlerFunc(ServefileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
