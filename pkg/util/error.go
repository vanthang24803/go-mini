package util

import "time"

func ErrorResponse(err any) BaseResponse {
	return BaseResponse{
		Status:  400,
		Success: false,
		Error:   err,
		Metadata: Metadata{
			Timestamp: time.Now().UTC(),
			Version:   "v1.0",
		},
	}
}
