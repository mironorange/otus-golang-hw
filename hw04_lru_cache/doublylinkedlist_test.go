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

// Если в список добавлять элементы, то его размер ростет
func TestDoublyLinkedLenght(t *testing.T) {
	input := NewDoublyLinkedList(3)
	expected := 0
	require.Equal(t, input.Len(), expected)

	input.PushFront(nil)
	expected = 1
	require.Equal(t, input.Len(), expected)

	input.PushBack(nil)
	expected = 2
	require.Equal(t, input.Len(), expected)
}

// Если в список добавлять элеенты, то они будут связаны по мере добавления
func TestPushFrontToDoublyLinkedList(t *testing.T) {
	input := NewDoublyLinkedList(3)

	input.PushFront(expectedTailValue)
	input.PushFront(expectedMiddleValue)
	input.PushFront(expectedHeadValue)

	expected, err := input.Front()
	require.Equal(t, expected.Value, expectedHeadValue)
	require.NoError(t, err)

	expectedMiddleElem := expected.Next
	require.Equal(t, expectedMiddleElem.Value, expectedMiddleValue)

	expectedTailElem := expectedMiddleElem.Next
	require.Equal(t, expectedTailElem.Value, expectedTailValue)
}

// Если в список добавлять элеенты, то они будут связаны по мере добавления
func TestPushBackToDoublyLinkedList(t *testing.T) {
	input := NewDoublyLinkedList(3)

	input.PushBack(expectedHeadValue)
	input.PushBack(expectedMiddleValue)
	input.PushBack(expectedTailValue)

	expected, err := input.Back()
	require.Equal(t, expected.Value, expectedTailValue)
	require.NoError(t, err)

	expectedMiddleElem := expected.Prev
	require.Equal(t, expectedMiddleElem.Value, expectedMiddleValue)

	expectedHeadElem := expectedMiddleElem.Prev
	require.Equal(t, expectedHeadElem.Value, expectedHeadValue)
}

// Если у списка пустая емкость, то в него нельзя добавить элементы
func TestDoublyLinkedListZeroCapacity(t *testing.T) {
	tests := []struct {
		build          func() DoublyLinkedLister
		expectedLength int
	}{
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(0)
				dll.PushFront(1)
				return dll
			},
			expectedLength: 0,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(0)
				dll.PushBack(5)
				return dll
			},
			expectedLength: 0,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(0)
				dll.PushFront(1)
				dll.PushBack(5)
				return dll
			},
			expectedLength: 0,
		},
	}

	for _, tc := range tests {
		dll := tc.build()
		require.Equal(t, dll.Len(), tc.expectedLength)

		result, err := dll.Front()
		require.Equal(t, result, listItemNilPoiner)
		require.Error(t, err)

		result, err = dll.Back()
		require.Equal(t, result, listItemNilPoiner)
		require.Error(t, err)
	}
}

// Если в списке один элемент, то он будет возвращаться и с вершины и с конца
func TestDoublyLinkedListOverflow(t *testing.T) {
	tests := []struct {
		build          func() DoublyLinkedLister
		expectedFront  int
		expectedBack   int
		expectedLength int
	}{
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(4)
				dll.PushFront(1)
				dll.PushFront(3)
				dll.PushFront(5)
				dll.PushFront(7)
				dll.PushFront(9)
				dll.PushFront(11)
				return dll
			},
			expectedFront:  11,
			expectedBack:   5,
			expectedLength: 4,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(4)
				dll.PushBack(13)
				dll.PushBack(15)
				dll.PushBack(17)
				dll.PushBack(19)
				dll.PushBack(21)
				dll.PushBack(23)
				return dll
			},
			expectedFront:  17,
			expectedBack:   23,
			expectedLength: 4,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(1)
				dll.PushFront(1)
				dll.PushFront(3)
				dll.PushFront(5)
				return dll
			},
			expectedFront:  5,
			expectedBack:   5,
			expectedLength: 1,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(1)
				dll.PushBack(13)
				dll.PushBack(15)
				dll.PushBack(17)
				return dll
			},
			expectedFront:  17,
			expectedBack:   17,
			expectedLength: 1,
		},
	}

	for _, tc := range tests {
		dll := tc.build()
		require.Equal(t, dll.Len(), tc.expectedLength)

		result, err := dll.Front()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedFront)

		result, err = dll.Back()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedBack)
	}
}

// Если в списке один элемент, то он будет возвращаться и с вершины и с конца
func TestDoublyLinkedListWithOneElement(t *testing.T) {
	tests := []struct {
		build    func() DoublyLinkedLister
		expected int
	}{
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(3)
				dll.PushFront(1)
				return dll
			},
			expected: 1,
		},
		{
			build: func() DoublyLinkedLister {
				dll := NewDoublyLinkedList(3)
				dll.PushBack(2)
				return dll
			},
			expected: 2,
		},
	}

	for _, tc := range tests {
		dll := tc.build()
		result, err := dll.Front()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expected)

		result, err = dll.Back()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expected)
	}
}

