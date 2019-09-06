package main

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/Knetic/govaluate.v3"
)

func main() {

	//Initialise
	context := make(map[string]interface{})

	context["score"] = 56

	score := context["score"]

	cond1 := "$score>10"

	fmt.Println("cond1 :", cond1)

	cond1 = strings.Replace(cond1, "$score", strconv.Itoa(score.(int)), -1)

	fmt.Println("Rev cond1 :", cond1)

	expression, _ := govaluate.NewEvaluableExpression(cond1)
	expressionResult, _ := expression.Evaluate(nil)
	conditionStatus := expressionResult.(bool)

	if conditionStatus {
		fmt.Println("+Ve Activity")
	}
	if !conditionStatus {
		fmt.Println("-Ve Activity")
	}
}
