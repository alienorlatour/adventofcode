# Newmath parser

## Grammar

### List of symbols

- SCA (scalar) - the only terminal symbol.
- EXP (expression) 

### Definitions

|     |             |
|-----|-------------|
| SCA | `[0-9]+`    |
| EXP | `(EXP)`     |
| EXP | `EXP + EXP` |
| EXP | `EXP * EXP` |
| EXP | SCA         |

### Turing machine

I love Turing machines...

```puml
@startuml

title Turing machine

state start
state digit : [0-9]
state symbol : +*
state parens : (count parenthesis and parse the contents)

[*] -> start
start -> digit
digit -> digit
digit -> symbol
symbol -> start

start --> parens
parens -> start

@enduml
```

