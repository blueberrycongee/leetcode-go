package lfucache

// 题目：460. LFU 缓存
// 链接：https://leetcode.cn/problems/lfu-cache/
// 难度：Hard

// 思路：哈希表 + 双向链表 + 频次桶
// 1. 使用哈希表存储key到节点的映射，实现O(1)时间查找
// 2. 使用频次桶（每个频次对应一个双向链表）来维护相同频次的节点
// 3. 维护最小频次，以便快速找到需要淘汰的节点
// 4. 当访问节点时，将其从当前频次桶移动到下一频次桶

// 时间复杂度：get和put都是 O(1)
// 空间复杂度：O(capacity)

// 双向链表节点
type Node struct {
	key, value, freq int
	prev, next       *Node
}

// 双向链表
type DoublyLinkedList struct {
	head, tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (l *DoublyLinkedList) AddToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *DoublyLinkedList) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.head.next == l.tail
}

func (l *DoublyLinkedList) RemoveTail() *Node {
	if l.IsEmpty() {
		return nil
	}
	node := l.tail.prev
	l.RemoveNode(node)
	return node
}

type LFUCache struct {
	capacity   int
	size       int
	minFreq    int
	cache      map[int]*Node                    // key -> 节点
	freqMap    map[int]*DoublyLinkedList // 频次 -> 双向链表
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		cache:    make(map[int]*Node),
		freqMap:  make(map[int]*DoublyLinkedList),
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 更新节点频次
	this.increaseFreq(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.cache[key]; ok {
		// key存在，更新value并增加频次
		node.value = value
		this.increaseFreq(node)
		return
	}

	// key不存在，需要插入新节点
	if this.size >= this.capacity {
		// 容量已满，需要删除频次最低、最久未使用的节点
		this.removeMinFreqNode()
	}

	// 创建新节点
	node := &Node{key: key, value: value, freq: 1}
	this.cache[key] = node
	this.size++

	// 将节点添加到频次为1的链表中
	if _, ok := this.freqMap[1]; !ok {
		this.freqMap[1] = NewDoublyLinkedList()
	}
	this.freqMap[1].AddToHead(node)
	this.minFreq = 1
}

func (this *LFUCache) increaseFreq(node *Node) {
	// 从当前频次的链表中移除节点
	oldFreq := node.freq
	oldList := this.freqMap[oldFreq]
	oldList.RemoveNode(node)

	// 如果当前频次是最小频次且该频次链表为空，则更新最小频次
	if oldFreq == this.minFreq && oldList.IsEmpty() {
		this.minFreq++
		delete(this.freqMap, oldFreq)
	} else if oldList.IsEmpty() {
		delete(this.freqMap, oldFreq)
	}

	// 更新节点频次并添加到新的频次链表
	node.freq++
	newFreq := node.freq
	if _, ok := this.freqMap[newFreq]; !ok {
		this.freqMap[newFreq] = NewDoublyLinkedList()
	}
	this.freqMap[newFreq].AddToHead(node)
}

func (this *LFUCache) removeMinFreqNode() {
	// 获取最小频次的链表
	list := this.freqMap[this.minFreq]
	// 删除最久未使用的节点（链表尾部）
	tailNode := list.RemoveTail()
	if list.IsEmpty() {
		delete(this.freqMap, this.minFreq)
	}

	// 从缓存中删除该节点
	delete(this.cache, tailNode.key)
	this.size--
}