package main

import "fmt"

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	// Clone folder appending the current name to "_clone"
	cloneFolder := &Folder{name: f.name + "_clone"}
	// Iterate throught the children of the folder
	var tempChildren []Inode
	for _, item := range f.children {
		// for each children, clone the item
		copiedItem := item.clone()
		// add the cloned item to the copied folder childrens
		tempChildren = append(tempChildren, copiedItem)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
