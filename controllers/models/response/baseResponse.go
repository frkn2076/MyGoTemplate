package response

var SingletonBaseResponseInstance *BaseResponse

func init() {
    instance := BaseResponse{
			IsSuccess: false,
			Message: "asdasd",
	}
	SingletonBaseResponseInstance = &instance
}


//#region Models

type BaseResponse struct {
	IsSuccess bool  `json:"isSuccess"`
	Message    string  `json:"message"`
}

//#endregion
