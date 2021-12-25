package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length    int
	firstNode *ListItem
	lastNode  *ListItem
}

func NewList() List {
	l := new(list)
	return l
}

func (l list) Front() *ListItem {
	return l.firstNode
}

func (l list) Back() *ListItem {
	return l.lastNode
}

func (l list) Len() int {
	return l.length
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.length == 0 {
		l.firstNode = li
		l.lastNode = li
	} else {
		l.firstNode.Prev = li
		li.Next = l.Front()
		l.firstNode = li
	}
	l.length++
	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.length == 0 {
		l.firstNode = li
		l.lastNode = li
	} else {
		l.lastNode.Next = li
		li.Prev = l.Back()
		l.lastNode = li
	}
	l.length++
	return li
}

func (l *list) Remove(i *ListItem) {
	// Если список пуст, сразу безопасно выходим
	if l.length == 0 {
		return
	}

	// Если список из одного элемента, просто убираем ссылки на него
	if l.Len() == 1 {
		l.firstNode = nil
		l.lastNode = nil
	} else {
		// Иначе смотрим на наличие соседей
		switch {
		case i == l.firstNode:
			l.firstNode = l.firstNode.Next
			l.firstNode.Prev = nil
		case i == l.lastNode:
			l.lastNode = l.lastNode.Prev
			l.lastNode.Next = nil
		default:
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
		}
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

/*
func main() {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]
	fmt.Println("3 = ", l.Len())
	//require.Equal(t, 3, l.Len())

	middle := l.Front().Next // 20
	l.Remove(middle)         // [10, 30]
	fmt.Println("2 = ", l.Len())
	fmt.Println("10 = ", l.Front().Value.(int))
	fmt.Println("30 = ", l.Back().Value.(int))

	// Удаление единственного элемента, удаление из пустого списка

	for i, v := range [...]int{40, 50, 60, 70, 80} {
		if i%2 == 0 {
			l.PushFront(v)
		} else {
			l.PushBack(v)
		}
	} // [80, 60, 40, 10, 30, 50, 70]

	fmt.Println("7 = ", l.Len())
	fmt.Println("80 = ", l.Front().Value)
	fmt.Println("70 = ", l.Back().Value)

	l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
	l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	fmt.Println("{70, 80, 60, 40, 10, 30, 50} = ", elems)
}
*/
