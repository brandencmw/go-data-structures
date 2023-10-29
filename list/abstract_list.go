package list

type AbstractList[T any] interface {
	Insert(T, int) int
	Append(T) int
	Prepend(T) int
	Remove(int) T
	Get(int) T
	Set(T, int)
}
