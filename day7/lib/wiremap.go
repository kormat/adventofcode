package day7lib

import (
	"fmt"
	"sort"
	"sync"
)

type Wire struct {
	Name  string
	Done  chan bool
	Value uint16
}

type WireMap struct {
	sync.Mutex
	M     map[string]*Wire
	names []string
}

func (wm *WireMap) add(name string) *Wire {
	wm.Lock()
	elem, ok := wm.M[name]
	if !ok {
		elem = &Wire{Name: name, Done: make(chan bool)}
		wm.M[name] = elem
		wm.names = append(wm.names, name)
	}
	wm.Unlock()
	return elem
}

func (wm *WireMap) print() {
	sort.Strings(wm.names)
	for _, name := range wm.names {
		fmt.Printf("  %s: %d\n", name, wm.M[name].Value)
	}
}

func (w *Wire) read() uint16 {
	<-w.Done
	return w.Value
}
