package pkgerror

import (
	"fmt"
)

type MyError struct {
	Raw       error
	ErrorCode string
	Message   string
}

func (e MyError) Error() string {
	if e.Raw != nil {
		return fmt.Sprintf("%s:%s", e.Raw, e.ErrorCode)
	}

	return e.Message
}
