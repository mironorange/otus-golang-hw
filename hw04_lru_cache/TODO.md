
- [x] Интерфейс двусвязанного списка
- [x] Структура двусвязанного списка
- [x] Структура элемента двусвязанного списка

## Описать интерфейс

```golang
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}
```

## Создать структуру данных

```golang
type DoublyLinkedList struct {
    capacity int
    head *DoublyLinkedItem
    tail *DoublyLinkedItem
}
type DoublyLinkedItem struct {
    Value interface{}
    Next *DoublyLinkedItem
    Prev *DoublyLinkedItem
}
```

## Запустить Unit Test из определенного файла

```bash
go test -v ./doublylinkedlist_test.go
go test -timeout 30s -run github.com/mironorange/otus-golang-hw/hw04_lru_cache
```

## Описать Unit Test

```golang
import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := "Hello, World!"
	expected := "Hello, World!"
	require.Equal(t, input, expected)
}
```

Ссылки:
- https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go-ru
