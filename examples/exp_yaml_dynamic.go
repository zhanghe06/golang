package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type Faq struct {
	FAQ map[string]FaqEvent `yaml:"FAQ"`
}

type FaqEvent struct {
	Equals   []string `yaml:"equals"`
	Contains []string `yaml:"contains"`
}

// Struct 处理动态 Yaml
func testYamlDynamicStruct(y string) {
	var faqStruct Faq
	err := yaml.Unmarshal([]byte(y), &faqStruct)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 动态结构处理
	for eventK, eventV := range faqStruct.FAQ {
		fmt.Println(eventK)
		// Equals
		fmt.Println("\t", "Equals")
		for equalsIndex, equalsValue := range eventV.Equals {
			fmt.Println("\t\t", equalsIndex, ":", equalsValue)
		}
		// Contains
		fmt.Println("\t", "Contains")
		for containsIndex, containsValue := range eventV.Contains {
			fmt.Println("\t\t", containsIndex, ":", containsValue)
		}
	}
}

// Map 处理动态 Yaml
func testYamlDynamicMap(y string) {
	// 获取faq
	faq := make(map[string]map[string]FaqEvent)

	err := yaml.Unmarshal([]byte(y), &faq)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 动态结构处理
	for faqK, faqV := range faq {
		if faqK == "FAQ" {
			for eventK, eventV := range faqV {
				fmt.Println(eventK)
				// Equals
				fmt.Println("\t", "Equals")
				for equalsIndex, equalsValue := range eventV.Equals {
					fmt.Println("\t\t", equalsIndex, ":", equalsValue)
				}
				// Contains
				fmt.Println("\t", "Contains")
				for containsIndex, containsValue := range eventV.Contains {
					fmt.Println("\t\t", containsIndex, ":", containsValue)
				}
			}
		}
	}
}

func keywordsContains(sentence, words, separator string) bool {
	keywords := strings.Split(words, separator)
	result := true
	for _, v := range keywords {
		// fmt.Println(sentence, " - ", v)
		result = result && strings.Contains(sentence, strings.Trim(v, " "))
	}
	return result
}

func testHit(y, txt string) {
	// 获取faq
	faq := make(map[string]map[string]FaqEvent)

	err := yaml.Unmarshal([]byte(y), &faq)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Dynamic FAQ
	for eventK, eventV := range faq["FAQ"] {
		// Equals
		for _, equalsValue := range eventV.Equals {
			// fmt.Println(equalsValue, " - ", txt)
			if strings.ToLower(equalsValue) == strings.ToLower(txt) {
				fmt.Println(eventK)
				return
			}
		}
		// Contains
		for _, containsValue := range eventV.Contains {
			// fmt.Println(containsValue, " - ", txt)
			if keywordsContains(strings.ToLower(txt), strings.ToLower(containsValue), ",") {
				fmt.Println(eventK)
				return
			}
		}
	}
	fmt.Println("OtherEvent")
	return
}

// 动态结构中, 部分结构的KEY是固定的:FAQ、equals、contains; 事件名称是动态的: EventQ01、EventQ02
func main() {
	y := `
FAQ:
    EventQ01:
        equals:
            - sentence_01_a
            - sentence_01_b
        contains:
            - keyword_01_a, keyword_01_b
            - keyword_01_c
            - keyword_01_d1 keyword_01_d2
    EventQ02:
        equals:
            - sentence_02_a
            - sentence_02_b
        contains:
            - keyword_02_a, keyword_02_b
            - keyword_02_c
            - keyword_02_d1 keyword_02_d2
`
	// 假设不包含FAQ节点
	// y = ``

	// Struct 处理动态 Yaml
	// testYamlDynamicStruct(y)

	// Map 处理动态 Yaml
	// testYamlDynamicMap(y)

	// 测试命中
	testHit(y, "sentence_01_a")   // equals
	testHit(y, "sentence_01_A")   // equals
	testHit(y, "keyword_02_a")    // contains 缺少必须关键词数量
	testHit(y, "keyword_02_A")    // contains 缺少必须关键词数量
	testHit(y, "keyword_02_c")    // contains
	testHit(y, "keyword_02_C")    // contains
	testHit(y, "keyword_02_d2")   // contains 缺少必须关键词文本
	testHit(y, "keyword_02_D2")   // contains 缺少必须关键词文本
	testHit(y, "keyword_02_c ha") // contains
	testHit(y, "keyword_02_C ha") // contains
}
