package canvas
import "C"
import (
	"errors"
)

type Canvas struct {
	Points 	[]byte
	Colors  color
	Width	int
	Height	int
}

func NewCanvas() *Canvas {
	c := &Canvas{
		Width: 1590,
		Height: 400,
	}

	for i:=0;i <= c.Width * c.Height + c.Height;i++{
		c.Points = append(c.Points, colors.GetRandomColor())
	}

	return c
}

func (c *Canvas) SetPoint(x int, y int, clr int) error{
	if x > c.Width || y > c.Height{
		return errors.New("wrong point position")
	}

	c.Points[x * c.Height + y] = colors.Get(clr)
	return nil
}

func (c *Canvas) GetPoint(x int, y int) (error, byte){
	if x > c.Width || y > c.Height{
		return errors.New("wrong point position"), 0
	}
	return nil, c.Points[x * c.Height + y]
}