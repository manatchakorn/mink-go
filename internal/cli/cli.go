package cli

import (
	"context"
	"minkgo/internal/cli/mink"
	"minkgo/internal/config"
)

type CLI struct {
	in *mink.Input
}

func New(in *mink.Input) *CLI {
	return &CLI{in: in}
}

func (cli *CLI) Run(ctx context.Context, args ...string) error {
	println("Run command called")
	config.New("mock config path")
	return nil
}
