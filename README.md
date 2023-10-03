### Limitation and Assumption
The tool only supports converting numbers from -999999999999 to 999999999999, but
is easily extensible. Output will only show in English format

Example:
```
1.15 -> One Dollar and Fifteen Cents
```

### Requirements

[go](https://go.dev)

### Usage

build the package first then add the binary to PATH. You can also compile and run
without building the binary with `go run .`

```bash
# build using
go build -ldflags="-s -w" .

# then run it
number-to-currency-converter

# compile and run without outputing compiled binary
go run .
```

if you are on mac you can do something fun like this

```bash
echo 1000000 | number-to-currency-converter | sed -E 's/.*://' | say
```

### Test

testing is done using golang [testing](https://pkg.go.dev/testing) package

```bash
go test .
```
