/*
 * Copyright (C) 2017 Do Thanh Trung <dothanhtrung.16@gmail.com>
 */

package deb822

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

/*
 * Find section by key-value
 */
func find(pathfile, skey, svalue string, limit uint) ([]map[string]string, error) {
	var results []map[string]string

	file, err := os.Open(pathfile)
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	section := make(map[string]string)
	key, value := "", ""
	found, ignore := false, false
	var count uint
	for scanner.Scan() {
		line := scanner.Text()
		if !ignore {
			keyMatch, _ := regexp.MatchString("^\\S*:.*", line)
			if keyMatch {
				if (key == skey || skey == "") && (value == svalue || svalue == "") {
					found = true
				} else if key != "" && key == skey && value != svalue {
					ignore = true
				}

				kv := strings.Split(line, ":")
				key = strings.TrimSpace(kv[0])
				value = strings.TrimSpace(kv[1])

			} else {
				valueMatch, _ := regexp.MatchString("^\\s+\\S*", line)
				if valueMatch {
					if value != "" {
						value += "\n"
					}
					value += strings.TrimSpace(line)
				}
			}
			section[key] = value
		}

		// Break between sections
		breakMatch, _ := regexp.MatchString("^\\s*$", line)
		if breakMatch {
			if found {
				results = append(results, section)
				count += 1
				found = false
				if limit > 0 && count >= limit {
					break
				}
			}

			ignore = false
			section = make(map[string]string)
			key = ""
			value = ""
		}
	}

	if len(results) == 0 {
		return results, errors.New("not found")
	}

	return results, nil
}

/*
 * Find a number of sections contain wanted key-value
 */
func Find(pathfile, skey, svalue string, limit uint) ([]map[string]string, error) {
	return find(pathfile, skey, svalue, limit)
}

/*
 * Find the first section contains wanted key-value
 */
func FindOne(pathfile, skey, svalue string) (map[string]string, error) {
	result := make(map[string]string)
	results, err := find(pathfile, skey, svalue, 1)
	if err == nil {
		result = results[0]
	}

	return result, err
}

/*
 * Find all the sections contain wanted key-value
 */
func FindAll(pathfile, skey, svalue string) ([]map[string]string, error) {
	return find(pathfile, skey, svalue, 0)
}
