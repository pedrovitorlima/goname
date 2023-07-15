package FileManager

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type FileOrganiser struct {
}

func (f FileOrganiser) SplitIntoOddEvenLists(filenames []string) ([]string, []string) {

	oddList := make([]string, 0)
	evenList := make([]string, 0)

	for _, filename := range filenames {
		if strings.HasPrefix(filename, "odd-") {
			oddList = append(oddList, filename)
		} else {
			evenList = append(evenList, filename)
		}
	}

	sort.Slice(oddList, func(index1, index2 int) bool {
		return sortByPageNumberFromFilename(index1, index2, oddList)
	})

	sort.Slice(evenList, func(index1, index2 int) bool {
		return sortByPageNumberFromFilename(index1, index2, evenList)
	})

	return oddList, evenList
}

func (f FileOrganiser) CreateRenameCommandMap(oddList []string, evenList []string) map[string]string {
	commands := make(map[string]string, 0)

	pageCounter := 1
	for index, oddPage := range oddList {
		commands[oddPage] = generateFileName(pageCounter)
		pageCounter++

		lastIteration := index+1 == len(oddList)
		if lastIteration && len(evenList) >= index {
			for _, evenPage := range evenList[index:] {
				commands[evenPage] = generateFileName(pageCounter)
				pageCounter++
			}
			break
		}

		if len(evenList) >= index {
			evenPage := evenList[index]
			commands[evenPage] = generateFileName(pageCounter)
			pageCounter++
		}
	}

	return commands
}

func generateFileName(pageNumber int) string {
	pageNumberWithDecimalCases := fmt.Sprintf("%04d", pageNumber)
	return "page-" + pageNumberWithDecimalCases + ".jpeg"
}

func sortByPageNumberFromFilename(index1 int, index2 int, oddList []string) bool {
	pageNumber1, err := extractPageNumber(oddList[index1])
	if err != nil {
		log.Fatal(err)
	}

	pageNumber2, err := extractPageNumber(oddList[index2])
	if err != nil {
		log.Fatal(err)
	}

	return pageNumber1 < pageNumber2
}

func extractPageNumber(filename string) (int, error) {
	from := strings.Index(filename, "-") + 1
	until := strings.Index(filename, ".")
	return strconv.Atoi(filename[from:until])
}
