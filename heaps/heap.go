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

// extract, return the largest key and remove it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	if len(h.array) == 0 {
		fmt.Println("Cant extract because the length of the array is 0")
		return -1
	}

	l := len(h.array)-1
	h.array[0] =h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array)-1
	l,r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		// when left child is only child
		if l == lastIndex {
			childToCompare = l
		}else if h.array[l] > h.array[r] {
			childToCompare = l //when left child is larger
		} else {
			childToCompare = r //when right child is larger
		}
		
		// compare array value and swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		}else{
			return
		}
		
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
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	// insert the heap into struct
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	//sort heap

	for i := 0; i <= 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}