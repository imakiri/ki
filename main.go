package main

import (
	"fmt"
	"github.com/imakiri/ki/interpreter"
	"log"
	"os"
)

func main() {
	var f, err = os.Open("test")
	if err != nil {
		log.Fatal(err)
	}

	var w interpreter.Worker
	if w, err = interpreter.NewJob(f); err != nil {
		log.Fatal(err)
	}

	if err = w.Parse(); err != nil {
		log.Fatal(err)
	}

	for _, v := range w.Interpret() {
		fmt.Println(v)
	}
}
