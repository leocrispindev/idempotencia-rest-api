package model

import "encoding/json"

type Status int

const (
	IN_PROCESS       Status = 1
	PROCESSED        Status = 2
	ERROR_ON_PROCESS Status = 3
)

type CacheMessage struct {
	ProcessStatus Status `json:"status"`
	Message       string `json:"message"`
}

func (c *CacheMessage) InProccess() bool {
	return c.ProcessStatus == PROCESSED || c.ProcessStatus == IN_PROCESS
}

func (c *CacheMessage) StatusError() bool {
	return c.ProcessStatus == PROCESSED || c.ProcessStatus == IN_PROCESS
}

func (c *CacheMessage) ToJson() string {
	data, err := json.Marshal(c)

	if err != nil {
		return ""
	}

	return string(data)

}
