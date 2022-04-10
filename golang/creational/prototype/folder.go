package main

import "fmt"

type folder struct {
	childrens []inode
	name      string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.childrens {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildrens []inode
	for _, child := range f.childrens {
		copy := child.clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}
