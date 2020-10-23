package reponse

type BaseResponse struct {
	IsSuccess bool  `json:"isSuccess"`
	Message    string  `json:"message"`
}
