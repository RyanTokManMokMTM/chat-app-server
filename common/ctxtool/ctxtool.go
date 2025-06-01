package ctxtool

import (
	"context"
	"encoding/json"
)

var CTXJWTUserId = "user_id"

func GetUserIDFromCTX(ctx context.Context) uint {
	if jwtUserID, ok := ctx.Value(CTXJWTUserId).(json.Number); ok {
		if id, err := jwtUserID.Int64(); err == nil {
			return uint(id)
		}
	}
	return 0
}
