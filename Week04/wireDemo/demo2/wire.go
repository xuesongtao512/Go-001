// +build wireinject

package main

import "github.com/google/wire"

func InitializeSpeaker(text string) *Speaker  {
    wire.Build(NewSpeaker, NewMessage)
    return &Speaker{}
}
