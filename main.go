package main

import (
	"flag"
	"fmt"
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
	fmt.Printf("Starting removal of: %v and it's children\n", f.Name)

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

	fmt.Printf("Finished in: %v", time.Since(t))
}

func main() {
	go Init()
}
