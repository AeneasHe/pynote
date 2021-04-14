package server

import (
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	livereload "github.com/mattn/echo-livereload"
)

func ConfigEcho(e *echo.Echo, server *Server) {
	// echo 配置模版目录
	t := NewTemplate(server.config.TemplatePath)
	e.Renderer = &t

	// echo 是否开启debug
	e.Debug = server.config.Debug

	e.Static("/static", server.config.StaticPath)
	//e.Use(middleware.Static(server.config.StaticPath))

	// 热加载
	lrconfig := livereload.LiveReloadConfig{
		Skipper: middleware.DefaultSkipper,
		Name:    os.Args[0],
		Dir:     server.config.StaticPath,
	}
	lr := livereload.LiveReloadWithConfig(lrconfig)
	e.Use(lr)

	// echo 自定义Context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})

	//  echo 日志定义
	timelocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timelocal

	loggerConfig := middleware.LoggerConfig{
		// 跳过的日志，行情查询的相关请求不记录
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/kline/query" || strings.Contains(c.Path(), "api/v1/private") ||
				strings.Contains(c.Path(), "api/v1/trades") ||
				strings.Contains(c.Path(), "api/v1/depth") {
				return true
			}
			return false
		},
		// 日志格式
		Format: `{"time":"${time_rfc3339_nano}","file":"${long_file}","line":"${line}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           os.Stdout,
	}
	e.Use(middleware.LoggerWithConfig(loggerConfig))

	if !e.Debug {
		e.Use(middleware.Recover())
	}

	// echo 跨域
	e.Use(middleware.CORS())
	// echo  使用xss protect
	e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))

}
