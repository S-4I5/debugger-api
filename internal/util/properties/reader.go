package properties

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadProperties(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := make(map[string]string)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		splitted := strings.Split(line, "=")
		if len(splitted) != 2 {
			return nil, fmt.Errorf("incorrect file format")
		}

		cut, _ := strings.CutSuffix(splitted[1], "\r\n")

		properties[splitted[0]] = cut

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return properties, nil
}
