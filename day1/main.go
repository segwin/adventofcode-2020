package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("day1/input")
	if err != nil {
		fmt.Printf("Read error\n")
		panic(err)
	}

	strValues := strings.Split(string(f), "\n")
	values := make([]int, 0, len(strValues))

	for i, strValue := range strValues {
		val, err := strconv.Atoi(strValue)
		if err != nil {
			fmt.Printf("Parse error on line %d\n", i)
			continue
		}

		values = append(values, val)
	}

	for i := range values {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Printf("FOUND: %d (%d) + %d (%d) + %d (%d) = 2020\n", values[i], i, values[j], j, values[k], k)
					fmt.Printf("PRODUCT: %d*%d+%d = %d\n", values[i], values[j], values[k], values[i]*values[j]*values[k])
				}
			}
		}
	}
}
