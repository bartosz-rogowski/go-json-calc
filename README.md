# go-json-calc

Simple calculator written in Go reading data from input file and saving results to output file 

## How to run the program

Run *swi.exe* (on Windows) or run program:
```
go run swi.go
```

## Project description

The main task of the program is to read valid json file containing list of operators and data, for example snippet:
```json
"obj1": {
    "operator": "add",
    "value1": 2,
    "value2": 3
}
```
will result in making operation of addition: `2+3` which of course equals `5`. Results are saved to file *output.txt* in ascending order.

## Assumptions
* the json file is valid
* available operation are:
  - addition `"add"` - takes two real arguments
  - substraction `"sub"` - takes two real arguments
  - multiplication `"mul"` - takes two real arguments
  - square root `"sqrt"` - takes one real non-negative argument
* additional arguments will be omitted
* if arguments are not valid (for example: strings, characters etc.), the returned result is `NaN`
* if given operator does not belong to above list, such one is considered invalid and the returned result is also `NaN`
* results are sorted in ascending order in a way that `NaN` values are listed first
* user can enter valid path for input file, however if left blank, the program will search for default name file (called *input.json*)
* if an input file does not exist (or is invalid) the program will prompt proper information and exit.

## Example

*input.json*:
```json
{
    "obj1": {
        "operator": "add",
        "value1": 2,
        "value2": 3
    },
    "obj2": {
        "operator": "sqrt",
        "value1": 16
    }
}
```

command line - information about number of valid operations:
```
[INFO] Valid operations: 2 (out of 2)
```

result - *output.txt*:
```
obj2: 4
obj1: 5

```

## Additional information:

Source file, inputs and executable file are located in [`bartosz.rogowski/source`](./bartosz.rogowski/source/) directory.