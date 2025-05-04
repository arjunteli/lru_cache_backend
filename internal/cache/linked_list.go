package cache

// DoublyLinkedList implements a doubly linked list for the LRU cache
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Node represents a node in the doubly linked list
type Node struct {
	item *CacheItem
	next *Node
	prev *Node
}

// NewDoublyLinkedList creates a new empty doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// AddToFront adds an item to the front of the list
func (dll *DoublyLinkedList) AddToFront(item *CacheItem) {
	node := &Node{item: item}
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
}

// MoveToFront moves an existing item to the front of the list
func (dll *DoublyLinkedList) MoveToFront(item *CacheItem) {
	node := dll.findNode(item)
	if node == dll.head {
		return
	}
	dll.removeNode(node)
	dll.AddToFront(item)
}

// Remove removes an item from the list
func (dll *DoublyLinkedList) Remove(item *CacheItem) {
	node := dll.findNode(item)
	dll.removeNode(node)
}

// RemoveOldest removes and returns the oldest item from the list
func (dll *DoublyLinkedList) RemoveOldest() *CacheItem {
	if dll.tail == nil {
		return nil
	}
	oldest := dll.tail.item
	dll.removeNode(dll.tail)
	return oldest
}

// findNode finds the node containing the given item
func (dll *DoublyLinkedList) findNode(item *CacheItem) *Node {
	for node := dll.head; node != nil; node = node.next {
		if node.item == item {
			return node
		}
	}
	return nil
}

// removeNode removes a node from the list
func (dll *DoublyLinkedList) removeNode(node *Node) {
	if node == nil {
		return
	}
	if node == dll.head {
		dll.head = node.next
		if dll.head != nil {
			dll.head.prev = nil
		}
	} else if node == dll.tail {
		dll.tail = node.prev
		if dll.tail != nil {
			dll.tail.next = nil
		}
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}
