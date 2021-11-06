package errors

// *withStack 이나 *withMessage 는 Cause()를 구현했다.
// *fundamental 은 Cause()를 구현하지 않았다.
type causer interface {
		Cause() error
}

func GetRootStackError(err error) error{
	// err이 *withStack 이나 *withMessage인 경우는
	// 그 녀석들이 RootStackError인지 확인해보고 RootStackError 이면 리턴
	// 그 녀석들 안에 RootStackError가 있으면 재귀적으로 그것을 호출.

	// err이 *withStack 이나 *withMessage가 아니면 무조건 걔네를 리턴.
	// *fundamental 이나 다른 error(스택이 지원되지 않는) 타입들
	if w, ok := err.(*withStack); ok{
		if IsRootStackError(w) {
			return err
		} else {
			return GetRootStackError(w.Cause())
		}
	} else if w, ok := err.(*withMessage); ok {
		if IsRootStackError(w) {
			return err
		} else {
			return GetRootStackError(w.Cause())
		}
	} else{
		return err
	}
}


// err causer 는 *withStack이나 *withMessage
func IsRootStackError(err causer) bool{
	// err의 원인이 *withStack, *withMessage, *fundamental 중 하나라면
	// err이 root stack을 가진 error가 아니다.
	if _, ok := err.Cause().(*withStack); ok{
		return false
	} else if _, ok := err.Cause().(*withMessage); ok{
		return false
	} else if _, ok := err.Cause().(*fundamental); ok{
		return false
	}

	return true
}
