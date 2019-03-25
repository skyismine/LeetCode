package commonds

/**
字典树结构和实现

详细文档参见 doc/commonds/Tire.md
*/
type TireNode struct {
	// 该节点下面的子节点
	Links []*TireNode
	// 该节点下子节点的真实个数
	Size int
	// 该节点是否作为结束节点，结束节点标识它代表一个单词的结束
	IsEnd bool
	// 该节点作为结束节点的次数，用于标识一个单词出现的次数，可用于词频统计
	EndCount int
	// 该节点被经过的次数，可用于前缀数目统计和前缀匹配
	PassCount int
}

/**
 打印字典树
 */
func (root *TireNode) Print() {

}

/**
 插入一个字符串到字典树
 */
func (root *TireNode) Insert(word string) {
	if word == "" {
		return
	}
	strbytes := []byte(word)
	node := root
	for index, value := range strbytes {
		newnode := node.Links[value-'a']
		if newnode == nil {
			newnode = &TireNode{make([]*TireNode, 26), 0, false, 0, 1}
			node.Links[value-'a'] = newnode
			node.Size++
		} else {
			newnode.PassCount++
		}
		if index == len(strbytes)-1 {
			newnode.IsEnd = true
			newnode.EndCount++
		}
		node = newnode
	}
}

/**
 查找字典树和给定字符串的最长公共前缀(Longest Common Prefix)
 */
func (root *TireNode) SearchLCP(word string) (prefix string) {
	if word == "" || root.Size != 1 {
		return ""
	}
	var prefixbytes []byte
	wordbytes := []byte(word)
	node := root;
	for _,value := range wordbytes {
		// nextnode 用于做字符匹配，判断当前字符是否存在，如果是 nil 则说明当前字符不存在
		// 如果以当前字符-'a'为索引的 TrieNode 存在则说明当前字符存在
		nextnode := node.Links[value-'a']
		// node 有结束标记说明此处有字符串已经查找完了，所以已经找到最大公共前缀了
		// node 子节点数目不为1说明该节点下有多个节点，字符串再次处分裂了，所以该节点就是最大公共前缀节点
		if node.IsEnd || node.Size != 1 || nextnode == nil {
			return string(prefixbytes)
		}
		prefixbytes = append(prefixbytes, value)
		node = nextnode
	}

	return string(prefixbytes)
}

func NewTireRoot() (*TireNode) {
	return &TireNode{make([]*TireNode, 26), 0, false, 0, 0}
}
