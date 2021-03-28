package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"moul.io/captcha/captcha"
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

	// init captcha and print the question
	captcha := captcha.NewBannerCaptcha()
	question, err := captcha.Question()
	if err != nil {
		return fmt.Errorf("init captcha: %w", err)
	}
	fmt.Println(question)

	// read and check the user input
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.ReplaceAll(text, "\n", "")
		valid, err := captcha.Validate(text)
		if err != nil {
			return fmt.Errorf("validate answer: %w", err)
		}
		if valid {
			return nil
		}
	}
	return fmt.Errorf("too many fails") // nolint:goerr113
}
