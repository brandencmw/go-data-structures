package list

type AbstractList[T comparable] interface {
	Insert(T, int) int
	Append(T) int
	Prepend(T) int
	Remove(int) T
	Get(int) T
	Set(int, T)
	Equals(AbstractList[T]) bool
	Clone() *AbstractList[T]
	Contents() []T
}
