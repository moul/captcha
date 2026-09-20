package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/captcha/captcha"
	"moul.io/srand"
)

func main() {
	if err := run(os.Args, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, inputStream io.Reader) error {
	// CLI configuration
	var (
		fs           = flag.NewFlagSet("generate-fake-data", flag.ExitOnError)
		seed         = fs.Int64("seed", 0, "random seed (0 means automatic)")
		maxRetries   = fs.Uint64("retries", 3, "failures allowed (0 means unlimited)")
		engineChoice = fs.String("engine", "random", "captcha engine to use (banner, math, random by default)")
		timeOut      = fs.Uint64("timeout", 0, "timeout(in seconds) after which the program exits, 0 turns off timeout")
	)

	root := &ffcli.Command{
		Name:    "captcha [FLAGS]",
		FlagSet: fs,
		Exec: func(ctx context.Context, args []string) error {
			if len(args) > 0 {
				return flag.ErrHelp
			}
			if *seed != 0 {
				rand.Seed(*seed)
			} else {
				rand.Seed(srand.Fast())
			}

			// init captcha and print the question
			var engine captcha.Captcha
			switch *engineChoice {
			case "math":
				engine = captcha.NewMathCaptcha()
			case "banner":
				engine = captcha.NewBannerCaptcha()
			case "random":
				engine = captcha.NewRandomCaptcha()
			default:
				return fmt.Errorf("undefined captcha engine: %q", *engineChoice) // nolint:goerr113
			}
			question, err := engine.Question()
			if err != nil {
				return fmt.Errorf("init captcha: %w", err)
			}
			fmt.Println(question)

			// Exit once the timeout is reached
			if *timeOut != 0 {
				go QuitAfter(timeOut)
			}

			// read and check the user input
			reader := bufio.NewReader(inputStream)
			for i := uint64(0); *maxRetries == 0 || i < *maxRetries; i++ {
				fmt.Print("-> ")
				text, _ := reader.ReadString('\n')
				// convert CRLF to LF
				text = strings.ReplaceAll(text, "\n", "")
				valid, err := engine.Validate(text)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					continue
				}
				if valid {
					return nil
				}
			}
			return fmt.Errorf("too many failures") // nolint:goerr113
		},
	}
	return root.ParseAndRun(context.Background(), args[1:])
}

func QuitAfter(timeout *uint64) {
	time.Sleep(time.Duration(*timeout) * time.Second)
	os.Exit(255)
}
