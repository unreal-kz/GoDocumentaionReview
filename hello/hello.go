package main

import (
	"fmt"

	"github.com/unreal-kz/GoDocumentaionReview/greetings"
)

func main() {
	fmt.Println("Testings")
	msg := greetings.Hello("Baubek")
	fmt.Println(msg)
}
