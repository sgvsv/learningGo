package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type nodesList []tree
type activeLevels []bool
type tree struct {
	name  string
	nodes nodesList
}

func (n nodesList) Len() int           { return len(n) }
func (n nodesList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodesList) Less(i, j int) bool { return n[i].name < n[j].name }

// buildTree генерирует рекурсивно дерево по заданному пути к корню
func buildTree(fullPath string, printFiles bool) tree {
	t := tree{name: fullPath}
	nodes, _ := os.ReadDir(fullPath)
	for _, el := range nodes {
		if el.IsDir() {
			subTree := buildTree(fullPath+string(os.PathSeparator)+el.Name(), printFiles)
			subTree.name = el.Name()

			sort.Sort(subTree.nodes)

			t.nodes = append(t.nodes, subTree)
		} else if printFiles {
			info, _ := el.Info()
			size := ""
			if info.Size() == 0 {
				size = " (empty)"
			} else {
				size = " (" + strconv.FormatInt(info.Size(), 10) + "b)"
			}
			t.nodes = append(t.nodes, tree{name: el.Name() + size})
		}
	}
	return t
}

// visual отрисовывает дерево ASCI графикой
func visual(out io.Writer, nodes nodesList, level int, levels activeLevels) {
	for n, el := range nodes {
		lastNode := n == len(nodes)-1
		var subLevels = append(levels, !lastNode)

		for i := 0; i < level; i++ {
			if levels[i] {
				fmt.Fprint(out, "│	")
			} else {
				fmt.Fprint(out, "	")
			}
		}
		if lastNode {
			fmt.Fprintln(out, "└───"+el.name)
		} else {
			fmt.Fprintln(out, "├───"+el.name)
		}
		visual(out, el.nodes, level+1, subLevels)
	}
}
func dirTree(out io.Writer, path string, printFiles bool) error {
	t := buildTree(path, printFiles)
	visual(out, t.nodes, 0, nil)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
