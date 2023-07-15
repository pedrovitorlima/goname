package FileManager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitIntoOddEvenListsAndOrdered(t *testing.T) {
	organiser := FileOrganiser{}

	filenames := []string{"odd-1.jpeg", "odd-2.jpeg", "even-1.jpeg", "odd-3.jpeg",
		"even-3.jpeg", "even-2.jpeg", "odd-5.jpeg", "odd-4.jpeg", "even-4.jpeg",
		"even-101.jpeg", "even-11.jpeg", "even-102.jpeg", "even-120.jpeg", "even-106.jpeg"}

	odd, even := organiser.SplitIntoOddEvenLists(filenames)

	expectedOddList := []string{"odd-1.jpeg", "odd-2.jpeg", "odd-3.jpeg", "odd-4.jpeg", "odd-5.jpeg"}
	assert.Equal(t, expectedOddList, odd)

	expectedEvenList := []string{"even-1.jpeg", "even-2.jpeg", "even-3.jpeg", "even-4.jpeg", "even-11.jpeg",
		"even-101.jpeg", "even-102.jpeg", "even-106.jpeg", "even-120.jpeg"}
	assert.Equal(t, expectedEvenList, even)
}

func TestCreateRenameCommandWhenMoreEvenThanOdds(t *testing.T) {
	oddList := []string{"odd-1.jpeg", "odd-2.jpeg", "odd-3.jpeg", "odd-4.jpeg"}
	evenList := []string{"even-1.jpeg", "even-2.jpeg", "even-3.jpeg", "even-4.jpeg", "even-5.jpeg"}

	organiser := FileOrganiser{}

	commandMap := organiser.CreateRenameCommandMap(oddList, evenList)

	expectedCommandMap := map[string]string{
		"odd-1.jpeg":  "page-0001.jpeg",
		"even-1.jpeg": "page-0002.jpeg",
		"odd-2.jpeg":  "page-0003.jpeg",
		"even-2.jpeg": "page-0004.jpeg",
		"odd-3.jpeg":  "page-0005.jpeg",
		"even-3.jpeg": "page-0006.jpeg",
		"odd-4.jpeg":  "page-0007.jpeg",
		"even-4.jpeg": "page-0008.jpeg",
		"even-5.jpeg": "page-0009.jpeg",
	}

	assertMapsAreEquivalent(t, expectedCommandMap, commandMap)
}

func TestCreateRenameCommandWhenSameLengthOfOddsAndEven(t *testing.T) {
	oddList := []string{"odd-1.jpeg", "odd-2.jpeg", "odd-3.jpeg", "odd-4.jpeg"}
	evenList := []string{"even-1.jpeg", "even-2.jpeg", "even-3.jpeg", "even-4.jpeg"}

	organiser := FileOrganiser{}

	commandMap := organiser.CreateRenameCommandMap(oddList, evenList)

	expectedCommandMap := map[string]string{
		"odd-1.jpeg":  "page-0001.jpeg",
		"even-1.jpeg": "page-0002.jpeg",
		"odd-2.jpeg":  "page-0003.jpeg",
		"even-2.jpeg": "page-0004.jpeg",
		"odd-3.jpeg":  "page-0005.jpeg",
		"even-3.jpeg": "page-0006.jpeg",
		"odd-4.jpeg":  "page-0007.jpeg",
		"even-4.jpeg": "page-0008.jpeg",
	}

	assertMapsAreEquivalent(t, expectedCommandMap, commandMap)
}

func TestCreateRenameCommandWhenMoreOddsThanEvens(t *testing.T) {
	oddList := []string{"odd-1.jpeg", "odd-2.jpeg", "odd-3.jpeg", "odd-4.jpeg", "odd-5.jpeg"}
	evenList := []string{"even-1.jpeg", "even-2.jpeg", "even-3.jpeg", "even-4.jpeg"}

	organiser := FileOrganiser{}

	commandMap := organiser.CreateRenameCommandMap(oddList, evenList)

	expectedCommandMap := map[string]string{
		"odd-1.jpeg":  "page-0001.jpeg",
		"even-1.jpeg": "page-0002.jpeg",
		"odd-2.jpeg":  "page-0003.jpeg",
		"even-2.jpeg": "page-0004.jpeg",
		"odd-3.jpeg":  "page-0005.jpeg",
		"even-3.jpeg": "page-0006.jpeg",
		"odd-4.jpeg":  "page-0007.jpeg",
		"even-4.jpeg": "page-0008.jpeg",
		"odd-5.jpeg":  "page-0009.jpeg",
	}

	assertMapsAreEquivalent(t, expectedCommandMap, commandMap)
}

func assertMapsAreEquivalent(t *testing.T, expectedCommandMap map[string]string, commandMap map[string]string) {
	for expectedPage, expectedValueToBeRenamed := range expectedCommandMap {
		actualValueToBeRenamed, found := commandMap[expectedPage]
		assert.True(t, found, "Key not found in actual map: %s", expectedPage)
		assert.Equal(t, expectedValueToBeRenamed, actualValueToBeRenamed)
	}

	if t.Failed() {
		for page, toBeRenamed := range commandMap {
			t.Logf("%s -> %s", page, toBeRenamed)
		}
	}
}
