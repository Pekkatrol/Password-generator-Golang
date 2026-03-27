# Password Generator in Golang

## Goal

The goal of this project is to create a command-line password generator in Golang. The script generates a random password based on user-defined settings: password length, and whether to include letters, digits, and/or symbols.

## What I used

- `fmt` for displaying output and error messages
- `os` for handling command-line arguments and exiting on error
- `strconv` for parsing string arguments to integers and booleans
- `math/rand` for generating random characters

### You need Golang

If you don't have Golang installed, run:
```bash
sudo apt update
sudo apt install golang-go
```

## How to use it?

Run the script with the following arguments:
```bash
go run . <length> <letters> <digits> <symbols>
```
- `<length>`: integer, the length of the password (e.g., 12)
- `<letters>`: true/false, include letters
- `<digits>`: true/false, include digits
- `<symbols>`: true/false, include symbols

### Example

```bash
go run . 16 true true false
```
You should see something like:
```bash
The generated password is: 8a7B2c9D1e3F4g5H
```

If you provide invalid arguments, the script will print an error message and exit.

---
