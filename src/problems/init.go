package problems

var METHOD string

const (
	METHOD_MINE = "MINE"
	METHOD_MINE2 = "MINE2"
	METHOD_LCA = "LCA"
)

func init() {
	SetMethod(METHOD_MINE)
}

func SetMethod(method string) {
	METHOD = method
}
