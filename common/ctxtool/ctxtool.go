package ctxtool

import (
	"context"
	"encoding/json"
)

var CTXJWTUserID = "user_id"

func GetUserIDFromCTX(ctx context.Context) uint {
	if jwtUserId, ok := ctx.Value(CTXJWTUserID).(json.Number); ok {
		if id, err := jwtUserId.Int64(); err == nil {
			return uint(id)
		}
	}
	return 0
}
