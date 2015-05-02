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

func (f Folder) Remove() {
	fmt.Printf("Starting removal of: %v and it's children\n", f.Name)

	os.RemoveAll(f.Name)
}

func RemoveThisAndChildren(r Remover) {
	r.Remove()
}

func getFlagVal() *string {

	fd := flag.String("folder", "", "folder to be removed")

	flag.Parse()

	return fd
}

func main() {

	t := time.Now()

	f := Folder{*getFlagVal()}

	RemoveThisAndChildren(f)

	fmt.Printf("Finished in: %v", time.Since(t))
}
