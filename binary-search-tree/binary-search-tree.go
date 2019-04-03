// API:
//
// type SearchTreeData struct {
//	left  *SearchTreeData
//	data  int
//	right *SearchTreeData
// }
//
// func Bst(int) *SearchTreeData
// func (*SearchTreeData) Insert(int)
// func (*SearchTreeData) MapString(func(int) string) []string
// func (*SearchTreeData) MapInt(func(int) int) []int

package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(value int) SearchTreeData {
	return SearchTreeData{data: 4}
}

func (sTree *SearchTreeData) Insert(value int) {
	if value <= sTree.data {
		if sTree.left != nil {
			sTree.left.Insert(value)
		} else {
			sTree.left = &SearchTreeData{data: value}
		}
	} else {
		if sTree.right != nil {
			sTree.right.Insert(value)
		} else {
			sTree.right = &SearchTreeData{data: value}
		}
	}
}

type stringCallback func(int) string

func (std *SearchTreeData) MapString(f stringCallback) (result []string) {
	if std.left != nil {
		result = append(std.left.MapString(f), result...)
	}
	result = append(result, []string{f(std.data)}...)
	if std.right != nil {
		result = append(result, std.right.MapString(f)...)
	}
	return result
}

type intCallback func(int) int

func (std *SearchTreeData) MapInt(f intCallback) (result []int) {
	if std.left != nil {
		result = append(std.left.MapInt(f), result...)
	}
	result = append(result, []int{f(std.data)}...)
	if std.right != nil {
		result = append(result, std.right.MapInt(f)...)
	}
	return result
}