// Если попытаться получить вершину пустого списка, то получаем ошибку
func TestEmptyDoublyLinkedListHead(t *testing.T) {
	input := NewDoublyLinkedList(1)
	_, err := input.Front()
	require.Error(t, err)
}

// Если попытаться получить хвост пустого списка, то получаем ошибку
func TestEmptyDoublyLinkedTail(t *testing.T) {
	input := NewDoublyLinkedList(1)
	_, err := input.Back()
	require.Error(t, err)
}

// Если удалять элементы из разных частей списка, то меняется количество элементов, вершина и хвост
func TestRemoveDoublyLinkedListItem(t *testing.T) {
	tests := []struct {
		build          func() (DoublyLinkedLister, *DLListItem)
		expectedFront  int
		expectedBack   int
		expectedLength int
	}{
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(3)
				dll.PushFront(3)
				dll.PushFront(5)
				firstDllItem := dll.PushFront(1)
				return dll, firstDllItem
			},
			expectedFront:  5,
			expectedBack:   3,
			expectedLength: 2,
		},
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(3)
				lastDllItem := dll.PushFront(6)
				dll.PushFront(2)
				dll.PushFront(4)
				return dll, lastDllItem
			},
			expectedFront:  4,
			expectedBack:   2,
			expectedLength: 2,
		},
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(7)
				dll.PushFront(1)
				dll.PushFront(3)
				middleDllItem := dll.PushBack(2)
				dll.PushBack(5)
				dll.PushBack(7)
				return dll, middleDllItem
			},
			expectedFront:  3,
			expectedBack:   7,
			expectedLength: 4,
		},
	}

	for _, tc := range tests {
		dll, input := tc.build()

		dll.Remove(input)

		result, err := dll.Front()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedFront)

		result, err = dll.Back()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedBack)

		require.Equal(t, dll.Len(), tc.expectedLength)
	}
}

// Если перемещать элементы из разных частей списка, то меняется вершина и хвост списка
func TestMoveToFrontDoublyLinkedListItem(t *testing.T) {
	tests := []struct {
		build         func() (DoublyLinkedLister, *DLListItem)
		expectedFront int
		expectedBack  int
	}{
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(3)
				dll.PushFront(3)
				dll.PushFront(5)
				firstDllItem := dll.PushFront(1)
				return dll, firstDllItem
			},
			expectedFront: 1,
			expectedBack:  3,
		},
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(3)
				lastDllItem := dll.PushFront(6)
				dll.PushFront(2)
				dll.PushFront(4)
				return dll, lastDllItem
			},
			expectedFront: 6,
			expectedBack:  2,
		},
		{
			build: func() (DoublyLinkedLister, *DLListItem) {
				dll := NewDoublyLinkedList(7)
				dll.PushFront(1)
				dll.PushFront(3)
				middleDllItem := dll.PushBack(2)
				dll.PushBack(5)
				dll.PushBack(7)
				return dll, middleDllItem
			},
			expectedFront: 2,
			expectedBack:  7,
		},
	}

	for _, tc := range tests {
		dll, input := tc.build()

		dll.MoveToFront(input)

		result, err := dll.Front()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedFront)

		result, err = dll.Back()
		require.NoError(t, err)
		require.Equal(t, result.Value, tc.expectedBack)
	}
}

// Если из очереди сделать слайс, то порядок элементов в нем будет соответсвовать порядку элементов в списке
func TestGetAllDoublyLinkedListItems(t *testing.T) {
	dll := NewDoublyLinkedList(10)
	expected := []interface{}{1, 3, 5, 7, 9}

	for _, value := range expected {
		dll.PushBack(value)
	}
	require.Equal(t, dll.All(), expected)
}

// Если из очереди удалить все элементы, то в очереде их не будет, вершина и хвост будут ссылаться на пустоту
func TestRemoveAllDoublyLinkedListItems(t *testing.T) {
	dll := NewDoublyLinkedList(10)
	input := []interface{}{1, 3, 5, 7, 9}

	items := make([]*DLListItem, 0)
	for _, value := range input {
		items = append(items, dll.PushBack(value))
	}
	for _, item := range items {
		dll.Remove(item)
	}

	front, err := dll.Front()
	require.Nil(t, front)
	require.Error(t, err)

	back, err := dll.Back()
	require.Nil(t, back)
	require.Error(t, err)

	require.Equal(t, dll.Len(), 0)
}
