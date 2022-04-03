package hw04lrucache

import (
	"errors"
)

var (
	ErrElemIsMissing  = errors.New("element is missing")
	listItemNilPoiner *DLListItem
)

type DLList struct {
	capacity int
	lenght   int
	head     *DLListItem
	tail     *DLListItem
}

type DLListItem struct {
	Value interface{}
	Next  *DLListItem
	Prev  *DLListItem
}

type DoublyLinkedLister interface {
	Len() int
	Front() (*DLListItem, error)
	Back() (*DLListItem, error)
	PushFront(v interface{}) *DLListItem
	PushBack(v interface{}) *DLListItem
	Remove(i *DLListItem)
	MoveToFront(i *DLListItem)
	All() []interface{}
}

// Возвращает количество элементов в списке
func (dll *DLList) Len() int {
	return dll.lenght
}

// Возвращает элемент, находящийся сверху списка или ошибку если элемента нет
func (dll *DLList) Front() (*DLListItem, error) {
	switch {
	case dll.lenght == 0:
		return listItemNilPoiner, ErrElemIsMissing
	case dll.lenght == 1 && dll.head != nil:
		return dll.head, nil
	case dll.lenght == 1 && dll.tail != nil:
		return dll.tail, nil
	default:
		return dll.head, nil
	}
}

// Возвращает элемент, находящийся в конце списка или ошибку если элемента нет
func (dll *DLList) Back() (*DLListItem, error) {
	switch {
	case dll.lenght == 0:
		return listItemNilPoiner, ErrElemIsMissing
	case dll.lenght == 1 && dll.tail != nil:
		return dll.tail, nil
	case dll.lenght == 1 && dll.head != nil:
		return dll.head, nil
	default:
		return dll.tail, nil
	}
}

// Добавляет элемент на вершину списка
func (dll *DLList) PushFront(v interface{}) *DLListItem {
	if dll.capacity <= 0 {
		return listItemNilPoiner
	}

	curItem := NewDoublyLinkedListItem(v)

	switch {
	// Если элементов не было, то добавляемый элемент одновременно является и первым и последним
	case dll.lenght == 0:
		dll.tail = curItem
		dll.head = curItem
		dll.lenght++

	// Если был один элемент в списке
	// - первый элемент заменит последний
	// - новый элемент будет связан с предыдущем
	case dll.lenght == 1:
		prevHeadItem := dll.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		dll.head = curItem
		dll.tail = prevHeadItem
		dll.lenght++

	// В любой иной ситуации связываем новый элемент с текущей вершиной и заменяем им вершину
	default:
		prevHeadItem := dll.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		dll.head = curItem
		dll.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if dll.lenght > dll.capacity {
		prevTailItem := dll.tail
		dll.tail = prevTailItem.Prev
		dll.tail.Next = listItemNilPoiner
		prevTailItem.Prev = listItemNilPoiner
		dll.lenght--
	}

	return curItem
}

// Добавляет элемент в конец списка
func (dll *DLList) PushBack(v interface{}) *DLListItem {
	if dll.capacity <= 0 {
		return listItemNilPoiner
	}

	curItem := NewDoublyLinkedListItem(v)

	switch {
	// Если элементов нет, то добавляемый элемент одновременно является и первым и последним
	case dll.lenght == 0:
		dll.tail = curItem
		dll.head = curItem
		dll.lenght++

	// Если был один элемент в списке
	// - последний элемент заменит первый
	// - новый элемент будет связан с предыдущем
	case dll.lenght == 1:
		prevTailItem := dll.tail
		curItem.Prev = prevTailItem
		prevTailItem.Next = curItem
		dll.head = prevTailItem
		dll.tail = curItem
		dll.lenght++

	// В любой иной ситуации связываем новый элемент с последним и заменяем им конец списка
	default:
		prevTailItem := dll.tail
		curItem.Prev = prevTailItem
		prevTailItem.Next = curItem
		dll.tail = curItem
		dll.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if dll.lenght > dll.capacity {
		prevHeadItem := dll.head
		dll.head = prevHeadItem.Next
		dll.head.Prev = listItemNilPoiner
		prevHeadItem.Next = listItemNilPoiner
		dll.lenght--
	}

	return curItem
}

// Удалить элемент из списка
func (dll *DLList) Remove(i *DLListItem) {
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
		dll.lenght -= 1

	// Если это первый элемент
	case i.Prev == nil && i.Next != nil:
		dll.head = i.Next
		i.Next = listItemNilPoiner
		dll.lenght -= 1

	// Если это последний элемент
	case i.Prev != nil && i.Next == nil:
		dll.tail = i.Prev
		i.Prev = listItemNilPoiner
		dll.lenght -= 1
	}
}

// Удалить элемент из списка
func (dll *DLList) MoveToFront(i *DLListItem) {
	switch {
	// Если это элемент посередине, то очищаем ссылки и переносим на вершину списка
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
		dll.lenght -= 1

	// Если это первый элемент, то он уже на вершине списка
	case i.Prev == nil && i.Next != nil:
		return

	// Если это последний элемент, то очищаем ссылки и переносим на вершину списка
	case i.Prev != nil && i.Next == nil:
		dll.tail = i.Prev
		i.Prev = listItemNilPoiner
		dll.lenght -= 1

	default:
		return
	}

	curItem := i
	switch {
	// Если элементов не было, то добавляемый элемент одновременно является и первым и последним
	case dll.lenght == 0:
		dll.tail = curItem
		dll.head = curItem
		dll.lenght++

	// Если был один элемент в списке
	// - первый элемент заменит последний
	// - новый элемент будет связан с предыдущем
	case dll.lenght == 1:
		prevHeadItem := dll.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		dll.head = curItem
		dll.tail = prevHeadItem
		dll.lenght++

	// В любой иной ситуации связываем новый элемент с текущей вершиной и заменяем им вершину
	default:
		prevHeadItem := dll.head
		curItem.Next = prevHeadItem
		prevHeadItem.Prev = curItem
		dll.head = curItem
		dll.lenght++
	}

	// Если после добавления элемента, размер превышает его емкость, срезать конец
	if dll.lenght > dll.capacity {
		prevTailItem := dll.tail
		dll.tail = prevTailItem.Prev
		dll.tail.Next = listItemNilPoiner
		prevTailItem.Prev = listItemNilPoiner
		dll.lenght--
	}
}

func (dll *DLList) All() []interface{} {
	item, err := dll.Front()
	if err != nil {
		return make([]interface{}, 0)
	}
	items := make([]interface{}, 0, dll.lenght)
	for i := item; i != nil; i = i.Next {
		items = append(items, i.Value)
	}
	return items
}

// Конструктор связанного списка
func NewDoublyLinkedList(capacity int) DoublyLinkedLister {
	return &DLList{
		capacity: capacity,
		lenght:   0,
	}
}

// Конструктор элемента связанного списка
func NewDoublyLinkedListItem(value interface{}) *DLListItem {
	return &DLListItem{
		Value: value,
	}
}
