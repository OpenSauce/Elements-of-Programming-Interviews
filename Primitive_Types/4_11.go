package chapter_4

import "math"

type Rect struct {
	x, y, height, width int
}

func RectangleIntersection(rect1, rect2 Rect) Rect {
	if rect1.x <= rect2.x+rect2.width && rect2.x <= rect1.x+rect1.width &&
		rect1.y <= rect2.y+rect2.height && rect2.y <= rect1.y+rect1.height {

		var x, y, height, width int
		x = int(math.Max(float64(rect1.x), float64(rect2.x)))
		y = int(math.Max(float64(rect1.y), float64(rect2.y)))
		height = int(math.Min(float64(rect1.y+rect1.height), float64(rect2.y+rect2.height))) - y
		width = int(math.Min(float64(rect1.x+rect1.width), float64(rect2.x+rect2.width))) - x
		return Rect{x, y, height, width}
	}

	return Rect{}
}
