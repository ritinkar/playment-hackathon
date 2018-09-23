package session

import (
	"strings"
)

type directoryNode struct {
	children []*directoryNode
	parent   *directoryNode
	value    string
}

// Session : Structure to represent a session
type Session struct {
	currentNode *directoryNode
	root        *directoryNode
}

// StartSession : Function to start a session and initialize root
func StartSession() *Session {
	root := &directoryNode{
		value:    "/",
		children: make([]*directoryNode, 0),
		parent:   nil}
	return &Session{root: root, currentNode: root}
}

// AddNodes : Add dirs as children to cwd
func (currentSession *Session) AddNodes(dirNames []string) (bool, []string, []string) {
	errorDirs := make([]string, 0)
	atLeastOneDirCreated := false
	invalidDirNames := make([]string, 0)
	for _, dirName := range dirNames {
		dirExists := false

		//Dirnames shouldn't contain a "/"
		if strings.ContainsAny(dirName, "/") {
			invalidDirNames = append(invalidDirNames, dirName)
		}
		for _, node := range currentSession.currentNode.children {
			if (node.value) == dirName {
				dirExists = true
				errorDirs = append(errorDirs, dirName)
			}

		}
		if !dirExists && !strings.ContainsAny(dirName, "/") {
			currentSession.currentNode.children = append(currentSession.currentNode.children,
				&directoryNode{value: dirName,
					parent:   currentSession.currentNode,
					children: make([]*directoryNode, 0)})
			atLeastOneDirCreated = true

		}
	}
	return atLeastOneDirCreated, errorDirs, invalidDirNames
}

// GetCwd : Get the value of current working Directory
func (currentSession *Session) GetCwd() string {
	pathValues := make([]string, 0)
	for node := currentSession.currentNode; node != currentSession.root; node = node.parent {
		pathValues = append(pathValues, node.value)
	}
	path := "/"
	for i := len(pathValues) - 1; i >= 0; i-- {
		path += (pathValues[i]) + "/"
	}
	return path

}

// GetChildrenDirectories : Get the values of children of current working Directory
func (currentSession *Session) GetChildrenDirectories() []string {
	childrenDirectories := make([]string, 0)
	for _, node := range currentSession.currentNode.children {
		childrenDirectories = append(childrenDirectories, node.value)
	}
	return childrenDirectories
}

// SwitchCurrentDir : Switch current working Directory
func (currentSession *Session) SwitchCurrentDir(path string) bool {
	splitPath := strings.Split(path, "/")

	nextNode := currentSession.currentNode
	startIndex := 0
	endIndex := len(splitPath)

	invalidPath := false

	// if the first arg is empty string then it's an absolute path
	if splitPath[0] == "" {
		nextNode = currentSession.root
		startIndex = 1
	}

	//If there's a trailing slash then the last arg is an empty string
	if splitPath[endIndex-1] == "" {
		endIndex--
	}

	for _, dir := range splitPath[startIndex:endIndex] {
		foundFlag := 0
		for _, node := range nextNode.children {
			if (node.value) == dir {
				nextNode = node
				foundFlag = 1
			}

		}
		if foundFlag == 0 {
			invalidPath = true
			return invalidPath
		}

	}
	currentSession.currentNode = nextNode
	return invalidPath
}

// Clear : Resets the filetree
func (currentSession *Session) Clear() {
	root := &directoryNode{
		value:    "/",
		children: make([]*directoryNode, 0),
		parent:   nil}
	currentSession.root = root
	currentSession.currentNode = root
}

// RemoveNodes : Remove specified Directories
func (currentSession *Session) RemoveNodes(paths []string) (bool, []string) {
	errorPaths := make([]string, 0)
	atLeastOneDirDeleted := false
	for _, path := range paths[0:] {
		splitPath := strings.Split(path, "/")

		nextNode := currentSession.currentNode
		startIndex := 0
		endIndex := len(splitPath)

		invalidPath := false

		// if the first arg is empty string then it's an absolute path
		if splitPath[0] == "" {
			nextNode = currentSession.root
			startIndex = 1
		}

		//If there's a trailing slash then the last arg is an empty string
		if splitPath[endIndex-1] == "" {
			endIndex--
		}

		for _, dir := range splitPath[startIndex:endIndex] {
			foundFlag := 0
			for _, node := range nextNode.children {
				if (node.value) == dir {
					nextNode = node
					foundFlag = 1
				}

			}
			if foundFlag == 0 {
				invalidPath = true
				errorPaths = append(errorPaths, path)
			}

		}

		// Delete the node that matched the path
		if !invalidPath {
			atLeastOneDirDeleted = true
			nodes := nextNode.parent.children
			for index, node := range nodes {
				if node == nextNode {
					nodes[index], nodes[len(nodes)-1] = nodes[len(nodes)-1], nodes[index]
				}
			}

			// set the current node to the parent node of the node to be deleted just in case
			// absolute path is provided and its a parent dir of current dir
			currentSession.currentNode = nextNode.parent
			nextNode.parent.children = nodes[:len(nodes)-1]

		}
	}
	return atLeastOneDirDeleted, errorPaths

}
