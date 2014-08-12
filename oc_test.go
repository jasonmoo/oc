package oc

import (
	"fmt"
	"testing"
)

func ExampleInsertionOrder() {

	m := NewOc()

	m.Increment("key1", 1)
	m.Increment("key1", 1)
	m.Increment("key1", 1)
	m.Increment("key2", 3)
	m.Increment("key3", 2)
	m.Increment("key4", 1)

	for m.Next() {
		fmt.Println(m.KeyValue())
	}
	// Output:
	// key1 3
	// key2 3
	// key3 2
	// key4 1

}

func ExampleKeySortedOrder() {

	m := NewOc()

	m.Increment("key4", 1)
	m.Increment("key2", 3)
	m.Increment("key3", 2)
	m.Increment("key1", 1)

	m.SortByKey(ASC)

	for m.Next() {
		fmt.Println(m.KeyValue())
	}

	m.SortByKey(DESC)

	for m.Next() {
		fmt.Println(m.KeyValue())
	}
	// Output:
	// key1 1
	// key2 3
	// key3 2
	// key4 1
	// key4 1
	// key3 2
	// key2 3
	// key1 1

}

func ExampleCtSortedOrder() {

	m := NewOc()

	m.Increment("key4", 1)
	m.Increment("key2", 3)
	m.Increment("key3", 2)
	m.Increment("key1", 1)

	m.SortByCt(ASC)

	for m.Next() {
		fmt.Println(m.KeyValue())
	}

	m.SortByCt(DESC)

	for m.Next() {
		fmt.Println(m.KeyValue())
	}
	// Output:
	// key4 1
	// key1 1
	// key3 2
	// key2 3
	// key2 3
	// key3 2
	// key4 1
	// key1 1

}

func BenchmarkIncrement(b *testing.B) {

	m := NewOc()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Increment("key", 1)
	}
}
