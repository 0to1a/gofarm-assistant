package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestCommand(t *testing.T) {
	for i, test := range []struct {
		Args   []string
		Output string
	}{
		{
			Args:   []string{"./gofarm-assistant"},
			Output: helpText,
		},
		{
			Args:   []string{"./gofarm-assistant", "create"},
			Output: helpText,
		},
		{
			Args:   []string{"./gofarm-assistant", "help"},
			Output: helpText,
		},
		{
			Args:   []string{"./gofarm-assistant", "version"},
			Output: fmt.Sprintf(versionNoProjectText, version),
		},
	} {
		t.Run("Test case #"+strconv.Itoa(i+1), func(t *testing.T) {
			os.Args = test.Args
			output = bytes.NewBuffer(nil)
			main()

			if actual := output.(*bytes.Buffer).String(); actual != test.Output {
				fmt.Println(actual, test.Output)
				t.Errorf("expected %s, but got %s", test.Output, actual)
			}
		})
	}
}
