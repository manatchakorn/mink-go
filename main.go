package main

import (
	"context"
	"errors"
	"fmt"
	"minkgo/internal/cli"
	"minkgo/internal/cli/mink"
	"os"
)

// mink-go entry point here
func main() {
	if err := run(); err != nil {
		fmt.Printf("[ERROR]: run program with error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {

	// run mink-go cli with default configuration
	cli := cli.New(&mink.Input{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Env:    os.Environ(),
	})

	if err := cli.Run(context.Background(), os.Args[1:]...); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil

		}
		return err
	}

	return nil

}
