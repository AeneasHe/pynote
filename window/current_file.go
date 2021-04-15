package window

import "sync"

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
