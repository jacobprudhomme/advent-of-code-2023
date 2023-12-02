# Advent of Code 2023
This year, for my 4th Advent of Code, I have decided to use Go, to learn the language. This is my first ever time using it.

## Execution Environment
For all problems, I used Go `1.21.3`, installed using `gobrew`.

## How to Use
To build the executable for a given day, run `go build -o ./bin ./cmd/dayXY` from the root of the project. This will output the binary in a `bin` subdirectory. Then, this program can be run by redirecting the input from the appropriate file under `inputs` (all programs expect input on `stdin`).

Alternatively, to execute the program for a given day directly, run `go run ./cmd/dayXY < inputs/dayXY > output` to obtain the output in a file.
