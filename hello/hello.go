package main

import (
	"fmt"

	"github.com/unreal-kz/GoDocumentaionReview/greetings"
)

func main() {
	msg := greetings.Hello("Baubek")
	fmt.Println(msg)
}
