package util

import "time"

type Metadata struct {
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	RequestID string    `json:"request_id,omitempty"`
}

type BaseResponse struct {
	Status   int      `json:"status"`
	Success  bool     `json:"success"`
	Message  any      `json:"message,omitempty"`
	Data     any      `json:"data,omitempty"`
	Error    any      `json:"error,omitempty"`
	Metadata Metadata `json:"metadata"`
}
