package pathes

import (
	"os"
	"testing"
)

func TestSplitByGroups(t *testing.T) {
	s := string(os.PathSeparator)

	type args struct {
		str           string
		groupsLengths []int
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"pass1", args{"abcdefg", []int{1, 2, 3}}, s + "a" + s + "bc" + s + "def", false},
		{"pass2", args{"abcdefg", []int{1, 1, 1, 1, 1}}, s + "a" + s + "b" + s + "c" + s + "d" + s + "e", false},
		{"pass3", args{"abcdefg", []int{1}}, s + "a", false},
		{"fail4", args{"abcdefg", []int{1, 32}}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitByGroups(tt.args.str, tt.args.groupsLengths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSplitToPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringSplitToPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
