package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"pynote/server"
	"pynote/window"
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

	// 初始化
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	// 窗口初始化
	mainWindow := window.NewMainWindow(args)
	defer mainWindow.UI.Close()

	mainWindow.StartCounter()
	mainWindow.AddPathWalk()

	// 服务器初始化
	//addr := "127.0.0.1:0"
	//url := server.StartServer(addr)
	//defer ln.Close()
	go server.Start()
	url := "http://127.0.0.1"

	// window2 := window.NewMainWindow(args)
	// defer window2.UI.Close()
	// window2.Load("http://127.0.0.1/show")

	// 窗口加载页面
	mainWindow.Load(url)

	// 执行js
	mainWindow.Dojs()

	startEventLoop(mainWindow)

}
