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
	t.Log("-----------------------------------------------------")
	t.Logf("%+v", GetRootStackError(mixed1()))
	t.Log("-----------------------------------------------------")
	t.Logf("%+v", GetRootStackError(mixed2()))
	t.Log("-----------------------------------------------------")
	t.Logf("%+v", GetRootStackError(G()))
}

func Root()error {
	return New("Root")
}

func wrapped() error {
	return Wrap(Wrap(Wrap(Root(), "wrapped1"), "wrapped2"), "wrapped3")
}

func withStacked() error{
	return WithStack(WithStack(WithStack(Root())))
}

func withMessaged() error{
	return WithMessage(WithMessage(WithMessage(Root(), "withMessaged"), "withMessaged2"), "withMessaged3")
}

func mixed1() error{
	return WithMessage(WithStack(Wrap(Root(), "wrap")), "withMessage")
}

func mixed2() error{
	return Wrap(WithStack(WithMessage(Root(), "wrap")), "withMessage")
}


func A() error{
	return New("A")
}

func B() error{
	return WithStack(A())
}

func C() error {
	return Wrap(B(), "")
}

func D() error{
	return WithMessage(C(), "")
}

func E() error{
	return WithMessage(D(), "")
}

func F() error{
	return Wrap(E(), "")
}

func G() error{
	return WithStack(F())
}
