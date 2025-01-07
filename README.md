## Fort script in go

his is a subset of the Forth language, written for a Coding KATA.

The code has some test coverage, and I've split it into smaller chunks.

It could be solved with less code. The reasoning behind this is to follow the methodology of how scripting languages work.

### The workflow.
1. Lexer: Tokenizes the source code.
2. AST Generation: Creates an Abstract Syntax Tree (AST).
3. Interpreter: Evaluates and executes the code represented by the AST.

### Tokenizing:
The main purpose of the lexical analyzer is to break down raw source code into manageable pieces called tokens.

### Abstract Syntax Tree (AST)
The Abstract Syntax Tree (AST) is a crucial data structure used in programming language interpreters and compilers. Its primary purpose is to provide a structured, hierarchical representation of the source code, making it easier to analyze and manipulate compared to raw text or tokens.

### Interpreter
The interpreter executes the AST, often recursively, by traversing its node structure. In this example, each word (or function) encapsulates its own AST, enabling modular execution.


### It supports:
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

word (function declaration)
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
```

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
I experimented with creating a scripting language some time ago, and here is a more complex example on my GitHub account. (Note: some refactoring and code cleaning are required, as this was just an experiment.)

https://github.com/olbrichattila/goscriptinglanguage

This example includes the following features:
- Complex mathematical expression evaluation (e.g., 1 + 5 + 6 / 2 * (15 - 2))
- Variable declaration with scope management (e.g., let foo = 1;)
- Constant declaration with scope management (e.g., const x = 8;)
- Variable reassignments
- Complex object definitions
- Internal functions
- Conditional expressions
- if / else statements
- Various for loop versions (e.g., do while and for)
- break statements
- switch / case statements
- Strings
So far, but it may be continued.
