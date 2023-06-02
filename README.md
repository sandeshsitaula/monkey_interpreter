### Monkey Interpreter: A simple recursive tree walking interpreter developed in go.

### Installing:

```
git clone https://github.com/sandeshsitaula/monkey_interpreter/
cd monkey_interpreter
go install 
```


If you dont have Go you can try it by going to directory and running ./monkey  (Windows exe file not Available)

### Usage
```
monkey  [Starts interpreter in REPL MOde]
monkey <filename.mn>  [mn as extension]

eg.monkey main.mn  or Without Go   ./monkey main.mn
```

It is syntactically Similar to C. But doesnot use data types explicitly 
```
eg.let a=12;  to define integer variable
````
Data Types Supported:

All primitive data types like int,char,float,bool<br>
Strings , Array <br>
Hashmap not implemented

### Examples:
``` golang
//primitive types and if expression
let a=12;
puts(a) ; //prints value of a
let val=if (a>18){return "can vote"} else {"cannot vote"}; //elseif not implemented
puts(val);  //prints cannot vote        

//Array
let arr=[1,"san",true];
puts(arr[1]) //gives san as output
puts(len(arr));  //prints 3  ;Here len is a builtin function

//Function as first Class citizen
let age = fn (x){return x};
puts(age(18))             //prints 18

let sum=fn(a,b){a+b;}     //return can be ommitted if it is last expression
puts(sum(10,15));         //prints 25

//String
let str1="hello";
let str2="world";
let str=str1+" "+str2;
puts(str) //prints hello world
puts(len(str)) //prints the length of hello world
```
### NOte: In REPL mode each line is scanned and given to lexer so  functions,ifstatements have to be written in same line.<br>
### NOTE: Repl mode is only for testing purpose  

### Credits to : Writing an interpreter in Go book  By Thorsten Ball
