package _15

import (
	"advent_of_code23/utils"
	"bytes"
	orderedmap "github.com/wk8/go-ordered-map"
	"regexp"
	"strconv"
)

func Problem1(inputFileName string) int {
	instructions := bytes.Split(utils.GetFileContent(inputFileName), []byte(","))
	sum := 0
	for _, instruction := range instructions {
		sum += hashString(instruction)
	}

	return sum
}

func Problem2(inputFileName string) int {
	instructions := bytes.Split(utils.GetFileContent(inputFileName), []byte(","))
	hmap := make(map[int]*orderedmap.OrderedMap)

	for _, instruction := range instructions {
		if instruction[len(instruction)-1] == '-' {
			instruction = append(instruction, '0')
		}
		portions := regexp.MustCompile(`(\w+)([=-])(\d+)`).FindStringSubmatch(string(instruction))

		id := portions[1]
		hashId := hashString([]byte(id))
		operation := portions[2][0] // getting at [0] is a shortcut to getting the char

		if operation == '-' {
			if _, ok := hmap[hashId]; ok {
				(*hmap[hashId]).Delete(id)
			}

			continue
		}

		lens, err := strconv.Atoi(portions[3])

		if err != nil {
			panic(err)
		}

		if _, ok := hmap[hashId]; !ok {
			hmap[hashId] = orderedmap.New()
		}
		(*hmap[hashId]).Set(id, lens)
	}

	sum := 0
	for k, v := range hmap {
		i := 1
		for item := v.Oldest(); item != nil; item = item.Next() {
			sum += (1 + k) * i * item.Value.(int)
			i++
		}
	}

	return sum
}

func hashString(b []byte) int {
	hashVal := 0
	for _, c := range b {
		hashVal += int(c)
		hashVal *= 17
		hashVal %= 256
	}

	return hashVal
}
