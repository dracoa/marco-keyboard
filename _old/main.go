package _old

type WebRequest struct {
	Command   string      `json:"command"`
	Parameter interface{} `json:"parameter"`
}

var onReceive bool

func main() {

	/*
		log.Println("Keyboard Control Server v1.0.0")
		server := websvr.Start("localhost:2303")

		go func() {
			EvChan := hook.Start()
			defer hook.End()
			for ev := range EvChan {
				if onReceive {
					server.SendJson(ev)
				}
			}
		}()

		for {
			select {
			case income := <-server.In:
				req := &WebRequest{}
				_ = json.Unmarshal(income, &req)
				if req.Command == "ScreenCapture" {
					server.SendBinary(screen.Capture())
				} else if req.Command == "StartReceiveEvent" {
					onReceive = true
				} else {
					onReceive = false
				}
			}
		}
	*/
}
