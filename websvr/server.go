package websvr

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Server struct {
	Router *gin.Engine
	Out    chan *OutMessage
	In     chan []byte
}

type OutMessage struct {
	MessageType int
	Data        []byte
}

func (s *Server) SendJson(content interface{}) {
	bytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	s.Out <- &OutMessage{
		MessageType: websocket.TextMessage,
		Data:        bytes,
	}
}

func (s *Server) SendBinary(content []byte) {
	s.Out <- &OutMessage{
		MessageType: websocket.BinaryMessage,
		Data:        content,
	}
}

func Start(bindAddress string) *Server {
	r := gin.Default()
	server := &Server{
		Router: r,
		In:     make(chan []byte),
		Out:    make(chan *OutMessage),
	}
	r.Static("/static", "./static")
	r.GET("/wsEndpoint", server.wsEndpoint)
	go r.Run(bindAddress)
	return server
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) wsEndpoint(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	go func() {
		for {
			select {
			case msg := <-s.Out:
				err = ws.WriteMessage(msg.MessageType, msg.Data)
				if err != nil {
					break
				}
			}
		}
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		s.In <- message
	}
}
