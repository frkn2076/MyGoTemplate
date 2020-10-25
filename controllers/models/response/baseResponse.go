package response


type BaseResponse struct {
	IsSuccess bool  `json:"isSuccess"`
	Message    string  `json:"message"`
}

