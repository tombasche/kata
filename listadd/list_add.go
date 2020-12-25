package listadd

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Values() []int {
	curr := n
	var result []int
	for curr != nil {
		result = append(result, curr.Value)
		curr = curr.Next
	}
	return result
}

func ListAdd(a *Node, b *Node) *Node {

	remainder := 0
	p := a
	q := b
	result := &Node{Value: 0, Next: nil}
	curr := result
	for p != nil || q != nil || remainder != 0 {
		x := 0
		y := 0
		if p != nil {
			x = p.Value
			p = p.Next
		}
		if q != nil {
			y = q.Value
			q = q.Next
		}
		sum := x + y + remainder
		remainder = sum / 10
		n := &Node{Value: sum % 10, Next: nil}
		curr.Next = n
		curr = curr.Next
	}
	return result.Next
}
