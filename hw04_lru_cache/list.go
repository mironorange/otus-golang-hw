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
	capacity int
	lenght   int
	head     *ListItem
	tail     *ListItem
}

// Возвращает количество элементов в списке
func (l *list) Len() int {
	return l.lenght
}

// Возвращает элемент, находящийся сверху списка или ошибку если элемента нет
func (l *list) Front() *ListItem {
	switch {
	case l.lenght == 0:
		return listItemNilPoiner
	case l.lenght == 1 && l.head != nil:
		return l.head
	case l.lenght == 1 && l.tail != nil:
		return l.tail
	default:
		return l.head
	}
}

// Возвращает элемент, находящийся в конце списка или ошибку если элемента нет
func (l *list) Back() *ListItem {
	switch {
	case l.lenght == 0:
		return listItemNilPoiner
	case l.lenght == 1 && l.tail != nil:
		return l.tail
	case l.lenght == 1 && l.head != nil:
		return l.head
	default:
		return l.tail
	}
}

// Добавляет элемент на вершину списка
func (l *list) PushFront(v interface{}) *ListItem {
	if l.capacity <= 0 {
		return listItemNilPoiner
	}

	curItem := NewDoublyLinkedListItem(v)

	switch {
	// Если элементов не было, то добавляемый элемент одновременно является и первым и последним
	case l.lenght == 0:
		l.tail = curItem
		l.head = curItem
		l.lenght++

	// Если был один элемент в списке
	// - первый элемент заменит последний
	// - новый элемент будет связан с предыдущем
	case l.lenght == 1:
		prevHeadItem := l.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		l.head = curItem
		l.tail = prevHeadItem
		l.lenght++

	// В любой иной ситуации связываем новый элемент с текущей вершиной и заменяем им вершину
	default:
		prevHeadItem := l.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		l.head = curItem
		l.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if l.lenght > l.capacity {
		prevTailItem := l.tail
		l.tail = prevTailItem.Prev
		l.tail.Next = listItemNilPoiner
		prevTailItem.Prev = listItemNilPoiner
		l.lenght--
	}

	return curItem
}

// Добавляет элемент в конец списка
func (l *list) PushBack(v interface{}) *ListItem {
	if l.capacity <= 0 {
		return listItemNilPoiner
	}

	curItem := NewDoublyLinkedListItem(v)

	switch {
	// Если элементов нет, то добавляемый элемент одновременно является и первым и последним
	case l.lenght == 0:
		l.tail = curItem
		l.head = curItem
		l.lenght++

	// Если был один элемент в списке
	// - последний элемент заменит первый
	// - новый элемент будет связан с предыдущем
	case l.lenght == 1:
		prevTailItem := l.tail
		curItem.Prev = prevTailItem
		prevTailItem.Next = curItem
		l.head = prevTailItem
		l.tail = curItem
		l.lenght++

	// В любой иной ситуации связываем новый элемент с последним и заменяем им конец списка
	default:
		prevTailItem := l.tail
		curItem.Prev = prevTailItem
		prevTailItem.Next = curItem
		l.tail = curItem
		l.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if l.lenght > l.capacity {
		prevHeadItem := l.head
		l.head = prevHeadItem.Next
		l.head.Prev = listItemNilPoiner
		prevHeadItem.Next = listItemNilPoiner
		l.lenght--
	}

	return curItem
}

// Удалить элемент из списка
func (l *list) Remove(i *ListItem) {
	switch {
	// Если это элемент посередине
	case i.Prev != nil && i.Next != nil:
		// Сохраним во временных переменных предыдущий и следующий элементы
		prevItem := i.Prev
		nextItem := i.Next
		// Скорректируем связи элементов очереди
		nextItem.Prev = prevItem
		nextItem.Next = nextItem
		// Уберем указатели у удаляемого элемента
		i.Prev = listItemNilPoiner
		i.Next = listItemNilPoiner
		l.lenght--

	// Если это первый элемент
	case i.Prev == nil && i.Next != nil:
		l.head = i.Next
		i.Next = listItemNilPoiner
		l.lenght--

	// Если это последний элемент
	case i.Prev != nil && i.Next == nil:
		l.tail = i.Prev
		i.Prev = listItemNilPoiner
		l.lenght--
	}
}

// Удалить элемент из списка
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
		l.lenght--

	// Если это первый элемент, то он уже на вершине списка
	case i.Prev == nil && i.Next != nil:
		return

	// Если это последний элемент, то очищаем ссылки и переносим на вершину списка
	case i.Prev != nil && i.Next == nil:
		l.tail = i.Prev
		i.Prev = listItemNilPoiner
		l.lenght--

	default:
		return
	}

	curItem := i
	switch {
	// Если элементов не было, то добавляемый элемент одновременно является и первым и последним
	case l.lenght == 0:
		l.tail = curItem
		l.head = curItem
		l.lenght++

	// Если был один элемент в списке
	// - первый элемент заменит последний
	// - новый элемент будет связан с предыдущем
	case l.lenght == 1:
		prevHeadItem := l.head
		curItem.Prev = listItemNilPoiner
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		prevHeadItem.Next = listItemNilPoiner
		l.head = curItem
		l.tail = prevHeadItem
		l.lenght++

	// В любой иной ситуации связываем новый элемент с текущей вершиной и заменяем им вершину
	default:
		prevHeadItem := l.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		l.head = curItem
		l.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if l.lenght > l.capacity {
		prevTailItem := l.tail
		l.tail = prevTailItem.Prev
		l.tail.Next = listItemNilPoiner
		prevTailItem.Prev = listItemNilPoiner
		l.lenght--
	}
}

func (l *list) All() []interface{} {
	var items []interface{}
	item := l.Front()
	if item != nil {
		for i := item; i != nil; i = i.Next {
			items = append(items, i.Value)
		}
	}
	return items
}

func NewList(capacity int) List {
	return &list{
		capacity: capacity,
	}
}

// Конструктор связанного списка
func NewDoublyLinkedList(capacity int) List {
	return &list{
		capacity: capacity,
	}
}

// Конструктор элемента связанного списка
func NewDoublyLinkedListItem(value interface{}) *ListItem {
	return &ListItem{
		Value: value,
	}
}
