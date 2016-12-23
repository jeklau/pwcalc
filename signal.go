// +build !windows

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func signalHandler() error {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	err = terminal.Restore(0, oldState)
	if err.Error() != "errno 0" {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		terminal.Restore(0, oldState)
		fmt.Println()
		os.Exit(0)
	}()
	return nil
}
