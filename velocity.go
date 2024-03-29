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

	"sangupta.com/velocity/parser"
)

func main() {
	// what callee will pass
	template := "Hello ${name}."
	variables := make(map[string]interface{})

	variables["name"] = "velocity4go"
	variables["numbers"] = []int{10, 20, 30, 40, 50}

	parseAndRender(template, variables, true, false)
}

func parseAndRender(template string, vars map[string]interface{}, showTemplate bool, render bool) {
	// our local code
	chars := []rune(template)
	parser := parser.Parser{
		Chars:        chars,
		ResourceName: "template.vm",
	}

	start := time.Now()
	parsedTemplate, err := parser.Parse()
	duration := time.Since(start)

	fmt.Println("time taken to parse: " + duration.String())
	if err != nil {
		fmt.Println("Failed")
		return
	}

	if showTemplate {
		jsonStr, ok := json.MarshalIndent(parsedTemplate, "", "  ")
		if ok != nil {
			fmt.Println("cannot convert to JSON")

			return
		}
		fmt.Println(string(jsonStr))
	}

	if !render {
		return
	}

	startRender := time.Now()
	rendered := parsedTemplate.Evaluate(vars)
	durationRender := time.Since(startRender)

	fmt.Println("time taken to render: " + durationRender.String())
	fmt.Println()
	fmt.Println("rendered: " + rendered)
	fmt.Println()
}
