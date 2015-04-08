package helpers

import (
	// "errors"
	"fmt"
)

// MySQLError is an error type which represents a single MySQL error
type MSGError struct {
	Number  uint16
	Message string
}

func (me *MSGError) Err() string {
	return fmt.Sprintf("Error %d: %s", me.Number, me.Message)
}
