/*
@Time : 2019/11/13 14:53
@Author : yanKoo
@File : bst_test
@Software: GoLand
@Description: bst test
*/
package bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBST(t *testing.T) {
	bst := NewBST()
	arr := []int{40, 5, 42, 6, 10, 400, 77}
	for _, e := range arr {
		bst.Add(e)
	}
	assert.Equal(t, bst.Size(), len(arr))
	assert.Equal(t, bst.Search(4), false)
	assert.Equal(t, bst.Search(40), true)
	bst.Remove(40)
	assert.Equal(t, bst.Search(40), true)
	assert.Equal(t, bst.Size(), len(arr)-1)
	fmt.Println(bst.RemoveMax())
	assert.Equal(t, bst.Size(), len(arr)-2)
	fmt.Println(bst.RemoveMin())
	assert.Equal(t, bst.Size(), len(arr)-3)

}
