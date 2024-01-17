package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

type GameServer struct {
}

func newGameServer() actor.Receiver {
	return &GameServer{}
}

func (s *GameServer) startHTTP() {
	fmt.Println("starting HTTP Server on port 40000")
	go func() {
		http.HandleFunc("/ws", s.handleWS)
		http.ListenAndServe(":40000", nil)
	}()
}
func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		_ = msg
	}
}

// handles upgrade of websockets
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("ws upgrade err:", err)
	}
	fmt.Print("new client trying to connect")
	fmt.Print(conn)
}

func main() {
	engineConfig := actor.NewEngineConfig()
	e, err := actor.NewEngine(engineConfig)
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(newGameServer, "server")
	select {}
}
