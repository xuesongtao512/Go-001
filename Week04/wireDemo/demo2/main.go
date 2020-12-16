// @Desc: 依赖注入
// @CreateTime: 2020/12/16
package main

import "fmt"

type Message string

func NewMessage(text string) Message {
    return Message(text)
}

type Speaker struct {
    Message Message
}

func (s *Speaker) Say() {
    fmt.Println(s.Message)
}

func NewSpeaker(m Message) *Speaker {
    return &Speaker{Message: m}
}

func main() {
    s := InitializeSpeaker("hello, I am a speaker")
    s.Say()
}