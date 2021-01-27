// @Desc:
// @CreateTime: 2021/1/27
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 9999})
	if err != nil {
		log.Fatalln("listen tcp is failed, err:", err)
	}
	var isStop bool
	exitCh := make(chan bool)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listen accept is failed, err: ", err)
			continue
		}
		go handleFunc(conn, exitCh)

		select {
		case <-exitCh:
			isStop = true
		}
		if isStop {
			break
		}
	}
}

func handleFunc(conn net.Conn, exitCh chan<- bool) {
	// 消息管道, readMsg 收到的消息通过这个 chan 传递到 writeMsg,
	msgCh := make(chan string, 10)
	go readMsg(msgCh, exitCh, conn)
	go writeMsg(msgCh, exitCh, conn)
}

func writeMsg(msgCh <-chan string, exitCh chan<- bool, conn net.Conn) {
	var isStop bool
	for {
		select {
		case resp := <-msgCh:
			respByte := []byte(fmt.Sprintf("you are send msg: %s\n", resp))
			log.Println("resp to client...")
			conn.Write(respByte)
			// 如果收到 exit 的话, 程序退出
			if strings.Contains(resp, "exit") {
				isStop = true
			}
		}
		if isStop {
			break
		}
	}
}

func readMsg(msgCh chan<- string, exitCh chan<- bool, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println("readMsg is failed, err: ", err)
			break
		}
		msg = strings.Replace(msg, "\n", "", -1)
		log.Println("recv msg:", msg)
		msgCh <- msg
		// 如果收到 exit 的话, 程序退出
		if strings.Contains(msg, "exit") {
			close(exitCh)
			break
		}
	}
}
