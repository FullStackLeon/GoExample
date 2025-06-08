package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
)

type QueueType int

const (
	QueueTypeMin QueueType = 0
	QueueTypeMax QueueType = 1
)

// QueueItem 定义队列项
type QueueItem struct {
	value    string
	priority int
	index    int
}

// SetIndex 设置队列项的索引
func (qi *QueueItem) SetIndex(index int) {
	qi.index = index
}

// PriorityQueue 定义优先队列
type PriorityQueue struct {
	heapType QueueType // 0: 小根堆, 1: 大根堆
	items    []*QueueItem
	mu       sync.RWMutex
}

// 实现 heap.Interface 接口的 Len 方法
func (pq PriorityQueue) Len() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return len(pq.items)
}

// 实现 heap.Interface 接口的 Less 方法
func (pq PriorityQueue) Less(i, j int) bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	if pq.heapType == 0 {
		// 小根堆：优先级越小，越靠前
		return pq.items[i].priority < pq.items[j].priority
	} else {
		// 大根堆：优先级越大，越靠前
		return pq.items[i].priority > pq.items[j].priority
	}
}

// 实现 heap.Interface 接口的 Swap 方法
func (pq PriorityQueue) Swap(i, j int) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	// 更新索引 i, j位置的QueueItem交换后，j位置的index还是i，i位置的index还是j，需要将i,j位置QueueItem的index更新为i,j的正确索引
	pq.items[i].SetIndex(i)
	pq.items[j].SetIndex(j)
}

// 实现 heap.Interface 接口的 Push 方法
func (pq *PriorityQueue) Push(x any) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	n := len(pq.items)
	item := x.(*QueueItem)
	item.SetIndex(n)
	pq.items = append(pq.items, item)
}

// 实现 heap.Interface 接口的 Pop 方法
func (pq *PriorityQueue) Pop() any {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil    // 避免内存泄漏
	item.SetIndex(-1) // 无效索引
	pq.items = old[0 : n-1]
	return item
}

// NewQueue 创建并返回一个新的 OrderQueue 实例
func NewQueue(queueType QueueType) *OrderQueue {
	// 优先队列初始化
	pq := PriorityQueue{
		heapType: queueType,
		items:    make([]*QueueItem, 0),
		mu:       sync.RWMutex{},
	}

	// 切片转为堆结构
	heap.Init(&pq)

	// 订单队列初始化
	queue := OrderQueue{
		pq: &pq,
		m:  make(map[string]*QueueItem),
	}
	return &queue
}

// OrderQueue 包含一个优先队列和一个映射，用于存储和管理订单队列
type OrderQueue struct {
	pq *PriorityQueue
	m  map[string]*QueueItem
}

func main() {
	// 小根堆 PriorityQueue
	queueMin := NewQueue(QueueTypeMin)
	order1 := &QueueItem{value: "order1", priority: 5}
	order2 := &QueueItem{value: "order2", priority: 10}
	order3 := &QueueItem{value: "order3", priority: 1}

	heap.Push(queueMin.pq, order1)
	heap.Push(queueMin.pq, order2)
	heap.Push(queueMin.pq, order3)
	for queueMin.pq.Len() > 0 {
		order := heap.Pop(queueMin.pq).(*QueueItem)
		fmt.Printf("queueMin Popped order: %s with priority %d\n", order.value, order.priority)
	}

	// 大根堆 PriorityQueue
	queueMax := NewQueue(QueueTypeMax)

	heap.Push(queueMax.pq, order1)
	heap.Push(queueMax.pq, order2)
	heap.Push(queueMax.pq, order3)
	for queueMax.pq.Len() > 0 {
		order := heap.Pop(queueMax.pq).(*QueueItem)
		fmt.Printf("queueMax Popped order: %s with priority %d\n", order.value, order.priority)
	}

	// 并发测试优先队列
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 4; i < 10; i++ {
			order := &QueueItem{value: fmt.Sprintf("order%d", i), priority: i + rand.Intn(100)}
			heap.Push(queueMin.pq, order)
			queueMin.m[order.value] = order
		}
	}()

	go func() {
		wg.Done()
		for i := 10; i < 15; i++ {
			order := &QueueItem{value: fmt.Sprintf("order%d", i), priority: i + rand.Intn(100)}
			heap.Push(queueMax.pq, order)
			queueMax.m[order.value] = order
		}
	}()

	wg.Wait()

	for queueMin.pq.Len() > 0 {
		order := heap.Pop(queueMin.pq).(*QueueItem)
		fmt.Printf("queueMin Popped order: %s with priority %d\n", order.value, order.priority)
	}

	for queueMax.pq.Len() > 0 {
		order := heap.Pop(queueMax.pq).(*QueueItem)
		fmt.Printf("queueMax Popped order: %s with priority %d\n", order.value, order.priority)
	}
}
