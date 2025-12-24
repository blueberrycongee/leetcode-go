package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	// 测试用例1：基本操作
	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lru.Put(3, 3) // 淘汰key 2
	if got := lru.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	lru.Put(4, 4) // 淘汰key 1
	if got := lru.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
	if got := lru.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
	if got := lru.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLRUCache_UpdateValue(t *testing.T) {
	// 测试用例2：更新已存在的key
	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(1, 10) // 更新key 1的值
	if got := lru.Get(1); got != 10 {
		t.Errorf("Get(1) = %d, want 10", got)
	}

	lru.Put(3, 3) // 淘汰key 2（因为key 1刚被访问过）
	if got := lru.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}
}

func TestLRUCache_SingleCapacity(t *testing.T) {
	// 测试用例3：容量为1
	lru := Constructor(1)

	lru.Put(1, 1)
	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lru.Put(2, 2) // 淘汰key 1
	if got := lru.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
	if got := lru.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}
