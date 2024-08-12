package jsonx

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

func UnMarshal[T interface{}](content []byte, To T) (T, error) {
	var result = To
	err := json.Unmarshal(content, &result)
	if err != nil {
		logx.Error(err)
		return result, err
	}
	return result, nil
}
