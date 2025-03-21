package error

import (
	"errors"
	"fmt"
)

type BizError struct {
	Code int
	Msg  string
}

func (e *BizError) Error() string {
	return e.Msg
}

func HandleError(err error) {
	if bizErr, ok := err.(*BizError); ok {
		fmt.Println("状态码", bizErr.Code)
	}
}

var targetErr *BizError

func HandleAsError(err error) {
	if errors.As(err, &targetErr) {
		fmt.Println("状态码", targetErr.Code)
	}
}
