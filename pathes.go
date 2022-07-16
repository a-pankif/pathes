package pathes

import (
	"errors"
	"os"
)

// SplitByGroups returns separated (by os.PathSeparator) string with given length of each group.
//
// Split string "abcdefg" to "\a\b\c\d\e\f" for groupsLengths [1,1,1,1,1,1]
//
// Split string "abcdefg" to "\a\bc\de\f" for groupsLengths [1,2,2,1]
//
// Split string "abcdefg" to "\a\bc" for groupsLengths [1,2]
func SplitByGroups(str string, groupsLengths ...int) (string, error) {
	split := ""
	position := 0
	separator := string(os.PathSeparator)
	strLength := len(str)

	for _, length := range groupsLengths {
		if length+position > strLength {
			return "", errors.New("groups lengths sum greater then str length")
		}

		split += separator + str[position:length+position]
		position += length
	}

	return split, nil
}
