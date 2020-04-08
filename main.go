package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/timdong/golang-protobuf-websocket/message"
	"html/template"
	"net/http"
)

var data []byte

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Handler read done")
		fmt.Printf("Handler write %s", data)
		if err = conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			fmt.Printf("err %s", err)
			return
		}
		fmt.Println("write done")
	}
}

func handlerHtml(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./public/index.html")
	t.Execute(w, "Hello")
}

func main() {

	msg := &message.Message{
		Id: *proto.Int32(17),
		Author: &message.Message_Person{
			Id:   *proto.Int32(1),
			Name: *proto.String("othree"),
		},
		Text: *proto.String("Hi, this is message."),
	}

	fmt.Println(msg.GetAuthor().GetName() + ": " + msg.GetText())

	data, _ = proto.Marshal(msg)

	fmt.Printf("After Marshal: %s", data)

	mux := http.NewServeMux()

	mux.HandleFunc("/ws", handler)

	go http.ListenAndServe(":1337", mux)

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/index", handlerHtml)
	files := http.FileServer(http.Dir("./public"))
	mux1.Handle("/public/", http.StripPrefix("/public/", files))

	go http.ListenAndServe(":8080", mux1)

	//block
	select {}

}
