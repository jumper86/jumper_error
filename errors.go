package jumper_error

type Error struct{
    Code int32
    Message string
}

func (this *Error) Error() string{
	return this.Message
}

func New(code int32, message string) *Error{
	return &Error{
		Code:    code,
		Message: message,
	}
}