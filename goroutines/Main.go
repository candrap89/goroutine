package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	var num, k int

	// Input the first variable (num)
	num = 4

	// Input the second variable (wayPoints)
	wayPoints := make([]int, num)
	// for i := 0; i < num; i++ {
	// 		fmt.Scan(&wayPoints[i])
	// }
	wayPoints[0] = 1
	wayPoints[1] = 2
	wayPoints[2] = 3
	wayPoints[3] = 4

	// Input the third variable (k)
	k = 1

	// Convert array to linked list
	linkedList := arrayToLinkedList(wayPoints)

	// Remove kth element from the end
	linkedList = removeKthFromEnd(linkedList, k)

	// Print the modified linked list (for verification)
	printLinkedList(linkedList)

	//wg.Add(1) --> one because only have 1 go routines
	wg.Add(2)
	go struct_example()
	// to make delay
	go func() {
		fmt.Println("halo ganteng")
		wg.Done()
	}()
	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
}

var wg = sync.WaitGroup{}

type Doctor struct {
	number      int      //inisializer
	actorName   string   //inisializer
	compnanions []string //inisializer
}

func struct_example() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Candra Ganteng",
		compnanions: []string{
			"farida", "zafran", "sarwiyah",
		},
	}
	fmt.Println(aDoctor.actorName)
	wg.Done()
}

// Node struct for the linked list
type Node struct {
	Data int
	Next *Node
}

// Function to convert array to linked list
func arrayToLinkedList(arr []int) *Node {
	var head, current *Node

	for _, data := range arr {
		// Create a new node
		newNode := &Node{Data: data, Next: nil}

		// If it's the first node, set it as the head
		if head == nil {
			head = newNode
			current = head
		} else {
			// Link the current node to the new node
			current.Next = newNode
			// Move the current pointer to the new node
			current = newNode
		}
	}

	return head
}

// Function to print linked list (for verification)
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Printf("%d", head.Data)
		head = head.Next
	}
	fmt.Println()
}

// Function to remove kth element from the end of the linked list
func removeKthFromEnd(head *Node, k int) *Node {
	dummy := &Node{Next: head}
	fast := dummy
	slow := dummy

	// Move fast k steps ahead
	for i := 0; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	// Move both pointers until fast is at the end
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the kth node from the end
	slow.Next = slow.Next.Next

	return dummy.Next

}
