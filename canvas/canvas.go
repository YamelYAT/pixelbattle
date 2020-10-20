package canvas
import "C"
import (
	"errors"
)

type Canvas struct {
	Points [1591][401]byte
}

func NewCanvas() *Canvas {
	var c Canvas
	for i, a := range c.Points{
		for j, _ := range a{
			c.Points[i][j] = 0
		}
	}
	return &c
}

func (c *Canvas) SetPoint(x int, y int, color uint8) error{
	if x >= len(c.Points) || y >= len(c.Points[0]){
		return errors.New("wrong point position")
	}
	if color > 24{
		return errors.New("color not exist")
	}

	c.Points[x][y] = color

	return nil
}