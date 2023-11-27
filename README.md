# Go
Go is a complied language. The toolchain convert source language to native. The source code begins with package of the code, import statement, and program code. 
## Initialization
Format:
1. s := ""  - short initialization gives appropriate types based on initialized value.
2. var s string - empty initialization. Go initialize value with empty value when it is no explicit declaration
3. var s = ""
4. var s string = "" - (3) and (4) is not recommended
## For Loop
There is only one recursive function in Go, which is for loop. It is consisted of for initialization, condition, and update after each loop. The for loop also can consisted of condition statement or it can have none of them which will be infinite loop.
``` Go
for i:= 1; i < len(os.Args); i++ {
...
}
```
### Range
It gives index and value for each element in array
``` Go
for _, arg := range arr[1:] {
...
}
```
## Standard Libraries
### os
#### Printf
##### Percent Sign
- %d: decimal integer
- %s: string
- %f: floating-point number
- %t: boolean
- %v: default format
- %T: type of the value
- %p: pointer
#### Open
open files and return two value, *os.File that is read by Open function and error value. If Error value is nil the file is opened successfully. To read the error message, there is print method for standard error stream which is fmt.Fprintf.
``` Go
f, err := os.Open(arg)
if err != nil {
fmt.Fprintf(os.Stderr, "%v\n", err)
}
```
### bufio
#### Scanner (bufio.NewScanner(f)
reads input and breaks it into lines or words.
- Scan(): reads the next line and removes newline character. return true if there is new line false if there is no input.
- Text(): returns the results that is run by Scan()
``` Go
input := bufio.NewScanner(os.Stdin)
for input.Scan() { //this will read from user's keyboard
output := input.Text() //this will read value that is inputted
...
}
```
## 4. Composite Types
### 4.3 Map
a set of key/value pairs (like a dictionary)
#### Initialization
The built-in function make creates a new empty map. The below map initialized map where key is string and value is integer.
```Go
counts := make(map[string]int)
```
