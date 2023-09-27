package main
import(
	"fmt"
	
)
type Node struct{
	property int
	nextNode *Node
}
type LinkedList struct{
	headNode *Node
}

//AddToHead method
func (linkedList *LinkedList) AddToHead(property int){
	var node= Node{}
	node.property=property
	if node.nextNode !=nil{
		node.nextNode=linkedList.headNode
	}
	linkedList.headNode=&node
}
//NodeWithValue method 
func (linkedList *LinkedList) NodeWithValue(property int) *Node{
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
	if node.property == property {
	nodeWith = node
	break;
	}
	}
	return nodeWith
	}
//AddAfter method
func (linkedList *LinkedList) AddAfter(nodeProperty int,property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
	node.nextNode = nodeWith.nextNode
	nodeWith.nextNode = node
	}
	}
//LastNode
func (linkedList *LinkedList) LastNode() *Node{
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
	if node.nextNode ==nil {
	lastNode = node
	}
	}
	return lastNode
	}
//AddToEnd 
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
	lastNode.nextNode = node
	}
	}
//IterateList 
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
	fmt.Println(node.property)
	}
	}

//main
func main(){
	var linkedList LinkedList
	linkedList=LinkedList{}
	linkedList.AddToHead(-10);
	linkedList.AddToHead(5);
	fmt.Println(linkedList.headNode.property)
	linkedList.AddToEnd(50)
	linkedList.AddAfter(5,33)
	linkedList.LastNode();
	linkedList.NodeWithValue(50);
	linkedList.IterateList()
}
