package server

import "strings"

type treeNode struct {
	name       string
	children   []*treeNode
	routerName string
	isEnd      bool
}

func (t *treeNode) Put(path string) {
	root := t
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name {
				isMatch = true
				t = node
				break
			}
		}
		if !isMatch {
			node := &treeNode{
				name:       name,
				children:   make([]*treeNode, 0),
				routerName: t.routerName + "/" + name,
				isEnd:      true,
			}
			children = append(children, node)
			t.children = children
			t.isEnd = false
			t = node
		}
	}
	t = root
}

func (t *treeNode) Get(path string) *treeNode {
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name ||
				node.name == "*" ||
				strings.Contains(node.name, ":") {
				t = node
				if index == len(strs)-1 {
					if node.isEnd {
						isMatch = true
						return node
					}
				}
				break
			}
		}
		if !isMatch {
			for _, node := range children {
				if node.name == "**" {
					return node
				}
			}
		}
	}
	return nil
}
