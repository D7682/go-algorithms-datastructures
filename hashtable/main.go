package main

import (
	"fmt"
)

//ArraySize iss the size of the hash table array
const ArraySize = 7

// HashTable will hold an array.
type HashTable struct {
	array[ArraySize] *bucket
}
// bucket structure - needs to be a linked list.
type bucket struct {
	head *bucketNode // bucket will be the head of the bucket.
}
// bucket node structure
type bucketNode struct {
	key string // for the names
	next *bucketNode // next will point to the address of the next bucket node
}
// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
// Search will take in a key and return true if tht key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert
func (b *bucket) insert(k string) {
	// if a value we pass doesn't exist this conditional creates the new node to store the name in the key
	// if it does exist, then the else statement is executed.
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}
// search will take in a key and return true if the bucket has a match.
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete will take in a key and delete the node from the bucket.
func (b *bucket) delete(k string) {
	// if the key of the head matches the value of what were trying to delete we execute this first statement
	if b.head.key == k {
		// we delete the head by making the next node (b.head.next) the new head node.
		b.head = b.head.next
		return
	}

	// b.head will change any time the node after this node is not the key were searching for.
	previousNode := b.head
	// for loop executes if the node after the b.head is not a node that doesn't exist.
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
			return
		}
		// b.head = b.head.next
		previousNode = previousNode.next
	}
}

// hash - hash function
func hash(key string) int {
	sum := 0
	// will loop through each of the characters in a string.
	for _, value := range key {
		sum += int(value)
	}

	return sum % ArraySize
}
// Init - initializes the hash table and creates a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i, _ := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}


func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable.array[0].head.key)

	fmt.Println(hashTable.Search("KENNY"))
	fmt.Println(hashTable.Search("KYLE"))
	fmt.Println(hashTable.Search("STAN"))
	fmt.Println(hashTable.Search("CHRIS"))
	hashTable.Delete("BUTTERS")

	fmt.Println(hashTable.Search("BUTTERS"))
}