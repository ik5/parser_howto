package file

import "testing"

func TestReadFileValid(t *testing.T) {
	lines, err := ReadFile("file.txt")
	if err != nil {
		t.Errorf("Error returned: %s", err)
	}

	var output []string
	output = append(output, "Hello world")
	output = append(output, "")
	output = append(output, "End of Input")
	for i, line := range lines {
		if output[i] != line {
			t.Errorf("Lines '%s' and '%s' are not the same", output[i], line)
		}
	}
}
