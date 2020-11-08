package canvas

import "testing"

var(
	ExceptedColors = map[int]byte{
		0:"0"[0],
		1:"1"[0],
		2:"2"[0],
		3:"3"[0],
		4:"4"[0],
		5:"5"[0],
		6:"6"[0],
		7:"7"[0],
		8:"8"[0],
		9:"9"[0],
		10:"a"[0],
		11:"b"[0],
		12:"c"[0],
		13:"d"[0],
		14:"e"[0],
		15:"f"[0],
		16:"g"[0],
		17:"h"[0],
		18:"i"[0],
		19:"j"[0],
		20:"k"[0],
		21:"l"[0],
		22:"m"[0],
		23:"n"[0],
		24:"o"[0],
	}
	NonExceptedColors = []int{-9999,25,9999, -1, 256}
)

func TestColor_GetExistingColors(t *testing.T) {
	for num, color := range ExceptedColors{
		excepted := color
		returned := colors.Get(num)
		if excepted != returned{
			t.Errorf("wrong color got %v want %v", returned, excepted)
		}
	}
}

func TestColor_GetNonExistingColors(t *testing.T) {
	for _, colorNum := range NonExceptedColors{
		excepted := "0"[0]
		returned := colors.Get(colorNum)
		if excepted != returned{
			t.Errorf("wrong color got %v want %v", returned, excepted)
		}

	}
}