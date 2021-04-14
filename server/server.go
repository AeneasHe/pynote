package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/jinzhu/configor"
	"github.com/labstack/echo/v4"

	"pynote/asset"
	"pynote/utils"
)

const SiteID = "X-SITE-ID"
const DeviceInfo = "X-DEVICE-INFO"
const Lang = "lang"
const ExternalID = "EXTERNAL-AUTH-ID"

type Server struct {
	e      *echo.Echo
	config *EchoServerConfig
}

type EchoServerConfig struct {
	Host           string `default:""`
	Port           string `default:"80"`
	Debug          bool
	Sync           bool
	TemplatePath   string `required:"true"`
	StaticPath     string `required:"true"`
	RedirectHost   string `required:"true"`
	ServerHost     string `required:"true"`
	Secret         string `required:"true"`
	Redis          *utils.RedisConfig
	WhiteListRobot []uint
}

func NewServer(configPath string) *Server {
	// 载入配置
	server := new(Server)
	server.config = new(EchoServerConfig)
	err := configor.Load(server.config, configPath)
	if err != nil {
		panic(err)
	}

	// echo 服务
	e := echo.New()

	// echo 基础配置
	ConfigEcho(e, server)

	userAuth := e.Group("user")
	userAuth.GET("", server.getUserInfo)

	e.GET("/", server.index)
	e.GET("/show", server.show)

	server.e = e

	return server

}

func (s *Server) Run() {
	s.e.Logger.Fatal(s.e.Start(s.config.Host + ":" + s.config.Port))
}

func (s *Server) Close() {
	utils.RedisConnClose()
}

func Start() {
	config := "./config.json"
	s := NewServer(config)
	s.Run()
}

func StartServer(addr string) string {
	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, http.FileServer(asset.FS))
	url := fmt.Sprintf("http://%s", ln.Addr())
	return url
}
