package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mars-projects/mars/common/utils"
)

type Message struct {
	*Request
	context.Context
}

func (e *Message) UnMarshal(data interface{}) error {
	fmt.Println(e.GetData())
	return json.Unmarshal(e.Data, data)
}

func (e Message) getStringFromContext(key string) string {
	v := e.Value(key)
	switch v.(type) {
	case string:
		return v.(string)
	}
	return ""
}

func (e Message) GetUserId() int {
	s := e.getStringFromContext("user_id")
	toInt, err := utils.StringToInt(s)
	if err != nil {
		return 0
	}
	return toInt
}
