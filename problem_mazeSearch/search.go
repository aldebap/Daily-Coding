package main

import (
	"errors"
	"fmt"
)

/*
	map1 = [
		[ '+', '0', '+', '0', '+', '0', '0' ],
		[ '+', '+', '+', '0', '+', '0', '0' ],
		[ '0', '0', '0', '0', '0', '0', '0' ],
		[ '+', '0', '+', '0', '0', '+', '0' ],
		[ '+', '0', '0', '0', '+', '0', '+' ],
		[ '+', '+', '+', '+', '0', '0', '0' ],
	]
	searchNearestExit(map1, start_3_1) ==> result end_1_4
*/

type position struct {
	row, column int
}

type node struct {
	position              *position
	left, right, up, down *node
}

func searchNearestExit(grid [][]byte, start *position) (*position, error) {
	var root *node = &node{
		position: start,
		left:     nil,
		right:    nil,
		up:       nil,
		down:     nil,
	}
	var searchNodes []*node

	searchNodes = append(searchNodes, root)
	for searchIndex := 0; searchIndex < len(searchNodes); searchIndex++ {

		searchNode := searchNodes[searchIndex]
		fmt.Printf("[debug] search node: (%d, %d): %c\n", searchNode.position.row, searchNode.position.column, grid[searchNode.position.row-1][searchNode.position.column-1])

		//	check if the current node is a solution
		if searchNode != root && (searchNode.position.row == 1 || searchNode.position.row == len(grid) ||
			searchNode.position.column == 1 || searchNode.position.column == len(grid[0])) {
			fmt.Printf("[debug] solution found: (%d, %d)\n", searchNode.position.row, searchNode.position.column)
			return searchNode.position, nil
		}

		//	check if it's possible to go up
		if searchNode.position.row > 1 && grid[searchNode.position.row-2][searchNode.position.column-1] != '+' {
			searchNodeUp := &node{
				position: &position{row: searchNode.position.row - 1, column: searchNode.position.column},
			}
			if !isTreeNode(root, searchNodeUp) {
				searchNode.up = searchNodeUp
				searchNodes = append(searchNodes, searchNode.up)
				fmt.Printf("[debug] up: search node added: (%d, %d)\n", searchNode.up.position.row, searchNode.up.position.column)
			}
		}

		//	check if it's possible to go down
		if searchNode.position.row < len(grid) && grid[searchNode.position.row][searchNode.position.column-1] != '+' {
			searchNodeDown := &node{
				position: &position{row: searchNode.position.row + 1, column: searchNode.position.column},
			}
			if !isTreeNode(root, searchNodeDown) {
				searchNode.down = searchNodeDown
				searchNodes = append(searchNodes, searchNode.down)
				fmt.Printf("[debug] down: search node added: (%d, %d)\n", searchNode.down.position.row, searchNode.down.position.column)
			}
		}

		//	check if it's possible to go left
		if searchNode.position.column > 1 && grid[searchNode.position.row-1][searchNode.position.column-2] != '+' {
			searchNodeLeft := &node{
				position: &position{row: searchNode.position.row, column: searchNode.position.column - 1},
			}
			if !isTreeNode(root, searchNodeLeft) {
				searchNode.left = searchNodeLeft
				searchNodes = append(searchNodes, searchNode.left)
				fmt.Printf("[debug] left: search node added: (%d, %d)\n", searchNode.left.position.row, searchNode.left.position.column)
			}
		}

		//	check if it's possible to go right
		if searchNode.position.column < len(grid[0]) && grid[searchNode.position.row-1][searchNode.position.column] != '+' {
			searchNodeRight := &node{
				position: &position{row: searchNode.position.row, column: searchNode.position.column + 1},
			}
			if !isTreeNode(root, searchNodeRight) {
				searchNode.right = searchNodeRight
				searchNodes = append(searchNodes, searchNode.right)
				fmt.Printf("[debug] right: search node added: (%d, %d)\n", searchNode.right.position.row, searchNode.right.position.column)
			}
		}
	}

	return nil, errors.New("can't find an exit for this grid")
}

func isTreeNode(root *node, searchNode *node) bool {

	if searchNode.position.row == root.position.row && searchNode.position.column == root.position.column {
		return true
	}

	if root.left != nil {
		if isTreeNode(root.left, searchNode) {
			return true
		}
	}
	if root.right != nil {
		if isTreeNode(root.right, searchNode) {
			return true
		}
	}
	if root.up != nil {
		if isTreeNode(root.up, searchNode) {
			return true
		}
	}
	if root.down != nil {
		if isTreeNode(root.down, searchNode) {
			return true
		}
	}

	return false
}
