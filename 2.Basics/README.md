[⭐]
## Since go version 1.11 we now have concept of go modules and it comes with a new environment variable $GO111MODULE.

## Its default value is 'auto'. So if your code is inside $GOPATH:
	- And does not contain go.mod file then it will compile your code the old fashioned way.
	- And contains go.mod file then it will use go modules to compile your code

## If you are not using Go modules then

	- You can use $GOPATH/src directory for your project

	- You can use directory outside $GOPATH and use 'GO111MODULE=off' to compile your code

    $GOPATH:

    Ex 1: $GOPATH/src/yourName/yourCustomPackage       
    Ex 2: $GOPATH/src/github.com/yourHandle/yourPackage

We turn off the modules feature before calling `go run` because there is no go.mod file in this directory

If your root workspace does not have a go.mod file then your project belongs inside $GOPATH
If your project is outside $GOPATH then you must have go.mod file defined using go mod init <your-package-name>

This directory does not have go.mod file hence
To run this file you will need to set GO111MODULE to 'off'

[🐧]
Ex (Linux):
go run <list of all go files>
GO111MODULE=off go run <list of all go files>

// Alternately you can type just
go run .


[🖥]
Ex: (Windows Powershell ) [Notice there is a space around '=']
$env:GO111MODULE = 'off'; go run <list of all go files>

//
go run <path-to-directory-containing-main-file>
*/