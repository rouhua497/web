package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

//我们需要将设置了超时的 c.Request.Context 方法传递进去，在验证时可以调短默认超时时间来进行调试
//把父级的上下文信息（ctx）不断地传递下去，那么在统计超时控制的中间件中所设置的超时时间，其实是针对整条链路的。
//如果需要单独调整某条链路的超时时间，那么只需调用context.WithTimeout等方法对父级 ctx 进行设置，然后取得子级 ctx，再进行新的传递即可

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
