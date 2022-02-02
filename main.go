package main

import (
	"data-structures-go/linked_list"
)

func main() {
  list := linked_list.NewLinkedList(10);
  list.Append(5);
  list.Prepend(12);
  list.Insert(2, 2);
  list.PrintList();
  list.Reverse();
  list.PrintList();
}
