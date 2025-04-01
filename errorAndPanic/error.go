package errorAndPanic

import (
	"errors"
	"fmt"
)

type BizError struct {
	Code int
	Msg  string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("Code:%d,Msg:%v", e.Code, e.Msg)
}

func HandleError(err error) {
	if bizErr, ok := err.(*BizError); ok {
		fmt.Println("状态码", bizErr.Code)
		fmt.Println("状态码", bizErr.Msg)

	}
}

var targetErr *BizError

func HandleAsError(err error) {
	if errors.As(err, &targetErr) {
		fmt.Println("状态码!!!!", targetErr.Code)
	}
	if errors.Is(err, targetErr) {
		fmt.Println("状态码!!!!", targetErr.Code)
	}
}
