// Implement Queue using Stacks
//
// [Easy] [AC:68.7% 303.8K of 442.3K] [filetype:golang]
//
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
// void push(int x) 将元素 x 推到队列的末尾
//
// int pop() 从队列的开头移除并返回元素
//
// int peek() 返回队列开头的元素
//
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
//
// 说明：
//
// 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size,
// 和 is empty 操作是合法的。
//
// 你所使用的语言也许不支持栈。你可以使用 list 或者
// deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
// 示例 1：
//
// 输入：
//
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
//
// [[], [1], [2], [], [], []]
//
// 输出：
//
// [null, null, null, 1, 1, false]
//
// 解释：
//
// MyQueue myQueue = new MyQueue();
//
// myQueue.push(1); // queue is: [1]
//
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//
// myQueue.peek(); // return 1
//
// myQueue.pop(); // return 1, queue is [2]
//
// myQueue.empty(); // return false
//
// 提示：
//
// 1 <= x <= 9
//
// 最多调用 100 次 push、pop、peek 和 empty
//
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
//
// 进阶：
//
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n
// 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
//
// [End of Description]
package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	x := 100
	obj.Push(x)
	param_3 := obj.Peek()
	param_2 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)

}

// local end

type MyQueue struct {
	l1  *Stack
	tmp *Stack
}

type Stack struct {
	L list.List
}

func (this *Stack) push(x int) {
	this.L.PushBack(x)
}

func (this *Stack) pop() int {
	tmp := this.L.Remove(this.L.Back())
	return tmp.(int)
}

func (this *Stack) peek() int {
	tmp := this.L.Back()
	return tmp.Value.(int)
}

func (this *Stack) empty() bool {
	len := this.L.Len()
	return len == 0
}

func Constructor() MyQueue {
	myQuere := MyQueue{}
	myQuere.l1 = &Stack{}
	myQuere.tmp = &Stack{}

	return myQuere

}

func (this *MyQueue) Push(x int) {
	//empty tmp stack
	for this.l1.empty() != true {
		node := this.l1.pop()
		this.tmp.push(node)
	}
	this.tmp.push(x)

	for this.tmp.empty() != true {
		node := this.tmp.pop()
		this.l1.push(node)
	}
	return
}

func (this *MyQueue) Pop() int {
	return this.l1.pop()
}

func (this *MyQueue) Peek() int {

	return this.l1.peek()
}

func (this *MyQueue) Empty() bool {
	return this.l1.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
