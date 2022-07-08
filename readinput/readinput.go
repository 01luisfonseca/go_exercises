package readinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text, err
}
