package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

// 中间件 方便请求前请求后执行一些逻辑 （比如：权限校验
// clouwego 提供的 middleware 看 hertz-contrib
func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		begin := time.Now()

		err = next(ctx, req, resp)

		fmt.Println(time.Since(begin))

		return err
	}
}
