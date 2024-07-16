package main

type Inode interface {
	print(string)
	clone() Inode
}

type File struct {
	name string
}

type Folder struct {
	children []Inode
	name     string
}
