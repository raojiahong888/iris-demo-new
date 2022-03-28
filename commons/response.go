package commons

type ResponseCode int

const (
	OK ResponseCode = 0
	UnKnowError ResponseCode = -1
	ParamsFormatError ResponseCode = 3
	ServerError ResponseCode = 4
)

var ErrorText = map[ResponseCode]string{
	OK: "success",
	UnKnowError: "unknown error",
	ParamsFormatError: "Failed to get parameters",
	ServerError: "Server error",
}

type ResponseBean struct {
	code ResponseCode `json:"code"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(message string, data ...interface{}) *ResponseBean {
	if data != nil {
		return &ResponseBean{0,message, data[0]}
	}
	return &ResponseBean{0,message,nil}
}

func ResponseError(code ResponseCode, message ...interface{}) *ResponseBean {
	errMsg := GetCodeAndMsg(code)
	if message != nil {
		errMsg = message[0].(string)
	}
	return &ResponseBean{code, errMsg, nil}
}

func GetCodeAndMsg(code ResponseCode) string {
	value, ok := ErrorText[code]
	if ok {
		return value
	}
	return ""
}