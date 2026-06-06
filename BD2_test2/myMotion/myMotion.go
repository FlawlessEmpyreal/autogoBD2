// myMotion/motion.go
package myMotion

import (
	"math"
	"math/rand"
	"time"

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

func RandomPoint() (x1, y1, x2, y2 int) { //角色周围随机生成坐标
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	theta := r.Float64() * 2 * math.Pi

	// 计算第一次点击的坐标
	x1 = 640 + int(float64(150)*math.Cos(theta))
	y1 = 360 + int(float64(150)*math.Sin(theta))

	// 计算相反方向的坐标 (角度 + Pi)
	x2 = 640 + int(float64(150)*math.Cos(theta+math.Pi))
	y2 = 360 + int(float64(150)*math.Sin(theta+math.Pi))
	return x1, y1, x2, y2
}
