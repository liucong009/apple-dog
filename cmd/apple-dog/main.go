package main

import (
	"flag"
	"fmt"

	"github.com/liucong009/apple-dog/internal/greet"
)

func main() {
	name := flag.String("name", "", "name to greet")
	flag.Parse()
	fmt.Println(greet.Message(*name))
}
