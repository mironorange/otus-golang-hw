package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	expectedHeadValue   = "Head"
	expectedMiddleValue = "Middle"
	expectedTailValue   = "Tail"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

// Если в список добавлять элементы, то его размер увеличивается.
func TestDoublyLinkedLenght(t *testing.T) {
	l := NewList()
	expected := 0
	require.Equal(t, l.Len(), expected)

	l.PushFront(1)
	expected = 1
	require.Equal(t, l.Len(), expected)

	l.PushBack(3)
	expected = 2
	require.Equal(t, l.Len(), expected)
}

// Если в список добавлять элеенты сверху
// - то они будут связаны друг с другом по мере их добавления.
func TestPushFrontToDoublyLinkedList(t *testing.T) {
	l := NewList()

	l.PushFront(expectedTailValue)
	l.PushFront(expectedMiddleValue)
	l.PushFront(expectedHeadValue)

	expected := l.Front()
	require.Equal(t, expected.Value, expectedHeadValue)

	expectedMiddleElem := expected.Next
	require.Equal(t, expectedMiddleElem.Value, expectedMiddleValue)

	expectedTailElem := expectedMiddleElem.Next
	require.Equal(t, expectedTailElem.Value, expectedTailValue)
}

// Если в список добавлять элементы снизу
// - то они будут связаны друг с другом по мере их добавления.
func TestPushBackToDoublyLinkedList(t *testing.T) {
	l := NewList()

	l.PushBack(expectedHeadValue)
	l.PushBack(expectedMiddleValue)
	l.PushBack(expectedTailValue)

	expected := l.Back()
	require.Equal(t, expected.Value, expectedTailValue)

	expectedMiddleElem := expected.Prev
	require.Equal(t, expectedMiddleElem.Value, expectedMiddleValue)

	expectedHeadElem := expectedMiddleElem.Prev
	require.Equal(t, expectedHeadElem.Value, expectedHeadValue)
}

// Если в списке один элемент
// - то он будет возвращаться и с вершины и с конца.
func TestDoublyLinkedListWithOneElement(t *testing.T) {
	tests := []struct {
		build    func() List
		expected int
	}{
		{
			build: func() List {
				l := NewList()
				l.PushFront(1)
				return l
			},
			expected: 1,
		},
		{
			build: func() List {
				l := NewList()
				l.PushBack(2)
				return l
			},
			expected: 2,
		},
	}

	for _, tc := range tests {
		l := tc.build()
		result := l.Front()
		require.Equal(t, result.Value, tc.expected)

		result = l.Back()
		require.Equal(t, result.Value, tc.expected)
	}
}

// Если попытаться получить вершину пустого списка
// - то получаем Nil в ответ.
func TestEmptyDoublyLinkedListHead(t *testing.T) {
	l := NewList()
	i := l.Front()
	require.Nil(t, i, listItemNilPoiner)
}

// Если попытаться получить хвост пустого списка
// - то получаем Nil в ответ.
func TestEmptyDoublyLinkedTail(t *testing.T) {
	l := NewList()
	i := l.Back()
	require.Equal(t, i, listItemNilPoiner)
}

// Если удалять элементы из разных частей списка
// - то меняется количество элементов, вершина и хвост.
func TestRemoveDoublyLinkedListItem(t *testing.T) {
	tests := []struct {
		build          func() (List, *ListItem)
		expectedFront  int
		expectedBack   int
		expectedLength int
	}{
		{
			build: func() (List, *ListItem) {
				l := NewList()
				l.PushFront(3)                 // [3]
				l.PushFront(5)                 // [5, 3]
				firstDllItem := l.PushFront(1) // [1, 5, 3]
				return l, firstDllItem         // [5, 3]
			},
			expectedFront:  5,
			expectedBack:   3,
			expectedLength: 2,
		},
		{
			build: func() (List, *ListItem) {
				l := NewList()
				lastDllItem := l.PushFront(6)
				l.PushFront(2)
				l.PushFront(4)
				return l, lastDllItem
			},
			expectedFront:  4,
			expectedBack:   2,
			expectedLength: 2,
		},
		{
			build: func() (List, *ListItem) {
				l := NewList()
				l.PushFront(1)
				l.PushFront(3)
				middleDllItem := l.PushBack(2)
				l.PushBack(5)
				l.PushBack(7)
				return l, middleDllItem
			},
			expectedFront:  3,
			expectedBack:   7,
			expectedLength: 4,
		},
	}

	for _, tc := range tests {
		l, input := tc.build()

		l.Remove(input)

		result := l.Front()
		require.Equal(t, result.Value, tc.expectedFront)

		result = l.Back()
		require.Equal(t, result.Value, tc.expectedBack)

		require.Equal(t, l.Len(), tc.expectedLength)
	}
}

