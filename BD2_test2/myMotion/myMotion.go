// myMotion/motion.go
package myMotion

import (
	"math"

	"github.com/Dasongzi1366/AutoGo/motion"
)

const (
	ScreenCenterX = 640
	ScreenCenterY = 360
	HoldRadius    = 150 //保持点到屏幕中心的半径距离，这里设为 100。Hold 点最终会落在这个圆周上
)

var isHolding = false //表示当前是否处于“按住

func CalcHoldPoint(curX, curY, targetX, targetY int) (int, int) {
	dx := float64(targetX - curX)
	dy := float64(targetY - curY)
	length := math.Sqrt(dx*dx + dy*dy)
	if length > 0 {
		dx /= length
		dy /= length
	}
	return ScreenCenterX + int(dx*HoldRadius),
		ScreenCenterY + int(dy*HoldRadius)
}

func StartMoveXY(x, y int) {
	//fmt.Printf("isHolding=%v TouchDown/Move到(%d,%d)\n", isHolding, x, y)
	if !isHolding {
		motion.TouchDown(x, y, 0, 0)
		//println("touchdown: %d,%d", x, y)
		isHolding = true
	} else {
		motion.TouchMove(x, y, 0, 0)
	}
}

func StopMove() {
	if isHolding {
		motion.TouchUp(ScreenCenterX, ScreenCenterY, 0, 0)
		isHolding = false
	}
}
