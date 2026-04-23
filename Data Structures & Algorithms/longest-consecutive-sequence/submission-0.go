type Node struct {
	backward *Node
	forward  *Node
	path     **int
}

func NewNode() *Node {
	var node Node
	path := 1
	ppath := &path
	node = Node{path: &ppath}
	return &node
}

func (n *Node) setBackward(backward *Node) {
	**backward.path += **n.path
	*n.path = *backward.path

	backward.forward = n
	n.backward = backward
}

func (n *Node) setForward(forward *Node) {
	**n.path += **forward.path
	*forward.path = *n.path

	forward.backward = n
	n.forward = forward
} 

func (n *Node) Set(backward, forward *Node) {
	fmt.Println("using", backward, ",", forward, "for", n)
	if backward != nil {
		if backward.forward == nil {
			n.setBackward(backward)
		}
	}

	if forward != nil {
		if forward.backward == nil {
			n.setForward(forward)
		}
	}
}

func longestConsecutive(nums []int) int {
	m := make(map[int]*Node)
	maxV := 0
	
	for _, num := range nums { 
		this, ok := m[num]
		if !ok {
			this = NewNode()
			m[num] = this
		}

		backward, _ := m[num-1]
		forward, _ := m[num+1]

		this.Set(backward, forward)
	 }

	for _, node := range m {
		n := **node.path
		if n > maxV {
			maxV = n
		}
	}

	return maxV
}
