package linked_list

import (
	"fmt"
)

type Node struct {
  value int
  next  *Node
}

type LinkedList struct {
  head *Node
  tail *Node
  length int
}

func NewLinkedList(value int) LinkedList {
  node := Node{
      value: value,
      next: nil,
  };

  linkedList := LinkedList{
    head: &node,
    tail: &node,
    length: 1,
  };

  return linkedList;
}

func (list *LinkedList) Append(value int) {
  newNode := &Node{
    value: value,
    next: nil,
  };
  list.tail.next = newNode;
  list.tail = newNode;
  list.length++;
}

func (list *LinkedList) Prepend(value int) {
  newNode := &Node{
    value: value,
    next: list.head,
  }
  list.head = newNode;
  list.length++;
}

func (list *LinkedList) Insert(index, value int) {
  if index == 0 {
    list.Prepend(value);
  } else if index == list.length {
    list.Append(value);
  } else if index < 0 || index > list.length {
    fmt.Println("Not possible");
  } else {
    node := list.Traverse(index - 1);

    newNode := &Node{
      value: value,
      next: nil,
    }

    (*newNode).next = node.next;
    node.next = newNode;
    list.length++;
  }
}

func (list *LinkedList) Remove(index int) {
  if index < 0 || index >= list.length {
    fmt.Println("Not possible");
  } else if index == 0 {
    list.head = list.head.next;
    list.length--;
  } else {
    node := list.Traverse(index - 1);
    node.next = node.next.next;
    list.length--;
  }
}

func (list *LinkedList) Reverse() {
  prev := list.head;
  curr := list.head.next;

  for curr != nil {
    next := curr.next;
    curr.next = prev;
    prev = curr;
    curr = next;
  }

  list.head.next = nil;
  list.head, list.tail = list.tail, list.head;
}

func (list LinkedList) Traverse(index int) *Node {
  count := 0;
  node := list.head;

  for count < index {
    node = node.next;
    count++;
  }

  return node;
}

func (list LinkedList) PrintList() {
  values := []int{};
  node := list.head;
  for node != nil {
    values = append(values, (*node).value);
    node = (*node).next;
  }
  fmt.Println(values);
}