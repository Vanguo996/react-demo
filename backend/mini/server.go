package mini

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Llongfile)
}

func ChatAppStart() {

	log.Println("Server will start at http://localhost:8080/")

	route := mux.NewRouter()

	AppRoutes(route)

	log.Fatal(http.ListenAndServe(":8080", route))
}
