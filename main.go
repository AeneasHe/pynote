package main

import (
	"os"
	"os/signal"
	"runtime"

	"pynote/server"
	"pynote/utils"
	"pynote/window"

	"log"
)

func startEventLoop(mainWindow window.MainWindow) {
	// 交互信号处理
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
	case <-mainWindow.UI.Done():
	}

	log.Println("exiting...")
}

func main() {

	// 日志初始化
	utils.InitLog()

	// 获取配置文件
	configPath := utils.GetConfigPath()
	log.Println("配置文件", configPath)

	// 初始化
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	// 服务器初始化
	// go server.StartFromTemplate(configPath) // 从template中读取静态文件
	go server.StartFromAsset() // 从嵌入式asset读取静态文件
	// 开发阶段也可以用python独立脚本启动，有热加载，方便实时更新： python app.py

	// webview窗口初始化
	mainWindow := window.NewMainWindow(args)
	defer mainWindow.UI.Close()

	// 绑定监听方法
	//mainWindow.StartCounter()
	mainWindow.AddPathWalk()

	// 窗口加载页面
	mainWindow.Load(configPath)

	// 执行js
	mainWindow.Dojs()

	startEventLoop(mainWindow)

}
