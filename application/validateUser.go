package application

import "fmt"

func ValidateUser(user UserBody) error {
	if user.FirstName == "" || user.LastName == "" {
		return fmt.Errorf("first_name and last_name are required")
	}

	if len(user.FirstName) < 2 || len(user.FirstName) > 20 {
		return fmt.Errorf("first_name must be between 2 and 20 characters")
	}

	if len(user.LastName) < 2 || len(user.LastName) > 20 {
		return fmt.Errorf("last_name must be between 2 and 20 characters")
	}

	if user.Biography == "" {
		return fmt.Errorf("biography is required")
	}

	if len(user.Biography) < 20 || len(user.Biography) > 450 {
		return fmt.Errorf("biography must be between 20 and 450 characters")
	}

	return nil
}
