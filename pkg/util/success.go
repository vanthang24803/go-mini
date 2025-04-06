package util

import "time"

func SuccessResponse(message string, data any) BaseResponse {
	return BaseResponse{
		Status:  200,
		Success: true,
		Message: message,
		Data:    data,
		Metadata: Metadata{
			Timestamp: time.Now().UTC(),
			Version:   "v1.0",
		},
	}
}
