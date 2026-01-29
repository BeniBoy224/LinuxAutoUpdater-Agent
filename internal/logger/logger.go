/*
 - Log errors
 - log successfull operations
 - Return to api in a json format
*/


package logger

import (
	"fmt"
	"log"
)

func Log(message string) {
	fmt.Println(message)
	log.Fatal("Fatal log: ", message)
}
