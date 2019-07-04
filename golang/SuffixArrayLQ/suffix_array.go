package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"unicode/utf8"
)

type StSuffixArray struct {
	Data []rune
	Sa   []int
}

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

	sort.Stable(sortRuneStruct{Index: &sa, Data: &data})
	return sa
}

// 生成SA(后缀数组): SA[i]存放排名第i大的子串首字符下标
func (self *StSuffixArray) buildSuffixByDoublingAlgorithm() []int {
	data := self.Data              //修改
	max_word_len := len(self.Data) //修改

	data_len := len(data)
	sa := sortByFirstRune(data)
	rank_x := make([]int, data_len)
	rank_y := make([]int, data_len)
	count := make([]int, data_len)
	wv := make([]int, data_len)

	rank_x[sa[0]] = 0
	rank_x_len := 1
	for i := 1; i < data_len; i++ {
		if data[sa[i-1]] == data[sa[i]] {
			rank_x[sa[i]] = rank_x_len - 1
		} else {
			rank_x[sa[i]] = rank_x_len
			rank_x_len++
		}
	}

	base_sort_range := rank_x_len
	for i := 1; i < max_word_len; i *= 2 {
		rank_y_len := 0
		for j := data_len - i; j < data_len; j++ {
			rank_y[rank_y_len] = j
			rank_y_len++
		}
		for j := 0; j < data_len; j++ {
			if sa[j] >= i {
				rank_y[rank_y_len] = sa[j] - i
				rank_y_len++
			}
		}
		for j := 0; j < data_len; j++ {
			wv[j] = rank_x[rank_y[j]]
		}

		for j := 0; j < base_sort_range; j++ {
			count[j] = 0
		}
		for j := 0; j < data_len; j++ {
			count[wv[j]]++
		}
		for j := 1; j < base_sort_range; j++ {
			count[j] += count[j-1]
		}
		for j := data_len - 1; j >= 0; j-- {
			count[wv[j]]--
			sa[count[wv[j]]] = rank_y[j]
		}
		rank_x, rank_y = rank_y, rank_x
		rank_x[sa[0]] = 0
		rank_x_len = 1
		for j := 1; j < data_len; j++ {
			if rank_y[sa[j-1]] == rank_y[sa[j]] && (sa[j-1]+i) < data_len &&
				(sa[j]+i) < data_len && rank_y[sa[j-1]+i] == rank_y[sa[j]+i] {
				rank_x[sa[j]] = rank_x_len - 1
			} else {
				rank_x[sa[j]] = rank_x_len
				rank_x_len++
			}
		}
		base_sort_range = rank_x_len
	}

	self.Sa = sa
	return sa
}

func (self *StSuffixArray) printSuffixArray() {
	runes := []rune{}
	for i := 0; i < len(self.Data); i++ {
		if 10+self.Sa[i] > len(self.Data) {
			runes = self.Data[self.Sa[i]:]
		} else {
			runes = self.Data[self.Sa[i] : self.Sa[i]+10]
		}

		fmt.Printf("Pos: %3d, Str: %s\n", self.Sa[i], string(runes))
	}
}

func (self *StSuffixArray) buildSuffix(data string) {
	self.Data = []rune{}
	for len(data) > 0 {
		r, size := utf8.DecodeRune([]byte(data))
		self.Data = append(self.Data, r)
		data = data[size:]
	}

	self.buildSuffixByDoublingAlgorithm()
}

func main() {
	fi, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer fi.Close()
	reader, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	data := string(reader)

	suffixArray := StSuffixArray{}
	suffixArray.buildSuffix(data)
	suffixArray.printSuffixArray()
}
