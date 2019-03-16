package response

type Response struct {
	Status int	`json:"status"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(status int, msg string, data interface{}) Response{
	return Response{
		status,
		msg,
		data,
	}
}
