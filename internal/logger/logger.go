package logger

import (
	"fmt"
	"log"
)

func Log(message string) {
	fmt.Println(message)
	log.Fatal("Fatal log: ", message)
}
