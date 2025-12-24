package lrucache

// 题目：146. LRU 缓存
// 链接：https://leetcode.cn/problems/lru-cache/
// 难度：Medium

// 思路：哈希表 + 双向链表
// 1. 哈希表：O(1)时间查找key对应的节点
// 2. 双向链表：O(1)时间移动节点到头部、删除尾部节点
// 3. 最近使用的放在头部，最久未使用的在尾部

// 时间复杂度：get和put都是 O(1)
// 空间复杂度：O(capacity)

// 双向链表节点
type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node // key -> 节点
	head     *Node         // 哨兵头节点（最近使用）
	tail     *Node         // 哨兵尾节点（最久未使用）
}

func Constructor(capacity int) LRUCache {
	// 创建哨兵节点，简化边界处理
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// 移动到头部（标记为最近使用）
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// key存在，更新value并移到头部
		node.value = value
		this.moveToHead(node)
	} else {
		// key不存在，创建新节点
		node := &Node{key: key, value: value}
		this.cache[key] = node
		this.addToHead(node)

		// 超出容量，删除尾部节点
		if len(this.cache) > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
		}
	}
}

// 添加节点到头部
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 删除节点
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移动节点到头部
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除尾部节点
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
