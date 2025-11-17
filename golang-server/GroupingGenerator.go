package main

func Partition(items []Goods) [][][]Goods {
	var result [][][]Goods
	backtracking(0, items, [][]Goods{}, &result)
	return result
}

func backtracking(index int, items []Goods, current [][]Goods, result *[][][]Goods) {
	if index == len(items) {
		copied := deepCopyGroups(current)
		*result = append(*result, copied)
		return
	}

	currentItem := items[index]

	for idx := range current {
		current[idx] = append(current[idx], currentItem)
		backtracking(index+1, items, current, result)
		current[idx] = current[idx][:len(current[idx])-1]
	}

	newGroup := []Goods{currentItem}
	current = append(current, newGroup)
	backtracking(index+1, items, current, result)
	current = current[:len(current)-1]
}

func deepCopyGroups(groups [][]Goods) [][]Goods {
	copyGroups := make([][]Goods, len(groups))
	for idx := range groups {
		copyGroups[idx] = make([]Goods, len(groups[idx]))
		copy(copyGroups[idx], groups[idx])
	}
	return copyGroups
}
