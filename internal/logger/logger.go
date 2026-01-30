/*
 Use of file - To pass logs back to central server.
 For local logs use fmt log package

 - Log errors
 - log successfull operations
 - Return to api in a json format
*/

package logger

import (
	"fmt"
)

func Log(message string) {
	fmt.Println(message)
}
