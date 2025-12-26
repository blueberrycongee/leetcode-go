package lfucache

import "testing"

func TestLFUCache(t *testing.T) {
	// 测试用例1：基本操作
	lfu := Constructor(2)

	lfu.Put(1, 1) // freq: 1
	lfu.Put(2, 2) // freq: 1

	if got := lfu.Get(1); got != 1 { // freq: 1->2
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lfu.Put(3, 3) // 淘汰key 2，因为freq相同但更久未使用

	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	if got := lfu.Get(3); got != 3 { // freq: 1->2
		t.Errorf("Get(3) = %d, want 3", got)
	}

	if got := lfu.Get(1); got != 1 { // freq: 2->3
		t.Errorf("Get(1) = %d, want 1", got)
	}

	// 现在key 3的频次是2，key 1的频次是3，插入key 4会淘汰key 3
	lfu.Put(4, 4) // 淘汰key 3

	if got := lfu.Get(3); got != -1 {
		t.Errorf("Get(3) = %d, want -1", got)
	}

	if got := lfu.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLFUCache_UpdateValue(t *testing.T) {
	// 测试用例2：更新已存在的key
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1) // 访问后freq增加
	lfu.Put(1, 10) // 更新值，但保持相同的freq

	if got := lfu.Get(1); got != 10 {
		t.Errorf("Get(1) = %d, want 10", got)
	}
}

func TestLFUCache_SingleCapacity(t *testing.T) {
	// 测试用例3：容量为1
	lfu := Constructor(1)

	lfu.Put(1, 1)
	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lfu.Put(2, 2) // 淘汰key 1
	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}

	if got := lfu.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}

func TestLFUCache_EmptyCapacity(t *testing.T) {
	// 测试用例4：容量为0
	lfu := Constructor(0)

	lfu.Put(1, 1)
	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
}

func TestLFUCache_FreqPriority(t *testing.T) {
	// 测试用例5：验证频次优先级
	lfu := Constructor(2)

	lfu.Put(1, 1) // freq: 1
	lfu.Put(2, 2) // freq: 1

	// 访问key 1，使其freq变为2
	lfu.Get(1) // freq: 1->2

	// 插入key 3，应该淘汰key 2而不是key 1，因为key 2频次低且更久未使用
	lfu.Put(3, 3)

	// key 1应该还在，因为freq更高
	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	// key 2应该被淘汰
	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	// key 3应该存在
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}

// TestLFUCache_ComplexScenario 测试复杂的使用场景
func TestLFUCache_ComplexScenario(t *testing.T) {
	lfu := Constructor(3)

	// 添加三个元素
	lfu.Put(1, 1) // freq: 1
	lfu.Put(2, 2) // freq: 1
	lfu.Put(3, 3) // freq: 1

	// 访问1和2，增加它们的频次
	lfu.Get(1) // 1的freq: 1->2
	lfu.Get(2) // 2的freq: 1->2
	lfu.Get(1) // 1的freq: 2->3

	// 此时频次: 1(3), 2(2), 3(1) -> 1(freq=3), 2(freq=2), 3(freq=1)

	// 添加第四个元素，应该淘汰频次最低的(3)
	lfu.Put(4, 4)

	// 检查各元素状态
	if got := lfu.Get(1); got != 1 { // 应该存在
		t.Errorf("Get(1) = %d, want 1", got)
	}
	if got := lfu.Get(2); got != 2 { // 应该存在
		t.Errorf("Get(2) = %d, want 2", got)
	}
	if got := lfu.Get(3); got != -1 { // 应该被淘汰
		t.Errorf("Get(3) = %d, want -1", got)
	}
	if got := lfu.Get(4); got != 4 { // 应该存在
		t.Errorf("Get(4) = %d, want 4", got)
	}
}