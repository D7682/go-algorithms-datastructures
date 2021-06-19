package main

import (
	"fmt"
)

const ArraySize = 7
// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket // of type pointer bucket, so it needs to point to a bucket type/struct..
}

// bucket structure - bucket will be the head of the linked list.
type bucket struct {
	head *bucketNode // the head variable needs to point to a bucketNode type/struct.
}

// bucketNode -- will hold the key and an address to the next node
type bucketNode struct {
	key string // bucketNode stores a key
	next *bucketNode // and a pointer to another Node's memory address to form a linked list. The node points to a Node of type bucketNode.
}
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	// the insert method below is a method of the bucket type h.array[0] remember that hashtable has a pointer to a bucket type.
	h.array[index].insert(key) // with this we go into the index of the array that the key is stored into.
}
// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert
func (b *bucket) insert(k string) {
	// in the if statement were using the true/false return value from the bucket's search method.
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}
// search - will return false of true which we will use in the bucket's insert method.
func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete will take in a key and delete the node from the bucket
func (b *bucket) delete (k string) {

	// we first check if the key matches the head of the bucket's key
	// if it is true then the next node would be the new head of the bucket and end the code here.
	// if it is false then the code will continue to the for loop.
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	headNode := b.head
	for headNode.next != nil {
		if headNode.next.key == k {
			// delete
			// ---------------------------- bucket ----------------------------
			// headNode --> headNode.next --> headNode.next.next --> etc.
			// headNode.next.next will be moved to the position of headNode.next if headNode.next.key's value matches the one we want to delete.
			headNode.next = headNode.next.next
			return
		}
		// if headNode.next.key is not equal to the key we want to delete then
		// we will make the new headNode be the headNode.next node in order to use it as a point of reference when executing the for loop again.
		headNode = headNode.next
	}

}


// hash
func hash(key string) int {
	sum := 0
	for _, value := range key {
		sum += int(value)
	}
	return sum % ArraySize
}

// Init - this init function will add bucket nodes to each array index of the hash table we initialized in the main function.
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{} // for each element in the array we will store a "bucket" type *remember* they have a head paramaeter.
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

	hashTable.Delete("STAN")

}