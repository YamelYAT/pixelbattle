package canvas

import (
	"testing"
)

type Data struct {
	x,
	y,
	color		int
}

var(
	successData = []Data{
		{0,0,0},
		{100,300,24},
		{1590,400,24},
		{1590,300,1},
	}
	failData = []Data{
		{1591, 401, 25},
		{100, 401, 24},
		{1591, 100, 24},
		{1590,401,25},
	}
	canvas = NewCanvas()
)

func TestCanvas_SetPoint_Success(t *testing.T) {
	for _, point := range successData{
		err := canvas.SetPoint(point.x, point.y, colors.Get(point.color))
		if err != nil{
			t.Errorf("set point %v failed: %v", point, err)
		}
	}
}

func TestCanvas_SetPoint_Fail(t *testing.T) {
	for _, point := range failData{
		err := canvas.SetPoint(point.x, point.y, colors.Get(point.color))
		if err == nil{
			t.Errorf("set point %v: %v", point, err)
		}
	}
}

func TestCanvas_GetPoint(t *testing.T) {
	for _, point := range successData{
		excepted := colors.Get(point.color)
		err, returned := canvas.GetPoint(point.x, point.y)

		if err != nil{
			t.Errorf("get point %v: %v", point, err)
		}

		if excepted != returned{
			t.Errorf("wrong color, got %v want %v. Point %v", returned, excepted, point)
		}
	}
}