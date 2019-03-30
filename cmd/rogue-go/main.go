package main

import (
	"runtime"

	rogue "github.com/lostinblue/rogue-go"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	rogue.RunAsRoot()
	rogue.PrintBanner()
	rogue.NewSignalHandler()
	hackEngine := rogue.NewEngine()
	// TODO: Add key handling
	hackEngine.Render()
}

// mmap will be useful later
//data, err = syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
//if err != nil {
//	panic("Unable to mmap file")
//}
