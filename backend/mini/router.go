package mini

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	handlers "react-demo/mini/handlers"
)

// func setStaticFolder(route *mux.Router) {
// 	fs := http.FileServer(http.Dir("./public/"))
// 	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
// }

// AddApproutes will add the routes for the application

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrader(rw http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	fmt.Println("upgrade the connection...")
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, nil
}

func chatHanlder(responseWriter http.ResponseWriter, request *http.Request, hub *handlers.Hub) {

	username := mux.Vars(request)["username"]

	// username := "van"

	connection, err := Upgrader(responseWriter, request)
	if err != nil {
		log.Println(err)
		return
	}

	handlers.CreateNewSocketUser(hub, connection, username)
}

func AppRoutes(route *mux.Router) {

	log.Println("Loadeding Routes...")

	// setStaticFolder(route)

	hub := handlers.NewHub()

	// hub具有注册与注销的功能
	go hub.Run()

	// route.HandleFunc("/", handlers.RenderHome)

	route.HandleFunc("/ws/{username}", func(responseWriter http.ResponseWriter, request *http.Request) {
		chatHanlder(responseWriter, request, hub)
	})

	log.Println("Routes are Loaded.")
}
