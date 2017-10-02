package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// Print the process ID to make it easier to send this program a signal.
	fmt.Printf("pid: %d\n", os.Getpid())

	// If this program's first argument is "trap" then trap SIGABRT.
	if len(os.Args) > 1 && strings.EqualFold(os.Args[1], "trap") {
		n := make(chan os.Signal, 1)
		signal.Notify(n, syscall.SIGABRT)
		go func() {
			for sig := range n {
				fmt.Println(sig)
				os.Exit(1)
			}
		}()
	}

	// Use a channel to block the program by waiting indefinitely
	// until something reads the channel (which will never happen).
	c := make(chan struct{})
	<-c
}
