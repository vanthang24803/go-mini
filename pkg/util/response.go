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
	Message  string   `json:"message"`
	Data     any      `json:"data,omitempty"`
	Error    string   `json:"error,omitempty"`
	Metadata Metadata `json:"metadata"`
}