// Если перемещать элементы из разных частей списка
// - то меняется вершина и хвост списка.
func TestMoveToFrontDoublyLinkedListItem(t *testing.T) {
	tests := []struct {
		build         func() (List, *ListItem)
		expectedFront int
		expectedBack  int
	}{
		{
			build: func() (List, *ListItem) {
				l := NewList()
				l.PushFront(3)                 // [3]
				l.PushFront(5)                 // [5, 3]
				firstDllItem := l.PushFront(1) // [1, 5, 3]
				return l, firstDllItem         // [1, 5, 3]
			},
			expectedFront: 1,
			expectedBack:  3,
		},
		{
			build: func() (List, *ListItem) {
				l := NewList()
				lastDllItem := l.PushFront(6) // [6]
				l.PushFront(2)                // [2, 6]
				l.PushFront(4)                // [4, 2, 6]
				return l, lastDllItem         // [6, 4, 2]
			},
			expectedFront: 6,
			expectedBack:  2,
		},
		{
			build: func() (List, *ListItem) {
				l := NewList()
				l.PushFront(1)                 // [1]
				l.PushFront(3)                 // [3, 1]
				middleDllItem := l.PushBack(2) // [3, 1, 2]
				l.PushBack(5)                  // [3, 1, 2, 5]
				l.PushBack(7)                  // [3, 1, 2, 5, 7]
				return l, middleDllItem        // [2, 3, 1, 5, 7]
			},
			expectedFront: 2,
			expectedBack:  7,
		},
		{
			build: func() (List, *ListItem) {
				l := NewList()
				repeat := l.PushFront(1) // [1]
				l.PushFront(3)           // [3, 1]
				l.MoveToFront(repeat)    // [1, 3]
				l.PushFront(5)           // [5, 1, 3]
				l.MoveToFront(repeat)    // [1, 5, 3]
				l.PushFront(7)           // [7, 1, 5, 3]
				l.MoveToFront(repeat)    // [1, 7, 5, 3]
				l.PushFront(9)           // [9, 1, 7, 5, 3]
				l.MoveToFront(repeat)    // [1, 9, 7, 5, 3]
				l.PushFront(11)          // [11, 1, 9, 7, 5, 3]
				l.PushFront(13)          // [13, 11, 1, 9, 7, 5, 3]
				return l, repeat
			},
			expectedFront: 1,
			expectedBack:  3,
		},
	}

	for _, tc := range tests {
		l, input := tc.build()
		l.MoveToFront(input)

		result := l.Front()
		require.Equal(t, result.Value, tc.expectedFront)

		result = l.Back()
		require.Equal(t, result.Value, tc.expectedBack)
	}
}

// Если из очереди сделать слайс
// - то порядок элементов в нем будет соответсвовать порядку элементов в списке.
func TestGetAllDoublyLinkedListItems(t *testing.T) {
	l := NewList()
	expected := []interface{}{1, 3, 5, 7, 9}

	for _, value := range expected {
		l.PushBack(value)
	}
	require.Equal(t, l.All(), expected)
}

// Если из очереди удалить все элементы
// - то в очереде их не будет, вершина и хвост будут ссылаться на пустоту.
func TestRemoveAllDoublyLinkedListItems(t *testing.T) {
	l := NewList()
	input := []interface{}{1, 3, 5, 7, 9}

	items := make([]*ListItem, 0)
	for _, value := range input {
		items = append(items, l.PushBack(value))
	}
	for _, item := range items {
		l.Remove(item)
	}

	front := l.Front()
	require.Nil(t, front)

	back := l.Back()
	require.Nil(t, back)

	require.Equal(t, l.Len(), 0)
}
