Go packages: http://pkg.go.dev

Install package using `go get {package}`, example `go get rsc.io/quote`, then it will generate a new file called `go.sum` and added package information there and will be used in `go.mod`.<br>
Alternatively you can use `go mod tidy` for auto-installing required packages.