package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
)

type MyTuple struct {
	key    string
	result float64
}

func close_program() {
	fmt.Println("Press ENTER to exit...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	os.Exit(0)
}

func make_operation(operator string, args ...float64) float64 {
	var result float64 = math.NaN()
	if operator == "add" && len(args) == 2 {
		result = args[0] + args[1]
		// fmt.Printf("[DEBUG] %v + %v = %v\n", args[0], args[1], result)
		return result
	}
	if operator == "sub" && len(args) == 2 {
		result = args[0] - args[1]
		// fmt.Printf("[DEBUG] %v - %v = %v\n", args[0], args[1], result)
		return result
	}
	if operator == "mul" && len(args) == 2 {
		result = args[0] * args[1]
		// fmt.Printf("[DEBUG] %v * %v = %v\n", args[0], args[1], result)
		return result
	}
	if operator == "sqrt" && len(args) == 1 {
		result = math.Sqrt(args[0])
		// fmt.Printf("[DEBUG] sqrt(%v) = %v\n", args[0], result)
		return result
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("[INFO] If left empty, program will search for \"input.json\" file.")
	fmt.Printf("Enter a valid path to the input file: ")
	scanner.Scan()
	jsonFileName := scanner.Text()
	if jsonFileName == "" {
		jsonFileName = "input.json"
		fmt.Printf("[INFO] Left empty, searching for \"%v\" file.\n", jsonFileName)
	}
	outputFileName := "output.txt"

	jsonFile, err := os.Open(jsonFileName)

	if err != nil {
		msg := fmt.Sprintf("[ERROR] Could not find \"%v\" file. Program will exit.\n",
			jsonFileName)
		fmt.Printf(msg)
		close_program()
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data map[string]map[string]interface{}
	json.Unmarshal([]byte(byteValue), &data)

	sortedTuple := make([]MyTuple, 0, 100)

	validOperations := 0
	for key, element := range data {
		operator, _ := element["operator"].(string)
		arg1, ok1 := element["value1"].(float64)
		arg2, ok2 := element["value2"].(float64)

		// searching if operator name is allowed
		var ok_operator bool
		for _, e := range []string{"add", "sub", "mul", "sqrt"} {
			if e == operator {
				ok_operator = true
				break
			}
		}

		// if operator is allowed than operation can be made
		if ok_operator {
			var result float64 = math.NaN()
			if ok1 && ok2 {
				result = make_operation(operator, arg1, arg2)
			}
			if ok1 && !ok2 && operator == "sqrt" {
				result = make_operation(operator, arg1)
			}

			if !math.IsNaN(result) {
				validOperations++
			}
			// fmt.Printf("%v: %v\n", key, result)
			sortedTuple = append(sortedTuple, MyTuple{key, result})
		}

	}

	sort.SliceStable(sortedTuple, func(i, j int) bool {
		// custom comparator needed for placing NaN values first
		val1 := sortedTuple[i].result
		val2 := sortedTuple[j].result

		isnan1 := math.IsNaN(val1)
		isnan2 := math.IsNaN(val2)

		if isnan1 && isnan2 {
			return true
		}

		if isnan1 {
			return true
		}

		if isnan2 {
			return false
		}

		return val1 < val2
	})

	// fmt.Println("------------------------------")
	// for _, e := range sortedTuple {
	// 	fmt.Printf("[DEBUG] %v: %v\n", e.key, e.result)
	// }

	fmt.Printf("[INFO] Valid operations: %v (out of %v)\n", validOperations, len(data))

	outputFile, err := os.Create(outputFileName)

	if err != nil {
		msg := "[ERROR] Could not create output file. Program will exit."
		fmt.Println(msg)
		close_program()
	}

	defer outputFile.Close()

	for _, e := range sortedTuple {
		_, err = fmt.Fprintf(outputFile, "%v: %v\n", e.key, e.result)

		if err != nil {
			msg := "[ERROR] An error occured while saving results to output file. Program will exit"
			fmt.Println(msg)
			close_program()
		}
	}

	fmt.Printf("[INFO] Results saved successfully to file: %v.\n", outputFileName)
	close_program()
}
