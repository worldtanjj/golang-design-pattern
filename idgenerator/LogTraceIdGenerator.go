package idgenerator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//LogTraceIdGenerator 结构体
//结构体a
type LogTraceIdGenerator struct {
}

func (l LogTraceIdGenerator) Generate() (string, error) {

	var substrOfHostName, err = getLastfieldOfHostName()
	if err != nil {
		return substrOfHostName, err
	}

	var currentTimeMillis = time.Now().UnixNano() / 1e6
	var randomString = generateRandomAlphameric(8)

	var id = fmt.Sprintf("%s-%d-%s", substrOfHostName, currentTimeMillis, randomString)
	return id, nil
}

func getLastfieldOfHostName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname:", hostname)
		return "", err
	}

	var tokens = strings.Split(hostname, "\\.")
	substrOfHostName := tokens[len(tokens)-1]
	return substrOfHostName, nil
}

func generateRandomAlphameric(length int) string {
	var count = 0
	var randomChars = make([]byte, 8)
	var maxASCII = 'z'
	for count < length {
		var randomASCII = rune(rand.Intn(int(maxASCII)))
		var isDigit = randomASCII >= '0' && randomASCII <= '9'
		var isUppercase = randomASCII >= 'A' && randomASCII <= 'Z'
		var isLowercase = randomASCII >= 'a' && randomASCII <= 'z'
		if isDigit || isUppercase || isLowercase {
			randomChars[count] = byte(randomASCII)
			count++
		}
	}
	return string(randomChars[:])
}
