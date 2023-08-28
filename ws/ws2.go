package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"time"
)

// for number of concurrent users, create an array, or channel of with buffer / lenghh of number of concurrent users
// for each connection in the DS created,listen and send messages concurrently

var in = 10

// runListen simulates the send and receive for one concurrent request
func runListen() {
	var conns []*websocket.Conn
	wsData := make(chan ChanWs, in)
	defer close(wsData)
	var sigChan = make(chan os.Signal, 1)
	done := make(chan bool)
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "ws",
	}
	go func() {

		for i := 0; i < in; i++ {
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Println("error connecting to ws", err)
				// handle error
				os.Exit(1)

				log.Println("connected successfully")
			}
			conns = append(conns, conn)
		}

		for i := 0; i < in; i++ {
			log.Println("operation ", i)
			listenDone := make(chan bool)
			sendDone := make(chan bool)
			//log.Print("sending and receiving on conn ", i)
			// 2 go routine performing read and write on  the same connection RED FLAG!!
			go sendMsg(conns[i], sendDone)
			go listen(conns[i], wsData, listenDone)

			<-sendDone
			<-listenDone

			log.Println("both routines are done already")
			conns[i].WriteMessage(websocket.CloseMessage, []byte("close connection"))
			conns[i].Close()
			log.Println("connection closed")

		}
		close(done)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	for {
		log.Println("inside for stmt")
		select {

		case <-ctx.Done():
			log.Println("time out here")
			log.Println("Time out, gracefully closing")
			return

		case <-sigChan:
			log.Println("Signal received, gracefully closing...")
			return

		case w := <-wsData:
			log.Println("-----------------------------------------")
			log.Println("Msg: ", w.Message)
			log.Println("Error ", w.Error)

		case <-done:
			return
		}

	}

}

func listen(conn *websocket.Conn, ws chan ChanWs, ch chan bool) {
	var messageType string
	for {

		i, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		switch i {
		case websocket.TextMessage:
			messageType = "Text Message"
		case websocket.BinaryMessage:
			messageType = "Binary Message"
		case websocket.CloseMessage:
			messageType = "Close Message"
			log.Println("got error message... closing now ")
			ws <- ChanWs{
				Message:     string("Error message received"),
				MessageType: messageType,
				Error:       err,
			}
			// log or send out message
			break

		case websocket.PingMessage:
			messageType = "Ping Message"
		case websocket.PongMessage:
			messageType = "Pong Message"
		}

		ws <- ChanWs{
			Message:     string(message),
			MessageType: messageType,
			Error:       err,
		}
		ch <- true
	}
}

// todo:  add a context to facilitate timeout
func sendMsg(conn *websocket.Conn, ch chan bool) {

	// change 100 to number of concurrent users sending a message
	// also simulate how many times you want to send the batch concurrent request

	for i := 0; i < 3; i++ {
		err := conn.WriteMessage(websocket.TextMessage, []byte("SENT"))
		if err != nil {
			// handle error
			return
		}
	}
	ch <- true

}
