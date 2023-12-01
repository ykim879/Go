# Go
Go is a complied language. The toolchain convert source language to native. The source code begins with package of the code, import statement, and program code. This is mostly referenced from the book The Go Programming Language by Alan and Brian, having same structure as well.
## 1. Program Structure 
### 1.1 Initialization
Format:
1. s := ""  - short initialization gives appropriate types based on initialized value.
2. var s string - empty initialization. Go initialize value with empty value when it is no explicit declaration. Using var followed with =. It needs to have either a declaration of type or an assignment value. 
3. var s = ""
4. var s string = "" - (3) and (4) is not recommended
#### 1.1.1 Tuple Assignment
``` Go
x,y = y, x+y
```
All of the right-hand are evaluated before any of the variables are updated, so it is useful when we use values that are updating at the same time
#### 1.1.2 Type Declaration
``` Go
type Celsius float64
type Fahrenheit float64
funct CToF(c Celsius) Fahrenheit {..}
```
Even the underlying type could be the same, if the type is different, they are not the same time. therefore comparison or assignment wouldn't happen because of different types. However, conversion will happen if the underlying type or same or pointer types point to variables. The conversion function can be manually written by below as well:
``` Go
func (c Celsius) String() String (return fmt.Sprintf("%gC", c)}
```
By using the above, below will execute the above function. Consider c is Celsius type.
``` Go
fmt.Println(c.String())
fmt.Printf("%v", c)
fmt.Printf("%s", c)
fmt.Println(c)
```
fmt package automatically executes the function.
### 1.2 Pointers
Go uses Pointers as same as C, & means address and * is value they are pointint
``` Go
x := 1
p := &x //pointer that points to x, stroing addresss of x.
*p = 2 //change x value
```
#### 1.2.1 new Function
initialize zero value and return address.
``` Go
p := new(int) // p is pointer with integer zero value, 0.
```
### 1.3 Lifetime and Scope of Variables
The scope is a compile-time property while lifetime is runtime. 
The scope is defined by implicit/explicit lexical blocks while the lifetime is until variables become unreachable. When the variables done with their lifetime, they can be recycled by garbage collector. The complier allocate local variables on the heap for long-lived and stack for short-lived.
## For Loop
There is only one recursive function in Go, which is for loop. It is consisted of for initialization, condition, and update after each loop. The for loop also can consisted of condition statement or it can have none of them which will be infinite loop.
``` Go
for i:= 1; i < len(os.Args); i++ {
...
}
```
### 1.3 Package 
#### 1.3.1 init Function
init function can't be called or referenced and automatically executed when the program starts. It becomes handy when there is package level variables that should be precomputed.
#### 1.3.2 Order of Compiling
Initialization proceeds from the bottom up; the main package is the last to be initialized. If a has dependency on b, b initialized before a.
## Function
``` Go
func name(f param) float64 {
return f / 9
}
```
Declaration of function is done by funct, name, list of parameters and return list if it has any.
### Range
It gives index and value for each element in array
``` Go
for _, arg := range arr[1:] {
...
}
```
## Standard Libraries
### default
#### strings
##### HasPrefix
``` Go
strings.HasPrefix(str, prefix)
```
### 1. os
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
### 2. bufio
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
### 3. http
#### Get
``` Go
resp, err := http.Get(url)
```
Get makes HTTP request and return result in the structure of resp, if there is no error, and err, if any. field of resp contains server response as a readable stream. It needs to be closed after it is done using to avoid leaking the resources.
### 4. ioutil
#### ReadAll
``` Go
body, err := ioutil.ReadAll(resp.Body)
```
reads the entire response. returns a response and the error message if it has any.
## 4. Composite Types
### 4.3 Map
a set of key/value pairs (like a dictionary)
#### Initialization
The built-in function make creates a new empty map. The below map initialized map where key is string and value is integer.
```Go
counts := make(map[string]int)
```
