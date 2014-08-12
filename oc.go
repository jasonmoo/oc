package oc

import (
	"container/list"
	"unicode/utf8"
)

type (
	Oc struct {
		set  map[string]*list.Element
		list *list.List

		cur *list.Element
	}

	element struct {
		key string
		ct  int64
	}
)

const (
	ASC  = 1
	DESC = -1
)

func NewOc() *Oc {
	return &Oc{
		set:  make(map[string]*list.Element),
		list: list.New(),
	}
}

func (o *Oc) Increment(key string, val int64) {

	if el, exists := o.set[key]; exists {
		el.Value.(*element).ct += val
	} else {
		o.set[key] = o.list.PushBack(&element{key: key, ct: val})
	}

}

func (o *Oc) Decrement(key string, val int64) {
	o.Increment(key, -val)
}

func (o *Oc) Delete(key string) {
	if el, exists := o.set[key]; exists {
		o.list.Remove(el)
		delete(o.set, key)
	}
}

func (o *Oc) Len() int {
	return len(o.set)
}

func (o *Oc) Next() bool {

	// first time through
	if o.cur == nil {
		o.cur = o.list.Front()
		return true
	}

	o.cur = o.cur.Next()

	return o.cur != nil

}

func (o *Oc) KeyValue() (string, int64) {
	e := o.cur.Value.(*element)
	return e.key, e.ct
}

func (o *Oc) SortByKey(dir int) {

	cursor := o.list.Front()

	for cursor != nil {

		// grab prev to process and next so we don't lose our place
		prev, next := cursor.Prev(), cursor.Next()

		// move backward until a cmp has been found
		for prev != nil && strcmp(prev.Value.(*element).key, cursor.Value.(*element).key)*dir > 0 {
			prev = prev.Prev()
		}

		if prev == nil {
			o.list.MoveToFront(cursor)
		} else if prev != cursor.Prev() {
			o.list.Remove(cursor)
			o.list.InsertAfter(cursor.Value, prev)
		}

		cursor = next

	}

}

func (o *Oc) SortByCt(dir int) {

	cursor := o.list.Front()
	d := int64(dir)

	for cursor != nil {

		// grab prev to process and next so we don't lose our place
		prev, next := cursor.Prev(), cursor.Next()

		// move backward until a cmp has been found
		for prev != nil && (prev.Value.(*element).ct-cursor.Value.(*element).ct)*d > 0 {
			prev = prev.Prev()
		}

		if prev == nil {
			o.list.MoveToFront(cursor)
		} else if prev != cursor.Prev() {
			o.list.Remove(cursor)
			o.list.InsertAfter(cursor.Value, prev)
		}

		cursor = next

	}

}

func strcmp(a, b string) int {

	for len(a) > 0 && len(b) > 0 {
		ra, sizea := utf8.DecodeRuneInString(a)
		rb, sizeb := utf8.DecodeRuneInString(b)
		if ra != rb {
			return int(ra - rb)
		}
		a, b = a[sizea:], b[sizeb:]
	}

	// return the shorter
	if len(a) == 0 {
		if len(b) == 0 {
			return 0
		}
		return -1
	}
	return 1

}
