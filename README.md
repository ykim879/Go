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
