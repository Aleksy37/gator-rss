package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("error reseting the users table: %v", err)
	}
	fmt.Println("users table was successfully reset")
	return nil
}