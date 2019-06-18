package main

import (
	"tidy-golang/lock"
)

// channel详解 https://colobu.com/2016/04/14/Golang-Channels/
// 协程详解 https://www.cnblogs.com/liang1101/p/7285955.html
// iops (Input/Output Operations Per Second) 是一个用于计算机存储设备（如硬盘（HDD）、固态硬盘（SSD）或存储区域网络（SAN））性能测试的量测方式，可以视为是每秒的读写次数

func main() {
	lock.DoGoLock()
}
