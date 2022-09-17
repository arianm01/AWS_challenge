package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Getenv("AWS_ACCESS_KEY_ID"))
}
