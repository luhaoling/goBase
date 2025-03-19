package linkList

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	doubleLinkList := NewDoubleLinkList()
	doubleLinkList.InsertAtHead(1)
	doubleLinkList.InsertAtHead(2)
	doubleLinkList.InsertAtTail(3)
	doubleLinkList.InsertAtTail(4)
	doubleLinkList.InsertAtRandomPosition(2, 5)
	doubleLinkList.PrintListForward()
	fmt.Println("--------------")
	doubleLinkList.PrintListBackward()
	doubleLinkList.DeleteAtPosition(2)
	fmt.Println("--------------")
	doubleLinkList.PrintListForward()

}
