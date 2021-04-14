package server

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DeviceInfoResp struct {
	Endpoint   string
	AppVersion string
	Platform   string
	OSVersion  string
	DeviceID   string
}

// 自定义context：包装echo.Context
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) SiteID() uint {
	siteIDStr := c.Request().Header.Get(SiteID)
	if siteIDStr == "" {
		siteIDStr = "1"
	}
	siteID, err := strconv.ParseUint(siteIDStr, 10, 64)
	if err != nil {
		return 1
	}
	return uint(siteID)
}

func (c *CustomContext) Lang() string {
	lang := c.Request().Header.Get(Lang)
	if lang == "" {
		lang = "zh-hans"
	}
	lang = strings.ToLower(lang)
	return lang
}

// 是否存在
func (c *CustomContext) DeviceInfo() (*DeviceInfoResp, bool) {
	deviceInfoStr := c.Request().Header.Get(DeviceInfo)
	if deviceInfoStr == "" {
		return nil, false
	}
	parts := strings.Split(deviceInfoStr, "_")
	if len(parts) != 5 {
		return nil, false
	}
	di := new(DeviceInfoResp)
	di.Endpoint = parts[0]
	di.AppVersion = parts[1]
	di.Platform = parts[2]
	di.OSVersion = parts[3]
	if strings.ToLower(di.Endpoint) == "web" {
		result, _ := base64.StdEncoding.DecodeString(parts[4])
		di.DeviceID = string(result)
	} else {
		di.DeviceID = parts[4]
	}
	return di, true
}
