package main

import (
	//"container/heap"
	"fmt"
	//"container/list"
	//"container/list"
)

type Node struct {
	data interface{}
	//flag bool
	//distance int
	next *Node
}
type stack []string
func (s stack) Push(value string)stack{
	return append(s,value)
}
func (s stack) Pop()(stack, string){
	l:=len(s)
	return s[:l-1],s[l-1]
}

func printList(head *Node) {

	for ; head != nil; head = head.next {

		fmt.Println(head.data)

	}

	if head == nil {
		return
	}
}

func createNode(data interface{}) *Node {
	var node *Node
	node = &Node{}
	node.data = data
	//node.flag = false
	//node.distance = 0
	node.next = nil
	return node
}
func AddNode(data interface{}, head *Node) {
	for ; head.next != nil; head = head.next {

	}
	temp := createNode(data)
	head.next = temp
}
func arrayToList(linkArr []interface{}) *Node {
	head := createNode(linkArr[0])

	for i := 1; i < len(linkArr); i++ {
		AddNode(linkArr[i], head)
	}
	return head

}
func GetNth(head *Node, index int) interface{} {
	if head == nil {
		fmt.Println("Link list is empty")
		return -1
	}
	for i := 0; i < index; i++ {
		fmt.Println("head", head)
		fmt.Println("head.next", head.next)
		if head.next == nil {
			return -1
		}
		head = head.next
	}
	return head.data
}
func GetNthFromLast(head *Node, posfromlast int) interface{} {
	arr := make([]interface{}, 0)
	i := 0
	for ; head.next != nil; i++ {
		arr = append(arr, head.data)
		head = head.next
	}
	arr = append(arr, head.data)
	return arr[i-posfromlast+1]
}

var currentPos int

func GetNthFromLastRecursive(head *Node, posfromlast int) interface{} {

	if head.next == nil {
		currentPos++
		if currentPos == posfromlast {
			return head.data
		}
		return -1
	}
	res := GetNthFromLastRecursive(head.next, posfromlast)
	currentPos++

	if currentPos == posfromlast {
		fmt.Println("Came inside if ", currentPos, posfromlast, head.data)
		return head.data
	}
	return res
}
func LengthOfList(head *Node) int {
	length := 0
	for ; head != nil; head = head.next {
		length++
	}
	fmt.Println("Length of the list", length)
	return length
}

func GetMiddleOfList(head *Node) interface{} {
	array := make([]interface{}, 0)
	length := 0
	for ; head != nil; head = head.next {
		array = append(array, head.data)
		length++
	}
	index := length / 2
	return array[index]

}
func GetMiddleOfListDoublePointer(head *Node) interface{} {
	temp := head
	for head.next != nil {
		if head.next.next != nil {
			head = head.next.next
			temp = temp.next
		} else {
			head = head.next
			temp = temp.next
			break
		}
	}
	fmt.Println("Double pointer middle of the list", temp.data)
	return temp.data
}
func MaxNoOfTimesKeyOccurs(head *Node) (int, int) {
	myMap := make(map[int]int, 0)
	for ; head != nil; head = head.next {
		count, exists := myMap[head.data.(int)]
		if exists {
			count++
			myMap[head.data.(int)] = count
		} else {
			myMap[head.data.(int)] = 1
		}
	}
	fmt.Println(myMap)
	max := 0
	key := -1
	for index, value := range myMap {
		if value >= max {
			max = value
			key = index
		}
	}
	if max == 0 {
		fmt.Println("No element count is largest")
		return -1, -1
	}
	return key, max
}
func NoOfTimesKeyOccurs(head *Node, key int) int {
	count := 0
	for ; head != nil; head = head.next {
		if head.data == key {
			count++
		}
	}
	return count
}

// func loopDetectWithFlag(head *Node) int {
// 	for ; head != nil; head = head.next {
// 		if head.flag == true {
// 			fmt.Println("Loop is detected")
// 			return 1

// 		} else {
// 			head.flag = true
// 		}

// 	}

// 	return 0
// }
func loopDetectWithMap(head *Node) int {
	myMap := make(map[*Node]int, 0)
	for ; head != nil; head = head.next {
		_, exists := myMap[head.next]
		if exists {
			fmt.Println("Loop detected")
			return 1
		} else {
			myMap[head.next] = 1
		}

	}
	return 0
}

//Floyd's cycle finding algorithm
func loopDetectWithDoublePointer(head *Node) int {
	temp := head
	for head.next != nil {
		if head.next.next != nil {
			head = head.next.next
			temp = temp.next
			fmt.Println("head", head, "temp", temp)
			if head == temp {
				fmt.Println("Loop detected")
				return 1
			}
		} else {
			head = head.next
			temp = temp.next
			fmt.Println("head", head, "temp", temp)
			if head == temp {
				fmt.Println("Loop detected")
				return 1
			}

		}
	}
	return 0
}
func loopDetectSoln4(head *Node) int {
	temp := createNode(-1)

	if head == nil {
		return 0
	}

	for ; head != nil; head = head.next {
		if head.next == nil {
			fmt.Println("Loop detected")
			return 1
		}
		if head.next == temp {
			fmt.Println("Loop detected")
			return 1
		}
		head.next = temp

	}
	return 0
}
func loopDetectFindingDistanceSoln5Part(head *Node, last *Node) int {
	count := 0
	for ; head != last; head = head.next {
		count++
	}
	fmt.Println("Distance", count)
	return count
}
func loopDetectWithDistanceSoln5(head *Node) int {
	last := head
	var current_dist, prev_dist int
	for ; last != nil; last = last.next {
		if last == nil {
			return 0
		}

		prev_dist = current_dist

		current_dist = loopDetectFindingDistanceSoln5Part(head, last)

		if current_dist < prev_dist {
			fmt.Println("Loop detected")
			return 1
		}
	}
	return 0
}
func DetectLoopAndCount(head *Node) int {
	var count, prev_count, val int
	var exists bool
	myMap := make(map[*Node]int)
	for ; head != nil; head = head.next {
		val, exists = myMap[head.next]
		if exists {
			fmt.Println("Loop detected at", val)
			prev_count = count
			return prev_count - myMap[head.next] + 1
		} else {
			prev_count = count
			count++
			myMap[head.next] = count
		}
	}
	return 0
}

