package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) int32 {
	uid := ctx.Value(SessionUserId)
	if uid == nil {
		return 0
	}
	return uid.(int32)
}
