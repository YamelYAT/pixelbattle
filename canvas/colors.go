package canvas

import (
	"math/rand"
)

type color [25]string

func (c *color)Get(num int) byte{
	if num >= len(c){
		return 0
	}
	return byte(c[num][0])
}

func (c *color)GetRandomColor() byte{
	return byte(c[rand.Intn(25)][0])
}

var colors = color{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"g",
	"k",
	"l",
	"m",
	"n",
	"o",
}
