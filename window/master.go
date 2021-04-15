package window

import (
	"log"
	"pynote/counter"
	"pynote/server"

	"pynote/fs"

	"github.com/jinzhu/configor"
	"github.com/zserge/lorca"
)

type MainWindow struct {
	UI lorca.UI
}

func NewMainWindow(args []string) MainWindow {

	mainWindow := MainWindow{}

	// 窗口初始化
	ui, err := lorca.New("", "", 1200, 600, args...)
	if err != nil {
		log.Fatal(err)
	}

	mainWindow.UI = ui

	// 绑定方法
	mainWindow.UI.Bind("start", func() {
		log.Println("UI is ready")
	})

	return mainWindow

}

// 加载配置文件，然后加载目标路径
func (mainWindow *MainWindow) Load(configPath string) {
	config := new(server.EchoServerConfig)
	err := configor.Load(config, configPath)
	if err != nil {
		println("配置文件读取错误")
		panic(err)
	}

	fs.RootPath = config.RootPath

	url := "http://" + config.Host + ":" + config.Port
	log.Println("=======>window url:", url)
	// 窗口加载页面
	//url := "http://127.0.0.1:8099"

	mainWindow.UI.Load(url)
}

// 执行js
func (mainWindow *MainWindow) Dojs() {
	// 窗口执行js程序
	mainWindow.UI.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)
}

// 开始计时器
func (mainWindow *MainWindow) StartCounter() {
	c := &counter.Counter{}
	mainWindow.UI.Bind("counterAdd", c.Add)
	mainWindow.UI.Bind("counterValue", c.Value)
}

// 绑定路径相关的方法
func (mainWindow *MainWindow) AddPathWalk() {
	// 显示路径
	mainWindow.UI.Bind("showPath", fs.ShowPath)

	// 打开
	mainWindow.UI.Bind("openPath", mainWindow.openPath)

	// 读取当前文件的内容
	mainWindow.UI.Bind("currentFile", cf.Read)

}

func (mainWindow *MainWindow) openPath(data interface{}, pathType string) {
	if pathType == "path-file" {
		var filename = data.(string)
		if filename != "" {
			log.Println("=======>filename:", filename)
			str := fs.ReadFile(filename)
			cf.Update(str)
		}
	}

	if pathType == "path-folder" {
		var foldername = data.(string)
		if foldername != "" {
			log.Println("=======>foldername:", foldername)
			fs.ShowPath(foldername, "all")

		}
	}
}
