package websvr

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Server struct {
	Router *gin.Engine
	Out    chan []byte
	In     chan []byte
}

func (s *Server) SendJson(content interface{}) {
	s.Out <- []byte{}
}

func (s *Server) SendBinary(content []byte) {
	s.Out <- content
}

func Start(bindAddress string) *Server {
	r := gin.Default()
	server := &Server{
		Router: r,
		In:     make(chan []byte),
		Out:    make(chan []byte),
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

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		s.In <- message
		select {
		case msg := <-s.Out:
			err = ws.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				break
			}
		}
	}
}
