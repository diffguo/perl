package main

import (
	"fmt"
	"sort"
	"time"
	"unicode/utf8"
)

type StSuffixArray struct {
	Data []rune
	Sa   []int
	Rank []int
}

/******************************内部使用******************************/

//直接使用Index []int这种方式比用指针慢不少
type sortRuneStruct struct {
	Index *[]int
	Data  *[]rune
}

func (s sortRuneStruct) Len() int {
	return len(*(s.Data))
}

func (s sortRuneStruct) Swap(i, j int) {
	(*s.Index)[i], (*s.Index)[j] = (*s.Index)[j], (*s.Index)[i]
}

func (s sortRuneStruct) Less(i, j int) bool {
	return (*s.Data)[(*s.Index)[i]] < (*s.Data)[(*s.Index)[j]]
}

func sortByFirstRune(data []rune) []int {
	data_len := len(data)
	sa := make([]int, data_len)
	for i := 0; i < data_len; i++ {
		sa[i] = i
	}

	//Stable 在比较简单文本情况下比 Sort 慢很多
	sort.Sort(sortRuneStruct{Index: &sa, Data: &data})
	return sa
}

type sortIntStruct struct {
	Index *[]int
	Data  *[]int
}

func (s sortIntStruct) Len() int {
	return len(*(s.Data))
}

func (s sortIntStruct) Swap(i, j int) {
	(*s.Index)[i], (*s.Index)[j] = (*s.Index)[j], (*s.Index)[i]
}

func (s sortIntStruct) Less(i, j int) bool {
	return (*s.Data)[(*s.Index)[i]] < (*s.Data)[(*s.Index)[j]]
}

// 针对第一个字母进行排序
func sortByFirstRune(data []rune) []int {
	data_len := len(data)
	sa := make([]int, data_len)
	for i := 0; i < data_len; i++ {
		sa[i] = i
	}

	//Stable 在比较简单文本情况下比 Sort 慢很多
	sort.Sort(sortRuneStruct{Index: &sa, Data: &data})
	return sa
}

/******************************内部使用******************************/

func NewSuffixArrayByDoublingAlgorithm(data []rune, max_word_len int) *StSuffixArray {

	data_len := len(data)
	sa := sortByFirstRune(data)
	rank := make([]int, data_len)

	fmt.Printf("sa: ")
	for i := 0; i < data_len; i++ {
		fmt.Printf("%d:%d:%d   ", i, sa[i], data[sa[i]])
	}
	fmt.Printf("\n")

	//初始化rank
	rank_num := 0
	rank[sa[0]] = 0
	for k := 1; k < data_len; k++ {
		if data[sa[k-1]] == data[sa[k]] {
			rank[sa[k]] = rank[sa[k-1]]
		} else {
			rank_num++
			rank[sa[k]] = rank_num
		}
	}

	rank_tmp := make([]int, data_len)
	for i := 1; i < max_word_len; i *= 2 {
		for j := 0; j < data_len; j++ {
			if j+i < data_len {
				rank_tmp[j] = rank[j]*10 + rank[j+i]
			} else {
				rank_tmp[j] = rank[j] * 10
			}
		}

		rank_num = 0
		for j := 1; j < data_len; j++ {

		}

		fmt.Printf("rank: ")
		for k := 0; k < data_len; k++ {
			fmt.Printf("%d:%d   ", rank[k], data[k])
		}
		fmt.Printf("\n")
	}

	return &StSuffixArray{Data: data, Sa: sa, Rank: rank}
}

func main() {

	//data := []byte("我们看数组 iarray1，只声明，并未赋值，Go语言帮我们自动赋值为0。再看 iarray2 和 iarray3 ，我们可以看到，Go语言的声明，可以表明类型，也可以不表明类型，var iarray3 = [5]int32{1, 2, 3, 4, 5} 也是完全没问题的。")
	data := []byte("我们看数组 iarray1")

	data_rune := []rune{}
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		data_rune = append(data_rune, r)
		data = data[size:]
	}

	t0 := time.Now()
	for i := 0; i < 1; i++ {
		fmt.Println(data_rune)
		NewSuffixArrayByDoublingAlgorithm(data_rune, 6)
	}
	t1 := time.Now()
	fmt.Println(t1.Sub(t0).String())
}
