package main

import (
	"fmt"
)


type Relation struct {
	parent string
	child string
}

func relate() {
	var relations []Relation

	relations = append(relations, Relation{ "animal", "mammal" })
	relations = append(relations, Relation{ "animal", "bird" })
	relations = append(relations, Relation{ "lifeform", "animal" })
	relations = append(relations, Relation{ "cat", "lion" })
	relations = append(relations, Relation{ "mammal", "cat" })
	relations = append(relations, Relation{ "animal", "fish" })

	printRelationsTree(relations)
}

func printRelationsTree(relations []Relation) {
	var rootName string
	for _, relation := range relations {
		isRootName := true
		for _, children := range relations {
			if children.child == relation.parent {
				isRootName = false
				continue
			}
		}
		if isRootName {
			rootName = relation.parent
			break
		}
	}

	adjMap := map[string][]string{}
	for _, relation := range relations {
		adjMap[relation.parent] = append(adjMap[relation.parent], relation.child)
	}

	dfs(rootName, 0, adjMap)
}

func dfs(rootName string, level int, adjMap map[string][]string) {
	for i := 0; i < level; i++ {
		fmt.Printf("\t")
	}

	fmt.Printf("%s\n", rootName)

	if _, ok := adjMap[rootName]; !ok {
		return
	}

	children := adjMap[rootName]

	for _, childName := range children {
		dfs(childName, level + 1, adjMap)
	}
}

type TreeNode struct {
	name string
	children []TreeNode
}