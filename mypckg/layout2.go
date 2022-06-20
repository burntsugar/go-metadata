package mypckg

import (
	"bufio"
	"log"
	"strings"
)

const logstring string = `1,2,3`

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func main() {

	scanner := bufio.NewScanner(strings.NewReader(logstring))
	scanner.Split(crunchSplitFunc)

	for scanner.Scan() {
		log.Print(scanner.Text())
	}
}
