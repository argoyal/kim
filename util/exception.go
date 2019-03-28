package util

// Block for exceptions
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception raising
type Exception interface{}

// Throw is for panic
func Throw(up Exception) {
	panic(up)
}

// Do will do something
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
