/**
 * velocity4go: Velocity template engine for Go
 * https://sangupta.com/projects/velocity4go
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 */

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"sangupta.com/velocity4go/parser"
)

func main() {
	template := "Hello, ${name}. #if ($user.isOldEnough($age)) You are ${age} years old. #end"
	parseTemplate(template)
}

func parseTemplate(template string) {
	chars := []rune(template)
	parser := parser.Parser{
		Chars:        chars,
		ResourceName: "template.vm",
	}

	start := time.Now()
	parsedTemplate, err := parser.Parse()
	duration := time.Since(start)

	fmt.Println("time taken: " + duration.String())

	if err != nil {
		fmt.Println("Failed")
		return
	}

	j, ok := json.MarshalIndent(parsedTemplate, "", "  ")
	if ok != nil {
		fmt.Println("cannot convert to JSON")
		return
	}

	fmt.Println(string(j))
}
