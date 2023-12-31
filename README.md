# Go
Go is a complied language. The toolchain convert source language to native. The source code begins with package of the code, import statement, and program code. This is mostly referenced from the book The Go Programming Language by Alan and Brian, having same structure as well.
## Table of Content
1. [Program Structure ](#item_one)
2. [Basic Data Types ](#item_two)

<a id = "item_one"></a>
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

<a id = "item_two"></a>
## 2 Basic Data Types
### 2.1 String
String is immutalbe so two copies of a string share same underlying memory, including substring. To manipulate slices of bytes, one can use bytes package. Since strings are immutable, building up strings incrementally an involves a lot of allocation and copying. Thus, using bytes.Buffer type will be helpful. \
Therefore substring operations are cheap. String uses escape sequence. However, you can use ` instead of double quote for raw string literal. Raw string literal won't process escape sequence. Therefore, it is good for HTML templates, JSON literals, command messages, etc.
#### 2.1.1 Packages
There are strings package, bytes package, strcnv package, unicode package, and etc for string data type. strcnv provides functions for converting boolean, integer, floating-point value to and from their string representations, and function for quoting and unquoting strings. unicode package provides IsDigit, IsLetter,IsUpper, IsLower, and etc which some are provided by string package as well.
##### (1) strings
###### HasPrefix
``` Go
strings.HasPrefix(str, prefix)
```
##### (2) Buffer
As string it self is immutalbe, Buffer type is efficient for amnipulating byte slices. Buffer starts out empty and grow as string, byte, and []byte grows into it. 
### 2.2 Constants
Constans is a value known to complier and whose evaluation is guaranteed to occur at compile time. Benefits from compile time are:
(1) reduce work on runtime which enable other compiler optimization.
(2) Errors are caught in compile time, which reduces runtime errors.

When sequence of constants is declared as a group, right-hand side can be ommited with imply previous expression is being used.
``` Go
const (
a = 1
b // b is 1
)
```
#### 2.2.1 iota : Constant Generator
iota create a sequence of related values without spelling out. Use cases are below
``` Go
type Weekday int
const (
Sunday Weekday = iota // generate value starting with zero
Monday // Monday is 1
Tuesday // 2
...
)
```
``` Go
type Flags unit
const (
FlagUp Flags = 1 << iota // 1
FlagBroadcast // 1 << 1 so two and sucessive will be power of previous
...
)
```

## 3. Composite Types
In Go, there are four composite types: arrays, slices, maps, and structs. Array is merely used because of fixed-length and alternatively slices are used because of its versatile.
### 3.1 Array
There are several attributes about array:
1. the size of the array is part of its type. Also it is computed on compiled time so different size cannot be placed during runtime.
2. to be both array same, they need to store same values for same indexes and same size.
Here are few example codes regarding array. 
``` Go
for i, v := range arr {
...
} // loop around array, if index is not usable, then replace it with _
q := [...]int{1,2,3} // size can be ommitted if it starts with initilizer
r := [...]int{99:-1} // this defines array w with 100 eleemnts all zero except for the last, which is value -1
```
### 3.2 Slices
1. variable-length sequce which is lightweight data structure that gives subsequence of the elements of the array (underlying array).
2. not comparable through ==.
3. zero value is nil, where lenght and capacity is 0.
#### 3.2.1 Composition
it has tree components: pointer, length, capacity. Length and capacity can be accessed by built-in function, len and cap. When the slice operator creates the slice of the array, it refers with the pointer. Therefore, the change of the slice will apply on original array and vice versa. The slice can be expanded by using [:j] operator if j does not exceed capacity
``` Go
// assume arr is array of size 12
slice = arr[6:9] // slice points to array[6] with len of 3 and capacity 9.
expanded = slice[:5] // expanded the slice length. 
```
#### 3.2.2 make function
it creates a slice of specified element type, length and optional capacity.
``` Go
make([]T, len)
make([]T, len, cap)
```
#### 3.2.3 append function
append value to the slice. when the capacity is less than the elements it is adding, it grows by growth strategy by copying one slice to another. Therefore, the runtime is constant in average.
``` Go
var x []int // initialize slice
x = append(x,1) //append value with growing the slices.
```
### 3.3 Maps
1. map type is written map[K]V where K and V is type of key and value
2. all keys need to be same type and values need to be same type but key and value do not have to be.
3. map eleement is not a variable so it does not have adress. It is because maps grows, the exisiting elements may change their storage during runtime.
4. key and value are retrieved by range function.
5. set does not exist in Go, therefore map[T]boolean can replace sets
#### 3.3.1 Initialization
``` Go
ages := map[string]int{
"alice": 31
}
```
``` Go
ages := make(map[string]int)
ages["alice"] = 31
```
#### 3.3.2 delete
``` Go
delete(ages, "alice")
```
#### 3.3.3 retrieve
If the key does not exist, it will give zero value. To know if the key exists in map, map[k] returns value of the key and boolean to see if the key exists.
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
