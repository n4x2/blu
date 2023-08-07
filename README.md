# blu
[![codecov](https://codecov.io/gh/n4x2/blu/branch/master/graph/badge.svg?token=FH1Z1IO7OM)](https://codecov.io/gh/n4x2/blu)
[![Go Report Card](https://goreportcard.com/badge/github.com/n4x2/blu)](https://goreportcard.com/report/github.com/n4x2/blu)

blu is struct value validation based on provided tags.

> **Note:** As of the current version, "blu" is intended for personal use.

## Usage
Install blu
```sh
go get github.com/n4x2/blu
```
Then import it:
```go
import "github.com/n4x2/blu"
```
Define struct with validation tags:
```go
type MyStruct struct {
    FieldA string `validate:"required"`
    // ...
}
```
Create a new validator instance:
```go
v := blu.NewValidator()
```
Validate the struct using the created validator:
```go
err := v.Validate(MyStruct)
if err != nil {
    // Handle error
}
```
## Custom Rule
Create custom rule by implementing the `Rule` interface:
```go
type CustomRule struct {}

func (r *CustomRule) Name() string {
    return "custom"
}

func (r *CustomRule) Validate(field string, value string, param string) error {
    // Implement your custom validation logic here.
    // Return an error if the validation fails.
    return nil
}
```
Register the custom rule:
```go
customRule := &CustomRule{}
err := v.RegisterRule(customRule)
if err != nil {
    // Handle error (e.g., duplicated rule name).
}
```
Use the "custom" rule in struct validation tags:
```go
type MyStruct struct {
    FieldA string `validate:"custom"`
    // ...
}
```