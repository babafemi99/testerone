package ws

//
//import (
//	"context"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/url"
//	"os"
//	"sync"
//	"time"
//)
//
//var messageType string
//var wsData = make(chan ChanWs)
//var sigChan = make(chan os.Signal, 1)
//var done = make(chan struct{})
//
////func Run() {
////
////	// create new http request and add the headers to it, then pass it into Dial parameter.
////
////	u := url.URL{
////		Scheme: "ws",
////		Host:   "localhost:8080",
////		Path:   "ws",
////	}
////
////	for i := 0; i < 5; i++ {
////		conn, response, err := websocket.DefaultDialer.Dial(u.String(), nil)
////		if err != nil {
////			// handle error
////			log.Println("Error connecting to web server...")
////			time.Sleep(2 * time.Second)
////			continue
////		}
////		log.Println(response.Status)
////	}
////}
//
//// runListen simulates the send and receive for one concurrent request
//func runListen() {
//
//	u := url.URL{
//		Scheme: "ws",
//		Host:   "localhost:8080",
//		Path:   "ws",
//	}
//	var wg sync.WaitGroup
//	conn, response, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Println("error connecting to ws", err)
//		// handle error
//		return
//	}
//	log.Println("connected successfully")
//	ctx := context.Background()
//	log.Println("status is ", response.Status)
//	go sendMsg(conn)
//	go listen(ctx, conn, wg)
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("time out here")
//			log.Println("Time out, gracefully closing")
//			wg.Wait()
//			return
//		case <-sigChan:
//			wg.Wait()
//			log.Println("Signal received, gracefully closing...")
//			return
//		case <-done:
//			wg.Wait()
//			log.Println("Error received, gracefully closing...")
//			return
//		}
//	}
//
//	time.Sleep(1 * time.Minute)
//}
//
//func listen(ctx context.Context, conn *websocket.Conn, wg sync.WaitGroup) {
//	// channels to listen on
//	wg.Add(1)
//
//	go func() {
//		defer wg.Done()
//		defer close(wsData)
//		i, message, err := conn.ReadMessage()
//		if err != nil {
//			return
//		}
//		log.Println("message received from conn", string(message))
//		log.Println("message type from conn", i)
//
//		switch i {
//		case websocket.TextMessage:
//			messageType = "Text Message"
//		case websocket.BinaryMessage:
//			messageType = "Binary Message"
//		case websocket.CloseMessage:
//			messageType = "Close Message"
//			conn.Close()
//			wsData <- ChanWs{
//				Message:     string(message),
//				MessageType: messageType,
//				Error:       err,
//			}
//			// log or send out message
//			close(done)
//			return
//
//		case websocket.PingMessage:
//			messageType = "Ping Message"
//		case websocket.PongMessage:
//			messageType = "Pong Message"
//		}
//
//		wsData <- ChanWs{
//			Message:     string(message),
//			MessageType: messageType,
//			Error:       err,
//		}
//
//		// log or send out message
//		for ws := range wsData {
//			log.Println("-----------------------------------------")
//			log.Println("Msg: ", ws.Message)
//			log.Println("Error ", ws.Error)
//		}
//	}()
//
//}
//
//// todo:  add a context to facilitate timeout
//func sendMsg(conn *websocket.Conn) {
//	// change 100 to number of concurrent users sending a message
//	// also simulate how many times you want to send the batch concurrent request
//
//	for i := 0; i < 100; i++ {
//		err := conn.WriteMessage(websocket.TextMessage, []byte("SENT"))
//		if err != nil {
//			// handle error
//			return
//		}
//	}
//
//}
