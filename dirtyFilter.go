package dirtyFilter

type node struct {
	end   bool
	child map[string]*node
}

type trieTree struct {
	root *node
}

func (t *trieTree) AddDirtyWord(words []string) {
	if len(words) < 1 {
		return
	}
	for _, word := range words {
		runeWord := []rune(word)
		nowNode := t.root
		for i := 0; i < len(runeWord); i++ {
			if _, ok := nowNode.child[string(runeWord[i])]; !ok {
				nowNode.child[string(runeWord[i])] = &node{child: make(map[string]*node)}
			}
			if i == len(runeWord)-1 {
				nowNode.end = true
			}
			nowNode = nowNode.child[string(runeWord[i])]
		}
	}
}

func (t *trieTree) SearchDirtyWord(content string) string {
	if len(content) < 1 {
		return content
	}
	runeContent := []rune(content)
	start := -1
	tag := -1
	nowNode := t.root
	for i := 0; i < len(runeContent); i++ {
		// 无效符号过滤
		if _, ok := InvalidWord[string(runeContent[i])]; ok {
			continue
		}
		if _, ok := nowNode.child[string(runeContent[i])]; ok {
			tag++
			if tag == 0 {
				start = i
			}
			if !nowNode.end {
				nowNode = nowNode.child[string(runeContent[i])]
			} else {
				// 全匹配 将敏感词的第一个文字到最后一个文字全部替换为“*”
				for y := start; y < i+1; y++ {
					runeContent[y] = 42
				}
				//重置标志数据
				nowNode = t.root
				start = -1
				tag = -1
			}
		} else {
			// 如果非全匹配， 终止 从下一个位置重新开始
			if start != -1 {
				i = start + 1
			}
			//重置标志参数
			nowNode = t.root
			start = -1
			tag = -1
		}
	}
	return string(runeContent)
}
