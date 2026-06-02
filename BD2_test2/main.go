package main

import "app/mainProcess"

func main() {

	mainProcess.MainProcess()

	//// 获取当前工作目录
	//dir, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("获取工作目录失败:", err)
	//	return
	//}
	//fmt.Println("当前工作目录:", dir)
}
