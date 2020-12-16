// @Desc: 非依赖注入
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

func NewSpeaker(text string) *Speaker {
    m := NewMessage(text)
    return &Speaker{Message: m}
}

func main() {
    s := NewSpeaker("hello")
    s.Say()
}