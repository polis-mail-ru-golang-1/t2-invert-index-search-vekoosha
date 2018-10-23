package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-vekoosha/invertIndex"
)

func main() {

	namesFiles := os.Args[1:]

	elements := make(map[string]map[string]int)

	for i := 0; i < len(namesFiles); i++ {
		invertIndex.MakeIndexMap(elements, namesFiles[i], contentFile(namesFiles[i]))
	}

	var slicePhrase []string
	slicePhrase = readPhrase(slicePhrase)

	resultMap := make(map[string]int)
	resultMap = invertIndex.SearchPhrase(elements, slicePhrase, namesFiles)
	resultMap = invertIndex.SortResult(resultMap)
	invertIndex.PrintResult(resultMap)
}

func readPhrase(slicePhrase []string) []string {
	fmt.Printf("Введите запрос: ")
	reader := bufio.NewReader(os.Stdin)
	request, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	helpStr := strings.TrimSpace(request)
	slicePhrase = strings.Split(helpStr, " ")
	return slicePhrase
}

func contentFile(nameFile string) []string {
	file, err := os.Open(nameFile)
	if err != nil {
		panic("Невозможно открыть файл.")
	}
	defer file.Close()

	var contentFile []string
	helpStr := nameFile
	nameByte, err := ioutil.ReadFile(helpStr)
	if err != nil {
		panic(err)
	}
	nameStr := string(nameByte)
	contentFile = strings.Split(nameStr, " ")
	return contentFile
}
