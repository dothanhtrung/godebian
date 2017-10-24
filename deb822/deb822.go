/*
 * Copyright (C) 2017 Do Thanh Trung <dothanhtrung.16@gmail.com>
 */

package deb822

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
 * Parse the whole file and return array of key-value
 */
func Parse(pathfile string) []map[string]string {
	var blocks []map[string]string

	file, err := os.Open(pathfile)
	if err != nil {
		fmt.Println(err)
		return blocks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	block := make(map[string]string)
	key, value := "", ""
	for scanner.Scan() {
		line := scanner.Text()
		key_match, _ := regexp.MatchString("^\\S*:.*", line)
		if key_match {
			kv := strings.Split(line, ":")
			key, value = kv[0], kv[1]
		}

		value_match, _ := regexp.MatchString("^\\s+\\S*", line)
		if value_match {
			value += line
		}
		block[key] = value

		break_match, _ := regexp.MatchString("^\\s*$", line)
		if break_match {
			blocks = append(blocks, block)
			block = make(map[string]string)
			key = ""
			value = ""
		}
	}

	return blocks
}
