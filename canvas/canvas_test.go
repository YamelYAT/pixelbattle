package canvas

import (
	"testing"
)

type Data struct {
	x,
	y			int
	color		uint8
}

var(

	successData = []Data{
		{0,0,0},
		{100,300,24},
		{1590,400,24},
		{1590,400,0},
	}
	failData = []Data{
		{1591, 401, 25},
		{100, 401, 24},
		{1591, 100, 24},
		{1590,400,25},
	}

)

func TestCanvas_SetPoint(t *testing.T) {
	canvas := NewCanvas()

	for _, point := range successData{
		err := canvas.SetPoint(point.x, point.y, point.color)

		if err != nil{
			t.Errorf("set point %v failed: %v", point, err)
		}

		excepted := point.color
		returned := canvas.Points[point.x][point.y]

		if excepted != returned{
			t.Errorf("wrong color, got %v want %v. Point %v", returned, excepted, point)
		}
	}

	for _, point := range failData{
		err := canvas.SetPoint(point.x, point.y, point.color)
		if err == nil{
			t.Errorf("point %v unexcepted", point)
		}
	}

}