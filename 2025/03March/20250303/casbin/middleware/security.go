package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 根据环境动态设置允许的源（解决CORS）
		allowedOrigin := "*" // 生产环境域名

		// 内容安全策略（CSP）：最重要的安全头
		// 作用：限制页面可以加载哪些资源
		c.Header("Content-Security-Policy",
			fmt.Sprintf(
				"default-src 'self'; "+ // 默认只允许同源
					"script-src 'self' 'unsafe-inline' https://trusted.cdn.com; "+ // 允许的JS来源
					"style-src 'self' 'unsafe-inline'; "+ // 允许的CSS来源
					"img-src 'self' data:; "+ // 允许的图片来源
					"connect-src 'self' %s; "+ // 允许的API请求来源
					"report-uri /csp-report", // CSP违规报告地址
				allowedOrigin,
			))

		// 其他关键安全头：
		c.Header("X-Content-Type-Options", "nosniff") // 禁止MIME嗅探
		c.Header("X-Frame-Options", "DENY")           // 禁止页面被嵌入iframe
		c.Header("X-XSS-Protection", "1; mode=block") // 启用浏览器XSS过滤

		// CORS相关设置：
		c.Header("Access-Control-Allow-Origin", allowedOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")

		c.Next() // 继续处理后续中间件/路由
	}
}

// XSSFilterMiddleware 过滤用户输入中的XSS代码
func XSSFilterMiddleware() gin.HandlerFunc {
	// 初始化HTML过滤策略：
	// UGCPolicy 允许常见安全标签（如<b>/<i>/<a>等）
	policy := bluemonday.UGCPolicy()
	// 按需放宽规则：允许div标签的class属性
	policy.AllowAttrs("class").OnElements("div", "span")

	return func(c *gin.Context) {
		// 过滤GET查询参数
		queryParams := c.Request.URL.Query()
		for key, values := range queryParams {
			cleaned := make([]string, len(values))
			for i, v := range values {
				cleaned[i] = policy.Sanitize(v) // 对每个值进行过滤
			}
			queryParams[key] = cleaned // 替换为安全值
		}
		c.Request.URL.RawQuery = queryParams.Encode() // 更新URL参数

		// 过滤POST表单数据
		if c.Request.Method == "POST" {
			// 解析表单数据（最大内存限制默认32MB）
			if err := c.Request.ParseForm(); err == nil {
				formParams := c.Request.PostForm
				for key, values := range formParams {
					cleaned := make([]string, len(values))
					for i, v := range values {
						cleaned[i] = policy.Sanitize(v)
					}
					formParams[key] = cleaned // 替换过滤后的值
				}
			}
		}
		c.Next() // 继续处理
	}
}
