package errors

import (
	"testing"
)

// 자기 다음이
func Test(t *testing.T){
	//t.Errorf("%+v", D())
	t.Logf("%+v", GetRootStackError(wrapped()))
	t.Log("-----------------------------------------------------")
	t.Logf("%+v", GetRootStackError(withStacked()))
	t.Log("-----------------------------------------------------")
	t.Logf("%+v", GetRootStackError(withMessaged()))


}

func Root()error {
	return New("Root")
}

func wrapped() error {
	return Wrap(Root(), "wrapped")
}

func withStacked() error{
	return WithStack(Root())
}

func withMessaged() error{
	return WithMessage(Root(), "withMessaged")
}