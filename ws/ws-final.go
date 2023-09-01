package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"time"
)

var u = url.URL{
	Scheme: "ws",
	Host:   "localhost:8080",
	Path:   "ws",
}
var conns []*websocket.Conn

func final(ctx context.Context) {
	var wsMessageChan = make(chan ChanWs, 30)
	fDone := make(chan bool)
	for i := 0; i < 10; i++ {
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Println("error connecting to ws", err)
			// handle error
			os.Exit(1)
		}
		log.Println("connected successfully")
		conns = append(conns, conn)
	}

	go func() {
		simulate(conns, wsMessageChan)

		close(wsMessageChan)
		close(fDone)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-fDone:
			log.Println("FDone")
			return
		}
	}
}

func simulate(conn []*websocket.Conn, ch chan ChanWs) {
	sg := make(chan bool)
	count := 0

	// send and receive on 10 concurrent users
	for i := 0; i < 10; i++ {
		go sendRec(conn[i], ch, sg, i)
	}

	// receives done signal 10 times before proceeding.
	for i := 0; i < 10; i++ {
		<-sg
	}

	for msg := range ch {
		log.Println("------------------------------------")
		log.Println("User is: ", msg.User)
		log.Println("Message is: ", msg.Message)
		log.Println("Error is: ", msg.Error)
		log.Println("Time take is: ", msg.ResponseTime)
		log.Println("------------------------------------", count)
		count++
		if count >= 30 {
			break
		}
	}
	log.Println("there")
	log.Println("exiting")
}

// this function should send messages and receive messages into the WS channel. Upon receiving 3 messages into WSChan
// the function returns, then sends a signal that it has received 3 messages (Synchronization)
func sendRec(conn *websocket.Conn, msgData chan ChanWs, signal chan bool, index int) {
	var nT time.Time
	// this baby ensures that receiving works until the message count is more than 3
	messageReceivedCount := 0
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("ReadMessage.Error:", err)
				return
			}
			log.Println("Received message")

			msgData <- ChanWs{
				User:         fmt.Sprint("user ", index),
				Message:      string(message),
				Error:        nil,
				ResponseTime: fmt.Sprintf("%.5f s", time.Since(nT).Seconds()),
			}

			messageReceivedCount++
			if messageReceivedCount >= 3 {
				return
			}
		}

	}()
	// while the go routine above is waiting to receive, we now send messages into it
	nT = time.Now()
	for i := 0; i < 3; i++ {
		err := conn.WriteMessage(websocket.TextMessage, []byte("SENT"))
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}

	log.Println("exiting sendRec")
	signal <- true
	log.Println("exited")
}
