package main

import (
	"fmt"
	"golang-server/init/cmd"
	"net/http"
)

func main() {
	cmd.NewCmd()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}
