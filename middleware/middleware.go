package middleware

import "github.com/valyala/fasthttp"

func Cors(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		origin := string(c.Request.Header.Peek("Origin")) //请求头部
		if origin != "" {
			// CORS
			c.Response.Header.Set(fasthttp.HeaderAccessControlAllowCredentials, "true")
			c.Response.Header.Set(fasthttp.HeaderAccessControlAllowHeaders, "content-type, Authorization, Content-Length, X-CSRF-Token, Token, session")
			c.Response.Header.Set(fasthttp.HeaderAccessControlAllowMethods, "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, origin)
			c.Response.Header.Set(fasthttp.HeaderAccessControlExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Response.Header.Set(fasthttp.HeaderAccessControlMaxAge, "172800")
		}
		// 如果是预检请求 (OPTIONS)，直接返回，不继续执行下一个处理器
		if string(c.Request.Header.Method()) == fasthttp.MethodOptions {
			c.Response.SetStatusCode(fasthttp.StatusOK)
			return
		}

		// TODO 加签验证

		// 继续执行下一个处理器
		next(c)
	}
}
