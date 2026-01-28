// trees/tree_test.go
package trees

import (
	"reflect"
	"testing"
)

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================

// createTestTree создаёт стандартное тестовое дерево:
//
//	    a
//	   / \
//	  b   c
//	 / \   \
//	d   e   f
func createTestTree() *TreeNode[string] {
	a := &TreeNode[string]{Val: "a"}
	b := &TreeNode[string]{Val: "b"}
	c := &TreeNode[string]{Val: "c"}
	d := &TreeNode[string]{Val: "d"}
	e := &TreeNode[string]{Val: "e"}
	f := &TreeNode[string]{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	return a
}

// createIntTree создаёт дерево с int значениями
func createIntTree() *TreeNode[int] {
	root := &TreeNode[int]{Val: 1}
	root.Left = &TreeNode[int]{Val: 2}
	root.Right = &TreeNode[int]{Val: 3}
	root.Left.Left = &TreeNode[int]{Val: 4}
	root.Left.Right = &TreeNode[int]{Val: 5}
	root.Right.Right = &TreeNode[int]{Val: 6}
	return root
}

// ==================== ТЕСТЫ НА NIL БЕЗОПАСНОСТЬ ====================

func TestNilSafety(t *testing.T) {
	var root *TreeNode[string]

	// Все методы должны корректно работать с nil
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"IterativeDFS", root.IterativeDFS(), []string(nil)},
		{"RecursiveDFS", root.RecursiveDFS(), []string{}},
		{"BFS", root.BFS(), []string(nil)},
		{"IncludesIterativeDFS", root.IncludesIterativeDFS("x"), false},
		{"IncludesRecursiveDFS", root.IncludesRecursiveDFS("x"), false},
		{"IncludesBFS", root.IncludesBFS("x"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s() on nil = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

// ==================== ТЕСТЫ ОБХОДА ====================

func TestTreeNode_IterativeDFS(t *testing.T) {
	root := createTestTree()
	got := root.IterativeDFS()
	want := []string{"a", "b", "d", "e", "c", "f"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("IterativeDFS() = %v, want %v", got, want)
	}
}

func TestTreeNode_RecursiveDFS(t *testing.T) {
	root := createTestTree()
	got := root.RecursiveDFS()
	want := []string{"a", "b", "d", "e", "c", "f"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RecursiveDFS() = %v, want %v", got, want)
	}
}

func TestTreeNode_BFS(t *testing.T) {
	root := createTestTree()
	got := root.BFS()
	want := []string{"a", "b", "c", "d", "e", "f"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS() = %v, want %v", got, want)
	}
}

// Тест на одноэлементное дерево
func TestSingleNodeTree(t *testing.T) {
	root := &TreeNode[string]{Val: "root"}

	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{"IterativeDFS", root.IterativeDFS(), []string{"root"}},
		{"RecursiveDFS", root.RecursiveDFS(), []string{"root"}},
		{"BFS", root.BFS(), []string{"root"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

// ==================== ТЕСТЫ ПОИСКА ====================

func TestTreeNode_IncludesBFS(t *testing.T) {
	root := createTestTree()

	tests := []struct {
		target string
		want   bool
	}{
		{"a", true},  // корень
		{"b", true},  // уровень 2, слева
		{"c", true},  // уровень 2, справа
		{"d", true},  // уровень 3
		{"e", true},  // уровень 3
		{"f", true},  // уровень 3
		{"x", false}, // нет в дереве
		{"", false},  // пустая строка
	}

	for _, tt := range tests {
		t.Run("target="+tt.target, func(t *testing.T) {
			got := root.IncludesBFS(tt.target)
			if got != tt.want {
				t.Errorf("IncludesBFS(%q) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}

func TestTreeNode_IncludesIterativeDFS(t *testing.T) {
	root := createTestTree()

	tests := []struct {
		target string
		want   bool
	}{
		{"a", true},
		{"f", true},
		{"x", false},
	}

	for _, tt := range tests {
		got := root.IncludesIterativeDFS(tt.target)
		if got != tt.want {
			t.Errorf("IncludesIterativeDFS(%q) = %v, want %v", tt.target, got, tt.want)
		}
	}
}

func TestTreeNode_IncludesRecursiveDFS(t *testing.T) {
	root := createTestTree()

	tests := []struct {
		target string
		want   bool
	}{
		{"a", true},
		{"f", true},
		{"x", false},
	}

	for _, tt := range tests {
		got := root.IncludesRecursiveDFS(tt.target)
		if got != tt.want {
			t.Errorf("IncludesRecursiveDFS(%q) = %v, want %v", tt.target, got, tt.want)
		}
	}
}

// Тест на поиск в пустом поддереве
func TestSearchInPartialTree(t *testing.T) {
	root := &TreeNode[string]{Val: "a"}
	root.Left = &TreeNode[string]{Val: "b"}
	// Правое поддерево пустое

	if !root.IncludesBFS("a") {
		t.Error("Должен найти корень")
	}
	if !root.IncludesBFS("b") {
		t.Error("Должен найти левый потомок")
	}
	if root.IncludesBFS("c") {
		t.Error("Не должен найти в пустом правом поддереве")
	}
}

// ==================== ТЕСТЫ С РАЗНЫМИ ТИПАМИ ====================

func TestWithIntType(t *testing.T) {
	root := createIntTree()

	// Проверяем обход
	bfsResult := root.BFS()
	wantBFS := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(bfsResult, wantBFS) {
		t.Errorf("BFS with int = %v, want %v", bfsResult, wantBFS)
	}

	// Проверяем поиск
	if !root.IncludesBFS(4) {
		t.Error("Должен найти 4")
	}
	if root.IncludesBFS(99) {
		t.Error("Не должен найти 99")
	}
}

// ==================== ТЕСТЫ СТЕКА И ОЧЕРЕДИ ====================

func TestStack(t *testing.T) {
	s := Stack[int]{}

	// Push/Pop
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Pop() != 3 {
		t.Error("Pop() должен возвращать последний добавленный элемент")
	}
	if s.Pop() != 2 {
		t.Error("Pop() должен соблюдать LIFO")
	}
	if s.Pop() != 1 {
		t.Error("Pop() должен возвращать все элементы")
	}

	// Pop на пустом стеке
	if zero := s.Pop(); zero != 0 {
		t.Errorf("Pop() на пустом стеке должен возвращать zero value, got %v", zero)
	}

	// Empty
	if !s.Empty() {
		t.Error("Empty() должен возвращать true для пустого стека")
	}
	s.Push(42)
	if s.Empty() {
		t.Error("Empty() должен возвращать false для непустого стека")
	}
}

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	// Enqueue/Dequeue
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Pop() != 1 {
		t.Error("Pop() должен возвращать первый добавленный элемент (FIFO)")
	}
	if q.Pop() != 2 {
		t.Error("Pop() должен соблюдать FIFO")
	}
	if q.Pop() != 3 {
		t.Error("Pop() должен возвращать все элементы")
	}

	// Pop на пустой очереди
	if zero := q.Pop(); zero != 0 {
		t.Errorf("Pop() на пустой очереди должен возвращать zero value, got %v", zero)
	}

	// Empty
	if !q.Empty() {
		t.Error("Empty() должен возвращать true для пустой очереди")
	}
	q.Push(99)
	if q.Empty() {
		t.Error("Empty() должен возвращать false для непустой очереди")
	}
}

// ==================== ТЕСТЫ НА ДЕРЕВЬЯ С ДУБЛИКАТАМИ ====================

func TestTreeWithDuplicates(t *testing.T) {
	// Дерево с повторяющимися значениями
	root := &TreeNode[string]{Val: "a"}
	root.Left = &TreeNode[string]{Val: "b"}
	root.Right = &TreeNode[string]{Val: "b"} // дубликат

	// Поиск должен находить оба
	if !root.IncludesBFS("b") {
		t.Error("Должен найти дубликат")
	}
}

// ==================== BENCHMARK ТЕСТЫ ====================

func BenchmarkIterativeDFS(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = root.IterativeDFS()
	}
}

func BenchmarkRecursiveDFS(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = root.RecursiveDFS()
	}
}

func BenchmarkBFS(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = root.BFS()
	}
}

func BenchmarkIncludesBFS(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = root.IncludesBFS("f") // поиск глубокого элемента
	}
}

func BenchmarkIncludesIterativeDFS(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = root.IncludesIterativeDFS("f")
	}
}
