package main

import "fmt"


//declear a new type struct
type MaxHeap struct {
	array []int
}

//append data to the struct
func (h *MaxHeap) Insert(key int){
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

//heapfy

func (h *MaxHeap) maxHeapifyUp(index int) {
	//swap whent the index is larger than parent
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// we need to get the parent
func parent (i int) int {
	return (i-1)/2
}

// get the left index
func left (i int) int {
	return 2 *i + 1
}

//get the right child index
func right (i int) int {
	return 2*i + 2
}

// swap the two values
func (h *MaxHeap) swap (i1,i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}


func main (){
	m := &MaxHeap{}
	fmt.Println(m)

	// create a heap
	buildHeap := []int{10, 20, 30, 15, 50, 33,34}

	// insert the heap into struct
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	//sort heap
}