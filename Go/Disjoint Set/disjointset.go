package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.

type node struct {
	parent *node
	val    int
	//rank   int
}

type DisjointStruct struct {
	allNodes map[int]*node //double check pointer impl
}

func (ds DisjointStruct) FindSet(u int) int {
	//check if in map
	_, ok := ds.allNodes[u]

	if ok == true { //exists
		currNode := ds.allNodes[u]

		for currNode.parent != currNode { //parent points to itself at root
			currNode = currNode.parent
		}

		return currNode.val //returns root
	}

	return u //set only made of int val
}

//helper func
func traverseToParent(val *node) *node {
	for val != val.parent {
		val = val.parent
	}

	return val
}

func (ds DisjointStruct) UnionSet(x, y int) int {
	_, okX := ds.allNodes[x]
	_, okY := ds.allNodes[y]

	if okX == false && okY == true {
		yNode := ds.allNodes[y]
		yNode = traverseToParent(yNode)

		xInit := node{val: x}
		var xNode *node
		xNode = &xInit
		xNode.parent = yNode
		ds.allNodes[x] = xNode
		return yNode.val
	} else if okY == false && okX == true {
		xNode := ds.allNodes[x]
		xNode = traverseToParent(xNode)

		yInit := node{val: y}
		var yNode *node
		yNode = &yInit
		yNode.parent = xNode
		ds.allNodes[y] = yNode
		return xNode.val
	} else if okX == true && okY == true {
		xNode := ds.allNodes[x]
		xNode = traverseToParent(xNode)

		yNode := ds.allNodes[y]
		yNode = traverseToParent(yNode)
		yNode.parent = xNode
		return xNode.val
	} else { //none in the map
		xInit := node{val: x}
		var xNode *node
		xNode = &xInit
		xNode.parent = xNode

		yInit := node{val: y}
		var yNode *node
		yNode = &yInit
		yNode.parent = xNode //x is root (arbritrary decision)

		ds.allNodes[x] = xNode
		ds.allNodes[y] = yNode

		return xNode.val
	}
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.

func NewDisjointSet() DisjointSet {
	//init the map and disjoint struct var
	var ds DisjointStruct
	ds.allNodes = make(map[int]*node)

	return ds
}
