package io

import (
	"fmt"
	"testing"
)

func TestMessage_UnMarshal(t *testing.T) {
	testByte := []byte{"15900400"}
	fmt.Println(testByte)
	p := UnMarshal(testByte)
	fmt.Println(p)
}
