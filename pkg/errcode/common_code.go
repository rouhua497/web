package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "内部服务错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "坚权失败，找不到对应的appkey")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeOut  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
)
