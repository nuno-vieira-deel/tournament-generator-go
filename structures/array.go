package structures

/*
 * Dependencies
 */

import (
	"encoding/json"
	"fmt"
)

/*
 * Array structure
 */

type Array[T any] struct {
	DataStructure
	Data []T
}

/*
 * Constructor
 */

func NewArray[T any](list []T) Array[T] {
	array := *new(Array[T])
	array.Data = list

	return array
}

/*
 * Append
 */

func (a *Array[T]) Append(list []T) {
	a.Data = append(list, a.List()...)
}

/*
 * Get
 */

func (a *Array[T]) Get(i int) T {
	return a.Data[i]
}

/*
 * Length
 */

func (a *Array[T]) Length() int {
	return len(a.Data)
}

/*
 * List
 */

func (a *Array[T]) List() []T {
	return a.Data
}

/*
 * Pop
 */

func (a *Array[T]) Pop() T {
	item := a.Get(a.Length() - 1)
	a.Data = a.Data[0 : a.Length()-1]

	return item
}

/*
 * Print
 */

func (a *Array[T]) Print() {
	val, _ := json.Marshal(a.Data)

	fmt.Println(string(val))
}

/*
 * Push
 */

func (a *Array[T]) Push(item T) int {
	a.Data = append(a.Data, item)

	return a.Length() - 1
}

/*
 * Shift
 */

func (a *Array[T]) Shift() T {
	item := a.Get(0)
	a.Data = a.Data[1:a.Length()]

	return item
}

/*
 * Unshift
 */

func (a *Array[T]) Unshift(item T) {
	a.Data = append([]T{item}, a.List()...)
}
