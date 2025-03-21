package error

import "testing"

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
