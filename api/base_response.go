package api

var Success *BaseResponse

func init() {
	Success = &BaseResponse{IsSuccess: true}
}

type BaseResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
}
