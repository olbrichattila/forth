## Fort script in go

This is a subset of forth language written for Coding KATA, It supports:

```
	+ :  Add,
	- : Sub,
	* : Multiply,
	dup : Dup,
	drop : Drop,
	swap : Swap,
	. : Print,
	: : Word (function),
	; : End of function,
```

### examples: 
- prints 30
```
10 20 + .
```

- prints:
```
30
1500
``

```
10 20 + . 50 * .
```
- Word: function Prints 20

```
:hello1 10 10 +;
hello .
```

## Usage:

- use go run ./cmd or build it:
```
make build
```

### Command line prompt mode:
This will give you a prompt waiting for you to type your command

```
./forth 
```

### Running a file
Example
```
./forth examples/example.forth

// where examples/example.forth is a file name
```

Note:
I've experimented with creating a script language before, so here is a more complex example in my github account
(note some refactoring code cleaning required, just experiment)

https://github.com/olbrichattila/goscriptinglanguage

this is a more complex example:
- complex mathematical expression evaluation: example: 1 + 5 + 6 /2 * (15 - 2) 
- variable declaration (including variable scope management) let foo = 1; 
- constant declaration (including variable scope management) const x = 8; 
- variable re assignments
- complex object definition
- internal functions
- conditional expressions
- if / else statement
- for loop (different versions, like do while and for)
- break
- switch case
- strings
so far, maybe to be continued.

