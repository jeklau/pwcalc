package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func calcPw(secret string, alias string, length int) string {
	str := []byte(secret + alias)
	hash := fmt.Sprintf("%s", sha1.Sum(str))
	base64 := base64.StdEncoding.EncodeToString([]byte(hash))[0:length]
	return base64
}

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
		os.Exit(0)
	}()
	return nil
}

func main() {
	var alias string
	var length = 16
	var err error

	if runtime.GOOS != "windows" {
		err = signalHandler()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if len(os.Args) > 1 {
		alias = os.Args[1]
		fmt.Printf("# alias: %s\n", alias)
	}
	if len(os.Args) > 2 {
		length, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if length > 28 || length < 1 {
			fmt.Println("Error: length must be between 1 and 28")
			os.Exit(1)
		}
	}

	if alias == "" {
		fmt.Print("# alias: ")
		reader := bufio.NewReader(os.Stdin)
		alias, _ = reader.ReadString('\n')
	}

	fmt.Print("# secret: ")
	input, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()

	secret := string(input)
	alias = strings.TrimSpace(alias)
	secret = strings.TrimSpace(secret)

	// print easy to remember hint based on secret
	fmt.Printf("%s\n", calcPw(secret, secret, 3))

	fmt.Printf("%s\n", calcPw(secret, alias, length))
}
