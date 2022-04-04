package hw04lrucache

import (
	"errors"
)

var (
	ErrElemIsMissing  = errors.New("element is missing")
	listItemNilPoiner *ListItem
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	All() []interface{}
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	switch {
	case l.length == 0:
		return listItemNilPoiner
	case l.length == 1 && l.head != nil:
		return l.head
	case l.length == 1 && l.tail != nil:
		return l.tail
	default:
		return l.head
	}
}

func (l *list) Back() *ListItem {
	switch {
	case l.length == 0:
		return listItemNilPoiner
	case l.length == 1 && l.tail != nil:
		return l.tail
	case l.length == 1 && l.head != nil:
		return l.head
	default:
		return l.tail
	}
}

func (l *list) pushToFront(i *ListItem) *ListItem {
	// Если элементов не было
	// - то добавляемый элемент одновременно является и первым и последним
	if l.length == 0 {
		l.tail = i
		l.head = i
		l.length++
		return i
	}

	// Если был один элемент в списке
	// - Существующий элемент списка станет последним
	// - Добавляемый элемент будет первым и связан с существующим
	if l.length == 1 {
		// Сохраним первый элемент, так как он будет будет завершать список
		// И встанет после текущего
		prev := l.head
		// Обнулим указатели на головы перед операцией связывания
		l.head = listItemNilPoiner
		l.tail = listItemNilPoiner
		// Добавляемый элемент будет первым, он ссылается только на следующий элемент
		i.Next = prev
		i.Prev = listItemNilPoiner
		l.head = i
		// Элемент, который был в списке станет последним, он ссылается на предыдущий
		prev.Prev = i
		prev.Next = listItemNilPoiner
		l.tail = prev
		l.length++
		return i
	}

	prev := l.head
	l.head = i
	i.Next = prev
	prev.Prev = i
	l.length++

	return i
}

func (l *list) PushFront(v interface{}) *ListItem {
	return l.pushToFront(newListItem(v))
}

func (l *list) PushBack(v interface{}) *ListItem {
	return l.pushToBack(newListItem(v))
}

func (l *list) pushToBack(i *ListItem) *ListItem {
	// Если элементов не было
	// - то добавляемый элемент одновременно является и первым и последним
	if l.length == 0 {
		l.tail = i
		l.head = i
		l.length++
		return i
	}

	// Если был один элемент в списке
	// - Существующий элемент списка станет последним
	// - Добавляемый элемент будет первым и связан с существующим
	if l.length == 1 {
		// Сохраним первый элемент, так как он будет будет завершать список
		// И встанет после текущего
		prev := l.tail
		// Обнулим указатели на головы перед операцией связывания
		l.head = listItemNilPoiner
		l.tail = listItemNilPoiner
		// Добавляемый элемент будет первым, он ссылается только на следующий элемент
		i.Next = listItemNilPoiner
		i.Prev = prev
		l.tail = i
		// Элемент, который был в списке станет последним, он ссылается на предыдущий
		prev.Prev = listItemNilPoiner
		prev.Next = i
		l.head = prev
		l.length++
		return i
	}

	prev := l.tail
	i.Prev = prev
	prev.Next = i
	l.tail = i
	l.length++

	return i
}

func (l *list) Remove(i *ListItem) {
	switch {
	// Если это элемент посередине
	case i.Prev != nil && i.Next != nil:
		// Сохраним во временных переменных предыдущий и следующий элементы
		prevItem := i.Prev
		nextItem := i.Next
		// Скорректируем связи элементов очереди
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
		l.length--

	// Если в списке единственный элемент
	case l.length == 1:
		l.head = listItemNilPoiner
		l.tail = listItemNilPoiner
		l.length--

	// Если это первый элемент
	case i.Prev == nil && i.Next != nil:
		l.head = i.Next
		l.length--

	// Если это последний элемент
	case i.Prev != nil && i.Next == nil:
		l.tail = i.Prev
		l.length--
	}

	// Уберем указатели у удаляемого элемента
	i.Prev = listItemNilPoiner
	i.Next = listItemNilPoiner
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	// Если это элемент посередине, то очищаем ссылки и переносим на вершину списка
	case i.Prev != nil && i.Next != nil:
		// Сохраним во временных переменных предыдущий и следующий элементы
		prevItem := i.Prev
		nextItem := i.Next
		// Скорректируем связи элементов очереди
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
		// Уберем указатели у удаляемого элемента
		i.Prev = listItemNilPoiner
		i.Next = listItemNilPoiner
		l.length--

	// Если это первый элемент, то он уже на вершине списка
	case i.Prev == nil && i.Next != nil:
		return

	// Если это последний элемент, то очищаем ссылки и переносим на вершину списка
	case i.Prev != nil && i.Next == nil:
		l.tail = i.Prev
		l.tail.Next = listItemNilPoiner
		// Уберем указатели у удаляемого элемента
		i.Prev = listItemNilPoiner
		i.Next = listItemNilPoiner
		l.length--

	default:
		return
	}

	l.pushToFront(i)
}

func (l *list) All() []interface{} {
	var elems []interface{}
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}

func NewList() List {
	return &list{}
}

func newListItem(value interface{}) *ListItem {
	return &ListItem{
		Value: value,
	}
}
