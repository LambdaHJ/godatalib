package skiplist


type SkipList struct {
	level []*ListItem
}

type ListItem struct {
	next *ListItem
}