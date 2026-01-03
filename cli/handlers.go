package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hugoivankm/gator/internal/config"
)

type Command struct {
	Name      string
	Arguments []string
}

type State struct {
	Cfg *config.Config
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("login command expects exactly one user as argument.")
	}

	var user string = cmd.Arguments[0]
	trimmedUser := strings.TrimSpace(user)
	err := s.Cfg.SetUser(trimmedUser)
	if err != nil {
		return errors.New("login command expects a non empty name.")
	}
	fmt.Printf("%s has logged in.\n", trimmedUser)
	return nil
}
