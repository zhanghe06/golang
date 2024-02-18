package main

import (
	"fmt"
	"regexp"
)

func testRegexp() {
	m := regexp.MustCompile(`ee`) // 默认大小写敏感
	//m := regexp.MustCompile(`(?i)ee`) // 忽略大小写敏感

	fmt.Println(m.FindAllStringIndex("GeeksgeeksGeeks, geeks", -1)) // [[1 3] [6 8] [11 13] [18 20]]
	fmt.Println(m.FindAllStringIndex("GeeksgeeksGeeks, geeks", 2)) // [[1 3] [6 8]]
	fmt.Println(m.FindAllStringIndex("Hello! geeksForGEEKs", -1)) // 大小写敏感：[[8 10]]；忽略大小写：[[8 10] [16 18]]
	fmt.Println(m.FindAllStringIndex("I like Go language", -1)) // []
	fmt.Println(m.FindAllStringIndex("Hello, Welcome", -1)) // []
}

type RegexpFilter struct {
	Pattern string `binding:"required" form:"pattern"`
	CaseSensitive string `binding:"omitempty" form:"caseSensitive,default=n"`
	FromOffset int `binding:"omitempty" form:"fromOffset,default=0"`
	ToOffset int `binding:"omitempty" form:"toOffset,default=-1"`
	NumResults int `binding:"omitempty" form:"numResults,default=1"`
}


func search(content string, filter RegexpFilter)  {
	var regStr string
	if filter.CaseSensitive == "y" {
		regStr = fmt.Sprintf(`%s`, filter.Pattern)
	} else {
		regStr = fmt.Sprintf(`(?i)%s`, filter.Pattern)
	}
	m := regexp.MustCompile(regStr)
	res := m.FindAllStringIndex(content, filter.NumResults)
	fmt.Printf("result: %d\n", len(res))

	var indexSlice []int
	for _, v := range res {
		indexSlice = append(indexSlice, v[0])
	}

	if filter.FromOffset == 0 && filter.ToOffset == -1 {
		fmt.Printf("indexs: %v", indexSlice)
	}
	if filter.FromOffset <= filter.ToOffset {
		var indexSliceNew []int
		for _, v := range indexSlice {
			if !(filter.FromOffset <= v && v <= filter.ToOffset) {
				continue
			}
			indexSliceNew = append(indexSliceNew, v)
		}
		indexSlice = indexSliceNew
		fmt.Printf("indexs: %v", indexSlice)
	} else if filter.FromOffset >= filter.ToOffset {
		var indexSliceNew []int
		for _, v := range indexSlice {
			if !(filter.ToOffset <= v && v <= filter.FromOffset) {
				continue
			}
			indexSliceNew = append(indexSliceNew, v)
		}
		indexSlice = indexSliceNew
		fmt.Printf("indexs: %v", indexSlice)
	}
}


func main()  {
	testRegexp()

	//全部
	//f := RegexpFilter{
	//	Pattern: "你",
	//	CaseSensitive: "n",
	//	FromOffset: 0,
	//	ToOffset: -1,
	//	NumResults: -1,
	//}

	//正序
	//f := RegexpFilter{
	//	Pattern: "你",
	//	CaseSensitive: "n",
	//	FromOffset: 7,
	//	ToOffset: 11,
	//	NumResults: -1,
	//}

	//反序
	f := RegexpFilter{
		Pattern: "你",
		CaseSensitive: "n",
		FromOffset: 11,
		ToOffset: 7,
		NumResults: -1,
	}
	search("你好，你今天吃了么？你要是没吃，给我带一份炒面", f)
}
