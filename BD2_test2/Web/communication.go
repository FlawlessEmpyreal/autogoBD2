package Web

import (
	"app/aStar"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Coord struct {
	Type string `json:"type"` // 固定填 "coord"
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Ts   int64  `json:"ts"`
}

// 低频发送：只在计算出新路径时发一次（大包）
type PathMsg struct {
	Type string        `json:"type"` // 固定填 "path"
	Path []aStar.Point `json:"path"`
	Ts   int64         `json:"ts"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Sender 封装UDP发送器
type Sender struct {
	conn *net.UDPConn //udp链接
	bx   int          //上次的x坐标
	by   int          //上次的y坐标
}

// 创建一个udp发送器,把坐标发到宿主机上，失败返回error
// 调用方式:NewSender("192.168.1.100:9999")
func NewSender(address string) (*Sender, error) {
	addr, err := net.ResolveUDPAddr("udp", address) //获取地址
	if err != nil {
		return nil, fmt.Errorf("地址解析失败: %w", err)
	}

	conn, err := net.DialUDP("udp", nil, addr) //和这个地址建立udp链接
	if err != nil {
		return nil, fmt.Errorf("UDP连接失败: %w", err)
	}

	return &Sender{conn: conn}, nil //返回链接
}

// Send 传入坐标，发送一次UDP包.建议使用 sync.Pool 复用对象,这样每次调用函数不用每次都内存操作
// (s *Sender)中的s相当于cpp中的this
func (s *Sender) SendCoord(x, y int) error {
	if s.bx == x && s.by == y { //和上次坐标相同的话就直接返回
		return nil
	}
	data, err := json.Marshal(Coord{
		Type: "coord",
		X:    x,
		Y:    y,
		Ts:   time.Now().UnixMicro(), //发送时间戳,用来检查延迟
	})
	if err != nil {
		return fmt.Errorf("json序列化坐标失败: %w", err)
	}
	//n, err := s.conn.Write(data)
	_, err = s.conn.Write(data)
	if err != nil {
		return fmt.Errorf("发送坐标失败: %w", err)
	}
	//fmt.Printf("发送: %s  字节数=%d  err=%v\n", string(data), n, err)
	s.bx = x
	s.by = y
	return nil
}

// 发送路径
func (s *Sender) SendPath(Path []aStar.Point) error {

	data, err := json.Marshal(PathMsg{
		Type: "path",
		Path: Path,
		Ts:   time.Now().UnixMicro(),
	})
	if err != nil {
		return fmt.Errorf("json序列化路径失败: %w", err)
	}

	n, err := s.conn.Write(data)
	fmt.Printf("发送路径: 字节数=%d  err=%v\n", n, err) // 先打印
	if err != nil {
		return fmt.Errorf("发送路径失败: %w", err) // 再判断错误
	}

	return nil
}

// Close 关闭连接
func (s *Sender) Close() {
	s.conn.Close()
}
