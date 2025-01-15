package response

func OK() Response {
	return Response{
		Code: Success,
		Info: "OK",
	}
}

func Bind(err error) Response {
	return Response{
		Code: BindError,
		Info: "BindError" + err.Error(),
	}
}

func Internal(err error) Response {
	return Response{
		Code: InternalError,
		Info: "InternalError" + err.Error(),
	}
}

func Register(err error) Response {
	return Response{
		Code: RegisterError,
		Info: err.Error(),
	}
}

func Password() Response {
	return Response{
		Code: PasswordErr,
		Info: PasswordError.Error(),
	}
}

func OkWithData(data interface{}) ResponseWithData {
	return ResponseWithData{
		Code: Success,
		Info: "OK",
		Data: data,
	}
}

func TokenError() Response {
	return Response{
		Code: TokenErr,
		Info: "unauthorized",
	}
}
