package invertIndex

import (
	"fmt"
	"sync"
)

func MakeIndexMap(Map map[string]map[string]int, name string, content []string, wg *sync.WaitGroup) {
	sliceWordFile := content
	for i := 0; i < len(sliceWordFile); i++ {
		word := sliceWordFile[i]
		if _, ok := Map[word]; !ok {
			newFile := make(map[string]int)
			newFile[name] = 1
			Map[word] = newFile
		} else {
			if _, ok := Map[word][name]; !ok {
				Map[word][name] = 1
			} else {
				Map[word][name]++
			}
		}
	}
	wg.Done()
}

func SearchPhrase(Map map[string]map[string]int, slicePhrase []string, sliceNameFile []string) map[string]int {
	phraseMap := make(map[string]map[string]int)
	for wordMap := range Map {
		for _, wordPhrase := range slicePhrase {
			if wordMap == wordPhrase {
				phraseMap[wordMap] = Map[wordMap]
			}
		}
	}

	resultMap := make(map[string]int)
	for i := 0; i < len(sliceNameFile); i++ {
		for _, fileMap := range phraseMap {
			for nameFile, count := range fileMap {
				if sliceNameFile[i] == nameFile {
					resultMap[sliceNameFile[i]] += count
				}
			}
		}
	}
	if len(resultMap) == 0 {
		fmt.Println("Ничего не найдено.")
	}
	return resultMap
}

type result struct {
	name  string
	count int
}

func SortResult(Map map[string]int) {
	var sorting []result

	for name, count := range Map {
		var pushElement result
		pushElement.name = name
		pushElement.count = count
		sorting = append(sorting, pushElement)
	}

	for j := len(sorting) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if sorting[i+1].count > sorting[i].count {
				tempElement := sorting[i]
				sorting[i] = sorting[i+1]
				sorting[i+1] = tempElement
			}
		}
	}

	for i := 0; i < len(sorting); i++ {
		fmt.Println("- ", sorting[i].name, " совпадений - ", sorting[i].count)

	}
}
