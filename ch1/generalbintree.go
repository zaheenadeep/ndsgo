package generalbintree

type GNode struct {
	Data		int
	LSon		*GNode
	LThread		bool
	RSon		*GNode
	RThread		bool
	Father		*GNode
}