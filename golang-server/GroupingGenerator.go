package main

import "golang-server/model"

func Partition(items []model.Goods) [][][]model.Goods {
	var result [][][]model.Goods
	backtracking(0, items, [][]model.Goods{}, &result)
	return result
}

func backtracking(index int, items []model.Goods, current [][]model.Goods, result *[][][]model.Goods) {
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

	newGroup := []model.Goods{currentItem}
	current = append(current, newGroup)
	backtracking(index+1, items, current, result)
	current = current[:len(current)-1]
}

func deepCopyGroups(groups [][]model.Goods) [][]model.Goods {
	copyGroups := make([][]model.Goods, len(groups))
	for idx := range groups {
		copyGroups[idx] = make([]model.Goods, len(groups[idx]))
		copy(copyGroups[idx], groups[idx])
	}
	return copyGroups
}
