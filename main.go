package main

import (
	"fmt"
	"os"

	"github.com/hugoivankm/gator/cli"
	"github.com/hugoivankm/gator/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments to run command")
		os.Exit(1)
	}

	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("Error reading configuration file %v\n", err)
		os.Exit(1)
		return
	}

	cmds := cli.NewCommands()

	cmds.Register("login", cli.HandlerLogin)

	s := &cli.State{
		Cfg: cfg,
	}

	cmd := cli.Command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	fmt.Print(">>> ")
	err = cmds.Run(s, cmd)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", cfg)
}
