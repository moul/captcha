package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"moul.io/banner"
	"moul.io/srand"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(_ []string) error {
	rand.Seed(srand.Fast())

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] // nolint:gosec
	}
	fmt.Println(banner.Inline(string(b)))
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.ReplaceAll(text, "\n", "")
		if strings.Compare(string(b), text) == 0 {
			return nil
		}
	}
	return fmt.Errorf("too many fails") // nolint:goerr113
}
