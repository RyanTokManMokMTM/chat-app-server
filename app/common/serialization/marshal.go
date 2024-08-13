package serialization

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

func Marshal[T interface{}](content T) ([]byte, error) {
	c, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return c, nil
}
