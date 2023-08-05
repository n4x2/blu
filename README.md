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
You can add custom validation rules by implementing the `Rule` interface and registering them with the validator. This allows you to extend the validation capabilities as per your specific requirements.
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
Use the "custom" rule in your struct's validation tags:
```go
type MyStruct struct {
    FieldA string `validate:"custom"`
    // ...
}
```
## Supported Rules
- `alpha_dash`: Validates that the field contains only alpha-numeric characters, underscores, and dashes.
- `alpha_num`: Validates that the field contains only alpha-numeric characters.
- `alpha_space`: Validates that the field contains only alpha-numeric characters and spaces.
- `alpha`: Validates that the field contains only alphabetic characters.
- `enum`: Validates that the field's value is within a predefined set of allowed values.
- `max`: Validates that the numeric field's value is less than or equal to a specified maximum value.
- `min`: Validates that the numeric field's value is greater than or equal to a specified minimum value.
- `number`: Validates that the field contains a valid number.
- `numeric`: Validates that the field contains only numeric characters.
- `required`: Validates that the field is present and not empty.
