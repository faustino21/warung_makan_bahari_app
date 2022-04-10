package resoponse

var (
	responseMessage = "00"
	responseFailed  = "XX"
	statusSuccess   = "Success"
	statusFailed    = "Failed"
)

type SuccessMessage struct {
	HttpResponse int         `json:"http_response"`
	RespCode     string      `json:"resp_code"`
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

type FailedMessage struct {
	HttpResponse int    `json:"http_response"`
	RespCode     string `json:"resp_code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}

func NewSuccessMessage(message string, data interface{}) *SuccessMessage {
	return &SuccessMessage{
		RespCode: responseMessage,
		Status:   statusSuccess,
		Message:  message,
		Data:     data,
	}
}

func NewFailedMessage(message string) *FailedMessage {
	return &FailedMessage{
		RespCode: responseFailed,
		Status:   statusFailed,
		Message:  message,
	}
}
