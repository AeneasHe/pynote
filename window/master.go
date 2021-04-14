package window

import (
	"fmt"
	"log"
	"pynote/counter"
	"sync"

	"pynote/fs"

	"github.com/zserge/lorca"
)

type MainWindow struct {
	UI lorca.UI
}

type CurrentFile struct {
	sync.Mutex
	Data string
}

func (f *CurrentFile) Update(data string) {
	f.Lock()
	defer f.Unlock()
	f.Data = data
}
func (f *CurrentFile) Read() string {
	f.Lock()
	defer f.Unlock()
	return f.Data
}

var cf = CurrentFile{}

func NewMainWindow(args []string) MainWindow {

	mainWindow := MainWindow{}

	// 窗口初始化
	ui, err := lorca.New("", "", 1200, 600, args...)
	if err != nil {
		log.Fatal(err)
	}

	mainWindow.UI = ui

	// 绑定方法
	// A simple way to know when UI is ready (uses body.onload event in JS)
	mainWindow.UI.Bind("start", func() {
		log.Println("UI is ready")
	})

	return mainWindow

}

func (mainWindow *MainWindow) Load(url string) {
	// 窗口加载页面
	fmt.Println(url)
	mainWindow.UI.Load(url)
}

func (mainWindow *MainWindow) Dojs() {
	// 窗口执行js程序
	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	mainWindow.UI.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)
}

func (mainWindow *MainWindow) StartCounter() {
	// Create and bind Go object to the UI
	c := &counter.Counter{}
	mainWindow.UI.Bind("counterAdd", c.Add)
	mainWindow.UI.Bind("counterValue", c.Value)
}

func (mainWindow *MainWindow) AddPathWalk() {
	mainWindow.UI.Bind("showPath", fs.Walk)
	mainWindow.UI.Bind("showFile", fs.OpenFile)
	mainWindow.UI.Bind("openWindow", mainWindow.openWindow)
	mainWindow.UI.Bind("currentFile", cf.Read)
}

func (mainWindow *MainWindow) openWindow(data interface{}) {
	var filename = data.(string)
	if filename != "" {
		str := fs.OpenFile(filename)
		cf.Update(str)
	}
	//mainWindow.Load("http://127.0.0.1/show")
}
