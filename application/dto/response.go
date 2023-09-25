package dto

const (
	Error   = "error"
	Message = "message"
)

type ResponseBody struct {
	Code        int         `json:"code"`
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(code int, messageType string, message string, data interface{}) ResponseBody {
	return ResponseBody{
		code,
		messageType,
		message,
		data,
	}
}
