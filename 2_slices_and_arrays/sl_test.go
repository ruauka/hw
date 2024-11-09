package __slices_and_arrays

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Кольцевая очередь (Circular Queue) — это структура данных, которая представляет собой очередь (FIFO) фиксированного размера.
// Кольцевая очередь использует буфер фиксированного размера таким образом,
// как будто бы после последнего элемента сразу же снова идет первый (как представлено на картинке справа).

// go test -v homework_test.go

type CircularQueue struct {
	values     []int
	size       int // кол-во значений в очереди
	head, tail int // индексы головы и хвоста очереди
}

// создать очередь с определенным размером буффера
func NewCircularQueue(capacity int) CircularQueue {
	return CircularQueue{
		values: make([]int, capacity),
	}
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	// добавляем новое значение
	q.values[q.tail] = value
	q.tail = (q.tail + 1) % cap(q.values)
	q.size++
	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % cap(q.values)
	q.size--
	return true
}

// получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.head]
}

// получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	if q.tail == 0 {
		return q.values[cap(q.values)-1]
	}
	return q.values[q.tail-1]
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.size == 0
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return cap(q.values) == q.size
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
