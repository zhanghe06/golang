package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 排名抽奖
func drawRank(userIndex, totalIndex int) (res bool)  {

	countZero := userIndex * (totalIndex)
	countOne := totalIndex - userIndex

	// 以防每次打乱顺序都一样
	rand.Seed(time.Now().UnixNano())

	// 奖品slice
	var prizeSlice []int

	// 追加空值
	indexZero := 0
	for {
		if indexZero >= countZero {
			break
		}
		prizeSlice = append(prizeSlice, 0)
		indexZero++
	}

	// 追加数值
	indexOne := 0
	for {
		if indexOne >= countOne {
			break
		}
		prizeSlice = append(prizeSlice, 1)
		indexOne++
	}

	rand.Shuffle(len(prizeSlice), func(i, j int) {
		prizeSlice[i], prizeSlice[j] = prizeSlice[j], prizeSlice[i]
	})
	res = prizeSlice[0] == 1
	//fmt.Println(prizeSlice)
	fmt.Println(res)
	return
}

// 积分抽奖
func drawScore(userScore, totalScore, prizeCount int) (res bool)  {

	countZero := totalScore - userScore
	countOne := userScore

	// 以防每次打乱顺序都一样
	rand.Seed(time.Now().UnixNano())

	// 奖品slice
	var prizeSlice []int

	// 追加空值
	indexZero := 0
	for {
		if indexZero >= countZero {
			break
		}
		prizeSlice = append(prizeSlice, 0)
		indexZero++
	}

	// 追加数值
	indexOne := 0
	for {
		if indexOne >= countOne {
			break
		}
		prizeSlice = append(prizeSlice, 1)
		indexOne++
	}

	rand.Shuffle(len(prizeSlice), func(i, j int) {
		prizeSlice[i], prizeSlice[j] = prizeSlice[j], prizeSlice[i]
	})
	res = prizeSlice[0] == 1
	fmt.Println(prizeSlice)
	fmt.Println(res)
	return
}

func main() {
	//drawRank(10, 100)
	//drawScore(500, 1000, 3)

	prizeSlice := []int{1,2,3,4,5}
	i := 0
	prizeSlice = append(prizeSlice[:i], prizeSlice[i+1:]...)
	fmt.Println(prizeSlice)
}


