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
	RootPath       string `required:"true"`
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
	log.Println("======>2", configPath)
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
	// 启动服务器
	addr := s.config.Host + ":" + s.config.Port
	s.e.Logger.Fatal(s.e.Start(addr))
}

func (s *Server) Close() {
	utils.RedisConnClose()
}

func StartFromTemplate(configPath string) {
	s := NewServer(configPath)
	s.Run()
}

func StartFromAsset() string {
	addr := "127.0.0.1:8099"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, http.FileServer(asset.FS))
	url := fmt.Sprintf("http://%s", ln.Addr())
	log.Println("-------->", url)
	return url
}
