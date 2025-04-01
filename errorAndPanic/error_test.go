package errorAndPanic

import (
	"errors"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	b := BizError{
		Code: 200,
		Msg:  "nihao",
	}
	c := BizError{
		Code: 500,
		Msg:  b.Error(),
	}
	HandleError(&b)
	HandleAsError(&c)
}

func TestFormatWFunction(t *testing.T) {
	b := BizError{
		Code: 400,
		Msg:  "errorMsg",
	}
	berr := errorS(&b)
	fmt.Println("errorW", berr)
	HandleAsError(berr)
	HandleError(berr)
	fmt.Println("unWrap")
	unerr := errors.Unwrap(berr)
	HandleError(unerr)

}

func errorS(err error) error {
	return fmt.Errorf("this is error,%w", err)
}
