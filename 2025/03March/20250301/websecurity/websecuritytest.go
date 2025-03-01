// 定义包名（每个Go程序必须从main包开始）
package main

// 导入依赖库
import (
	"bytes"          // 字节缓冲操作
	"fmt"            // 格式化输出
	"html/template"  // 安全HTML模板引擎（关键防XSS）
	"io"             // I/O接口
	"mime/multipart" // 处理文件上传
	"os"             // 操作系统功能（环境变量等）
	"strings"        // 字符串处理

	"github.com/disintegration/imaging"  // 图片处理库（清除元数据）
	"github.com/gin-gonic/gin"           // Web框架
	"github.com/microcosm-cc/bluemonday" // HTML过滤库（防XSS核心）
	"image"                              // 图片解码接口
)

// 主函数：程序入口
func main() {
	// 初始化Gin引擎（默认包含Logger和Recovery中间件）
	r := gin.Default()

	// ==================== 注册安全中间件 ====================
	// 中间件按注册顺序执行，先处理安全头再处理数据

	// 1. 安全响应头中间件（CSP/XSS防护等）
	r.Use(SecurityHeadersMiddleware())

	// 2. XSS输入过滤中间件（清理GET/POST参数）
	r.Use(XSSFilterMiddleware())

	// 3. 生产环境强制HTTPS（防止中间人攻击）
	if os.Getenv("ENV") == "production" { // 通过环境变量判断环境
		r.Use(RequireHTTPSMiddleware())
	}

	// ==================== 定义业务路由 ====================

	// 首页路由：演示安全模板渲染
	r.GET("/", func(c *gin.Context) {
		// 模拟用户输入（包含恶意脚本）
		userInput := "<script>alert(1)</script>"

		// 使用gin.H传递数据到模板：
		// - safeContent: 手动标记为安全的HTML（谨慎使用！）
		// - escapedContent: 自动转义的内容（默认安全）
		c.HTML(200, "index.html", gin.H{
			"safeContent":    template.HTML("<b>安全内容</b>"),
			"escapedContent": userInput,
		})
	})

	// 评论提交路由：演示输入过滤
	r.POST("/comment", func(c *gin.Context) {
		// 获取原始用户输入
		rawComment := c.PostForm("comment")

		// 使用bluemonday进行XSS过滤：
		// - UGCPolicy() 允许常见安全标签（如<b>/<i>）
		// - Sanitize() 移除危险内容
		safeComment := bluemonday.UGCPolicy().Sanitize(rawComment)

		// 返回处理后的安全内容
		c.JSON(200, gin.H{
			"status":  "success",
			"comment": safeComment,
		})
	})

	// 文件上传路由：演示安全文件处理
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中获取上传的文件
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "文件无效"})
			return
		}
		defer file.Close() // 确保关闭文件句柄

		// 验证文件扩展名是否在白名单内
		if !isFileTypeAllowed(header.Filename) {
			c.AbortWithStatusJSON(400, gin.H{"error": "仅支持 JPG/PNG"})
			return
		}

		// 处理图片文件（清除元数据等敏感信息）
		// 此处应调用 ProcessImage 函数处理（示例代码中未完全实现）
		_, err = ProcessImage(file)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "文件处理失败"})
			return
		}

		c.JSON(200, gin.H{"status": "文件已安全处理"})
	})

	// ==================== 模板配置 ====================
	// 加载HTML模板文件（自动转义功能在此生效）
	r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))

	// ==================== 启动服务 ====================
	// 从环境变量读取端口，默认8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port) // 启动HTTP服务
}

// ==================== 中间件实现部分 ====================

// SecurityHeadersMiddleware 设置安全相关的HTTP头
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

// RequireHTTPSMiddleware 强制跳转HTTPS（生产环境用）
func RequireHTTPSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过X-Forwarded-Proto判断原始协议（适用于反向代理场景）
		if c.Request.Header.Get("X-Forwarded-Proto") != "https" {
			// 构造HTTPS目标URL
			target := "https://" + c.Request.Host + c.Request.RequestURI
			// 301永久重定向
			c.Redirect(301, target)
			c.Abort() // 终止后续处理
			return
		}
		c.Next()
	}
}

// ==================== 文件处理函数 ====================

// isFileTypeAllowed 检查文件扩展名是否合法
func isFileTypeAllowed(filename string) bool {
	// 允许的文件类型白名单
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	// 提取文件扩展名并转为小写
	ext := strings.ToLower(filename[strings.LastIndex(filename, "."):])
	return allowed[ext]
}

// ProcessImage 处理上传的图片文件（清除元数据）
func ProcessImage(file multipart.File) ([]byte, error) {
	// 解码图片文件（自动验证文件格式）
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("图片解码失败: %v", err)
	}

	// 重置文件指针（如果后续需要再次读取）
	if seeker, ok := file.(io.Seeker); ok {
		seeker.Seek(0, io.SeekStart)
	}

	// 使用imaging库将图片转为JPEG格式（清除元数据）
	var buf bytes.Buffer
	if err := imaging.Encode(&buf, img, imaging.JPEG); err != nil {
		return nil, fmt.Errorf("图片编码失败: %v", err)
	}
	return buf.Bytes(), nil
}
