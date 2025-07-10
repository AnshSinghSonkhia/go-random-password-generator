
# go-random-password-generator

A simple, secure, and highly customizable command-line tool written in Go for generating random passwords. This tool uses cryptographically secure random number generation to ensure strong, unpredictable passwords for your accounts, applications, or scripts.

---

## Features

- **Cryptographically secure** password generation using Go's `crypto/rand` package
- **Customizable character sets**: include/exclude lowercase, uppercase, numbers, and symbols
- **Configurable password length**
- **Easy to use** with clear command-line flags
- **Sensible defaults**: if no character set is specified, all are included

---

## Installation

<!-- ### Prerequisites
- [Go](https://golang.org/dl/) 1.13 or higher installed -->

### Build from Source

```powershell
# Clone the repository (if not already cloned)
git clone <repo-url>
cd go-random-password-generator

# Build the executable
go build -o passwordgen.exe
```

---

## Usage

Run the program from your terminal or command prompt:

```powershell
# Basic usage (generates a 12-character password with all character types)
./passwordgen.exe

# Specify password length
./passwordgen.exe -length 20

# Include only lowercase and numbers
./passwordgen.exe -lower -number -length 16

# Include uppercase, symbols, and numbers
./passwordgen.exe -upper -symbol -number -length 24
```

### Run Directly with `go run`

You can also run the tool without building the executable by using `go run`:

```powershell
go run main.go
go run main.go -length 20
go run main.go -lower -number -length 16
go run main.go -upper -symbol -number -length 24
```

All command-line flags work the same way as with the built executable.

### Command-Line Flags

| Flag      | Description                        | Default |
|-----------|------------------------------------|---------|
| `-length` | Length of the password             | 12      |
| `-lower`  | Include lowercase letters          | false   |
| `-upper`  | Include uppercase letters          | false   |
| `-number` | Include numbers                    | false   |
| `-symbol` | Include symbols                    | false   |

> **Note:** If you do not specify any of the `-lower`, `-upper`, `-number`, or `-symbol` flags, the generator will include all character types by default.

#### Example Outputs

```
Generated password: 8@kQw!2zR#pL
Generated password: 7g4n2v9b1m0q
Generated password: XZQJYVBNMLKJ
Generated password: 1aB!2cD@3eF#
```

---

## Security Notes
- Passwords are generated using Go's `crypto/rand` for strong randomness.
- Always use a sufficiently long password (at least 12 characters recommended).
- Do **not** share your generated passwords with others.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to open an issue or submit a pull request.

---

## Author

- [Ansh Singh Sonkhia](https://github.com/AnshSinghSonkhia)

---

Enjoy secure password generation! 🔒