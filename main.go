package main

import (
	"flag"
	"os"
	"time"
)

type Remover interface {
	Remove()
}

type Folder struct {
	Name string
}

var folderName = flag.String("folder", "", "folder to be removed")

func (f Folder) Remove() {
	os.RemoveAll(f.Name)
}

func RemoveThisAndChildren(r Remover) {
	r.Remove()
}

func Init() {
	flag.Parse()

	t := time.Now()

	f := Folder{*folderName}

	RemoveThisAndChildren(f)
}

func main() {
	go Init()
}
