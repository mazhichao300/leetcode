/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	List []int
	Idx  int
}

func build(nestedList []*NestedInteger, ans *[]int) {
	for i := 0; i < len(nestedList); i++ {
		o := nestedList[i]
		if o.IsInteger() {
			*ans = append(*ans, o.GetInteger())
		} else {
			build(o.GetList(), ans)
		}
	}
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	list := []int{}
	build(nestedList, &list)
	return &NestedIterator{
		list,
		-1,
	}
}

func (this *NestedIterator) Next() int {
	next := this.Idx + 1
	this.Idx = next
	return this.List[next]
}

func (this *NestedIterator) HasNext() bool {
	next := this.Idx + 1
	if len(this.List) > next {
		return true
	}
	return false
}