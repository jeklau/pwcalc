package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
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

func main() {
	var alias string
	var length = 16
	var err error

	err = signalHandler()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		alias = os.Args[1]
		fmt.Printf("# alias: %s\n", alias)
	}
	if len(os.Args) > 2 {
		length, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
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
		fmt.Fprintf(os.Stderr, "%v\n", err)
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
