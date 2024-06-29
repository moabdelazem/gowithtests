package structs

import (
	"math"
	"testing"
)

// Shape Interface
type Shape interface {
	Area() float64
}

// Shape Structs
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.shape.Area()
			if output != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, output, tt.hasArea)
			}
		})
	}
}

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (rect Rectangle) Perimeter() float64 {
// 	return 2 * (rect.Width + rect.Height)
// }

// func (rect Rectangle) Area() float64 {
// 	return rect.Width * rect.Height
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.Radius
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius
// }

// func TestPerimeter(t *testing.T) {
// 	t.Run("calc the perimeter of rectangle", func(t *testing.T) {
// 		testRect := Rectangle{10.0, 10.0}
// 		got := testRect.Perimeter()
// 		want := 40.0

// 		if got != want {
// 			t.Errorf("got %v want %v", got, want)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("calc the area of rectangle", func(t *testing.T) {
// 		testRect := Rectangle{12.0, 6.0}
// 		want := 72.0

// 		checkArea(t, testRect, want)
// 	})

// 	t.Run("calc the area of circle", func(t *testing.T) {
// 		testCircle := Circle{
// 			10,
// 		}
// 		want := 314.1592653589793

// 		checkArea(t, testCircle, want)
// 	})
// }