// func loopDetectAndCountUsingDistance(head *Node) int {
// 	count := 0
// 	prev_distance := 0
// 	for ; head != nil; head = head.next {
// 		if head.distance == 0 {
// 			prev_distance = count
// 			count++
// 			head.distance = count
// 		}
// 		if head.distance < prev_distance {
// 			prev_distance = count
// 			fmt.Println("Loop detected")
// 			return prev_distance - head.distance + 1

// 		}
// 	}
// 	return 0
// }
func detectLoopAndCountFastSLowPointers(head *Node) int {
	slow_ptr := head
	count := 0
	for head.next != nil {
		if head.next.next != nil {
			head = head.next.next
			slow_ptr = slow_ptr.next

		} else {
			head = head.next
			slow_ptr = slow_ptr.next
		}
		if head == slow_ptr {
			count++
			head = head.next
			for head != slow_ptr {
				count++
				head = head.next
			}
			return count
		}
	}

	return 0
}
func ListIsPalindrome(head *Node) bool {

	array := make([]string, 0)
	for ; head.next != nil; head = head.next {

		array = append(array, head.data.(string))
	}
	array = append(array, head.data.(string))
	fmt.Println("Arrray", array)
	var i, j int
	for i, j = 0, len(array)-1; i != j && i < j; i, j = i+1, j-1 {
		if array[i] == array[j] {
			fmt.Println("array[i]",array[i],"array[j]",array[j])
			continue
		} else {
			break
		}
	}
	if array[i] == array[j] {
		return true
	} else {
		return false
	}

}
func ListIsPalindromeWithStack(head *Node)int{
	stack:= make(stack,0)
	temp:=head
	var val string
	for;head!=nil;head=head.next{
		stack = stack.Push(head.data.(string))
	}
	for;temp!=nil;temp=temp.next{
		stack,val=stack.Pop()
		fmt.Println("temp.data",temp.data,"pop",val)
		if temp.data == val{
			continue
		}else{
			fmt.Println("Not a Palindrome")
			return 0
		}
	}
	fmt.Println("Its a Palindrome")
	return 1
	
}
func PalindromeWithReversingList(head *Node)bool{
}CREATION
	// head := createNode(1)
	// node2 := createNode(2)
	// node3 := createNode(3)
	// node4 := createNode(4)
	// node5 := createNode(5)
	// head.next = node2
	// node2.next = node3
	// node3.next = node4
	// node4.next = node5
	// node5.next = node2
	// //data := DetectLoopAndCount(head)
	// //data:=loopDetectAndCountUsingDistance(head)
	// data := detectLoopAndCountFastSLowPointers(head)
	// fmt.Println("Count", data)
	//loopDetectWithDistanceSoln5(head)
	// data:=loopDetectWithFlag(head)
	// loopDetectWithMap(head)
	//loopDetectWithDoublePointer(head)
	//loopDetectSoln4(head)
	// fmt.Println("Loop", data)
	//printList(head)

	//data:= GetMiddleOfList(head)
	//fmt.Println("Middle element of the list is ", data)
	//GetMiddleOfListDoublePointer(head)
	// key, value := MaxNoOfTimesKeyOccurs(head)
	// fmt.Println("key", key, "value", value)
	// count:= NoOfTimesKeyOccurs(head, 1)
	// fmt.Println("No of times key occurs", count)

	//data := GetNthFromLast(head, 3)
	//data := GetNthFromLastRecursive(head, 5)
	//fmt.Println("Nth data from the last", data)

	//fmt.Println("Deleting by position")
	//deleteNodeRecursive(head, 7)
	//deleteAtPos(&head, 8)

	//printList(head)

}
func deleteAtPos(phead **Node, pos int) {
	head := *phead
	if pos == 1 {
		if head.next == nil {
			*phead = nil
			return
		}
		*phead = head.next
		return
	}
	for i := 1; i < pos-1; i++ {

		head = head.next

	}
	if head.next.next == nil {
		head.next = nil
		return
	} else {
		head.next = head.next.next
		return
	}

}

func deleteNodeRecursive(head *Node, val int) {

	fmt.Println("came in deleteNOdeRecursive", head)

	if head.next == nil {
		fmt.Println("Element not present in the link list")
		return
	}

	if head.next.data == val {
		head.next = head.next.next
		return
	}

	// temp := head

	// if temp.next.data == val {
	// 	fmt.Println("Entered the data if statement")
	// 	head.next = temp.next.next
	// 	fmt.Println(head)
	// 	//head.next = temp.next
	// 	return
	// }

	deleteNodeRecursive(head.next, val)
}
