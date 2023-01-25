package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	fmt.Println("child: Awaiting signal")
	s := <-c
	fmt.Println("child: Received signal:", s)
	time.Sleep(2 * time.Second)
	fmt.Println("child: Exiting")
}
