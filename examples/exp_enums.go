package main

import "fmt"

type ActiveStatus int32

const (
	Inactive  ActiveStatus = 0
	Activated ActiveStatus = 1
)

func (s ActiveStatus) String() string {
	switch s {
	case Inactive:
		return "未激活"
	case Activated:
		return "已激活"
	default:
		return "未知状态"
	}
}

func main() {
	fmt.Printf("Inactive: %d\n", Inactive) // Inactive: 0
	fmt.Printf("Activated: %d\n", Activated) // Activated: 1
	fmt.Printf("Inactive: %v\n", Inactive) // Inactive: 未激活
	fmt.Printf("Activated: %v\n", Activated) // Activated: 已激活
}
