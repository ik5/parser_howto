package file

import (
	"fmt"
	"os"
)

// MaxFileBufferSize is the maximum file buffer size
const MaxFileBufferSize = 50 * 1024 * 1024 * 1024

// ReadFile takes an existed file and convert it into a slice of lines
// If an error is found, then the lines are empty, and an error is returned
func ReadFile(path string) (lines []string, err error) {
	fl, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer fl.Close()

	var fileBuf []byte
	stat, _ := fl.Stat()
	fsize := stat.Size()
	// Do not read more then MaxFileBufferSize
	if fsize > MaxFileBufferSize {
		fileBuf = make([]byte, MaxFileBufferSize)
	} else {
		fileBuf = make([]byte, fsize)
	}

	count, err := fl.Read(fileBuf)
	if err != nil {
		return lines, err
	}
	if count == 0 {
		return lines, fmt.Errorf("File size is %d, but unable to read content", fsize)
	}

	line := ""

	for i, ch := range fileBuf {
		if ch == '\n' || i == count {
			lines = append(lines, line)
			line = ""
			continue
		}

		line = line + string(ch)
	}

	return
}
