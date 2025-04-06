package util

import "time"

func ErrorResponse(status int, message string, err string) BaseResponse {
	return BaseResponse{
		Status:  status,
		Success: false,
		Message: message,
		Error:   err,
		Metadata: Metadata{
			Timestamp: time.Now().UTC(),
			Version:   "v1.0",
		},
	}
}

func BadRequestError(message string) BaseResponse {
	return ErrorResponse(400, "Bad Request", message)
}

func NotFoundError(message string) BaseResponse {
	return ErrorResponse(404, "Not Found", message)
}

func InternalServerError(message string) BaseResponse {
	return ErrorResponse(500, "Internal Server Error", message)
}
