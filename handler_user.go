package main

import "fmt"

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Unexpected number of arguments, expected: 1, got: %d", len(cmd.Args))
	}

	err := s.cfg.SetUser(cmd.Args[0])

	if err != nil {
		return err
	}

	fmt.Println("User has been set successfully")
	return nil
}
