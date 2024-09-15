package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Dictionary struct {
	dict map[int]string
}

func (d *Dictionary) AddWords(words []string) {
	for _, word := range words {
		d.dict[utf8.RuneCountInString(word)] += strings.ToLower(strings.TrimSpace(word)) + " "
	}
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		dict: make(map[int]string),
	}
}

func Start(wordList []string, dict *Dictionary) map[string][]string {
	result := make(map[string][]string)
	for _, word := range wordList {
		wordAnagrams := anagrams(word, dict)
		if len(wordAnagrams) > 1 {
			sort.Strings(wordAnagrams)
			result[strings.ToLower(strings.TrimSpace(word))] = wordAnagrams
		}
	}
	return result
}

func anagrams(word string, dict *Dictionary) []string {
	validWord := strings.ToLower(strings.TrimSpace(word))
	validWorldLen := utf8.RuneCountInString(validWord)

	var regExpr strings.Builder
	regExpr.WriteString("[")

	for _, symbol := range validWord {
		regExpr.WriteString(string(symbol) + ",")
	}

	regExpr.WriteString("]{" + strconv.Itoa(validWorldLen) + "|")

	re := regexp.MustCompile(regExpr.String())

	return re.FindAllString(dict.dict[validWorldLen], -1)
}

func TaskFour() {
	myDict := NewDictionary()
	myDict.AddWords([]string{"АМКАР", "КАРМА", "КРАМА", "МАКАР", "МАКРА", "МАРКА", "РАМКА",
		"ПЯТАК", "ПЯТКА", "ТЯПКА", "КОСАЧ", "САЧОК", "ЧАСОК", "АВТОР", "ВАРТО", "ВТОРА", "ОТВАР",
		"РВОТА", "ТАВРО", "ТОВАР", "КАЧУР", "КРАУЧ", "КРУЧА", "КУРЧА", "РУЧКА", "ЧУРКА", "АБНЯ",
		"БАНЯ", "БАЯН", "КОРТ", "КРОТ", "ТРОК", "КОТ", "КТО", "ОТК", "ТОК",
	})

	fmt.Println(Start([]string{"кот"}, myDict))
}
