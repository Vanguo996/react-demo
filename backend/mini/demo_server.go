package mini

type Message struct {
	// Message string `json:"message"`
	Type int    `json:"type"`
	Body string `json:"body"`
}

// func readJson(conn *websocket.Conn) {

// 	var message Message

// 	err := conn.ReadJSON(message)

// 	if err != nil {
// 		log.Printf("readJson error: %v", err)
// 	}
// 	log.Println(message)

// 	//send
// 	if err = conn.WriteJSON(message); err != nil {
// 		log.Printf("writeJson error: %v", err)
// 	}
// }

// func WsServer() {
// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	e.GET("/ws", func(c echo.Context) error {

// 		// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 		conn, err := Upgrader(c.Response().Writer, c.Request())

// 		if err != nil {
// 			log.Println(err)
// 		}

// 		defer conn.Close()

// 		log.Println("ws Connected!!")

// 		// 添加一个for循环，这个循环监听来自客户端的任何消息，读取json之后，返回给客户端
// 		for {

// 			messageType, p, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			message := Message{Type: messageType, Body: string(p)}
// 			b, _ := json.Marshal(message)

// 			for {
// 				if err := conn.WriteMessage(messageType, b); err != nil {
// 					log.Println(err)
// 				}
// 				time.Sleep(time.Second * 5)
// 			}

// 		}

// 		// return c.String(http.StatusOK, "Hi, WebSocket!")
// 		// return nil
// 	})

// 	e.Logger.Fatal(e.Start(":8080"))
// }
