// @Desc:
// @CreateTime: 2021/1/27
package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		Port: 9999,
	})
	if err != nil {
		log.Fatalln("dial tcp is failed, err: ", err)
	}
	defer conn.Close()
	exitCh := make(chan bool)
	go sendMsg(conn, exitCh)
	go readMsg(conn, exitCh)
	<-exitCh
}

func sendMsg(conn net.Conn, exitCh <-chan bool) {
	rf := bufio.NewReader(os.Stdin)
	for {
		log.Print("请输入: ")
		b, err := rf.ReadBytes('\n')
		if err != nil {
			log.Println("input is err: ", err)
			break
		}
		_, err = conn.Write(b)
		if err != nil {
			log.Println("write send msg is failed, err:", err)
			break
		}
	}
}

func readMsg(conn net.Conn, exitCh chan<- bool) {
	rf := bufio.NewReader(conn)
	for {
		msg, err := rf.ReadString('\n')
		if err != nil {
			log.Println("client read msg is failed, err: ", err)
			close(exitCh)
			break
		}
		msg = strings.Replace(msg, "\n", "", -1)
		log.Println("recv server msg: ", msg)
	}
}
