package main

type node struct {
	Val  interface{}
	Prev *node
}
type stackLink struct {
	Top    *node
	Length int
}

func (sl *stackLink) push(val interface{}) {
	newNode := &node{
		Val:  val,
		Prev: sl.Top,
	}
	sl.Top = newNode
	sl.Length++
}

func (sl *stackLink) pop() interface{} {
	tempVal := sl.Top.Val
	sl.Top = sl.Top.Prev
	sl.Length--
	return tempVal
}

func (sl stackLink) length() int {
	return sl.Length
}

func (sl stackLink) isEmpty() bool {
	return sl.Length == 0
}

func (sl stackLink) peer() interface{} {
	return sl.Top.Val
}
