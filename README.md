# Gocek

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/gocek?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/gocek)](https://goreportcard.com/report/github.com/restuwahyu13/gocek)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/gocek/blob/master/CONTRIBUTING.md)


**Gocek** is a simple tools for BDD / TDD testing assertion library for golang, that can be delightfully paired with any golang testing framework or without framework.


## Table Of Content

- [Gocek](#gocek)
  - [Table Of Content](#table-of-content)
  - [Installation](#installation)
  - [API Reference](#api-reference)
  - [Testing](#testing)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

```bash
go get github.com/restuwahyu13/gocek
```

## API Reference

| **Method**                 | **Params** | **Type Data**    | **Description**                                                                                                           |
| -------------------------- | ---------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------- |
| **Expect**                 | **value**  | **interface{}**  | `Expect` function is used every time you want to test a value.                                                            |
| **Not**                    | **none**   | **-**            | `Not` If you know how to test something, .not lets you test its opposite                                                  |
| **ToBe**                   | **value**  | **interface{}**  | `ToBe` to compare primitive values, but not support for type object like slice, map, struct or array.                     |
| **ToBeFalsy**              | **none**   | **-**            | `ToBeFalsy` when you don't care what a value is and you want to ensure a value is false in a boolean context.             |
| **ToBeGreaterThan**        | **value**  | **interface{}**  | `ToBeGreaterThan `to compare received > expected for number or big integer values                                         |
| **ToBeGreaterThanOrEqual** | **value**  | **interface{}**  | `ToBeGreaterThan` to compare received >= expected for number or big integer values                                        |
| **ToBeInstanceOf**         | **value**  | **reflect.Kind** | `ToBeInstanceOf` to check that an object is an instance of a class. This matcher uses instanceof underneath               |
| **ToBeLessThan**           | **value**  | **interface{}**  | `ToBeLessThan`to compare received < expected for number or big integer values                                             |
| **ToBeLessThanOrEqual**    | **value**  | **interface{}**  | `ToBeLessThanOrEqual` to compare received <= expected for number or big integer values                                    |
| **ToBeMinus**              | **none**   | **-**            | `ToBeMinus` when you don't care what a value is and you want to ensure a value is minus in a float or integer context.    |
| **ToBeNil**                | **none**   | **-**            | `ToBeNil` when you don't care what a value is and you want to ensure a value is nil context.                              |
| **ToBeTruthy**             | **none**   | **-**            | `ToBeTruthy` when you don't care what a value is and you want to ensure a value is false in a boolean context.            |
| **ToBeZero**               | **none**   | **-**            | `ToBeZero` when you don't care what a value is and you want to ensure a value is zero in a integer or float context.      |
| **ToContain**              | **value**  | **string**       | `ToContain` when you want to check that an item is in an string                                                           |
| **ToEqual**                | **value**  | **interface{}**  | `ToEqual` to compare recursively primitive values or all properties of object instances like slice, map, struct or array. |
| **ToHaveLength**           | **value**  | **int**          | `ToHaveLength` to check that an object has a length property and it is set to a certain numeric value.                    |
| **ToHaveReturned**         | **none**   | **-**            | `ToHaveReturned` to check function returned value or not                                                                  |
| **ToMatchObject**          | **value**  | **interface{}**  | `ToMatchObject` to check that a slice, map, struct or array matches a subset of the properties of an object.              |
| **ToMatch**                | **regex**  | **string**       | `ToMatch` to check that a string matches a regular expression.                                                            |
| **ToBeError**              | **none**   | **-**            | `ToBeError` when you don't care what a value is and you want to ensure a value is error context.                          |

## Testing

- Testing Via Local

  ```sh
  go test .
  ```

- Testing Via Docker

  ```sh
  docker build -t gocek --compress . && docker run gocek go test --cover -v --failfast .
  ```

## Bugs

For information on bugs related to package libraries, please visit [here](https://github.com/restuwahyu13/gocek/issues)

## Contributing

Want to make **gocek** more perfect ? Let's contribute and follow the
[contribution guide.](https://github.com/restuwahyu13/gocek/blob/main/CONTRIBUTING.md)

## License

- [MIT License](https://github.com/restuwahyu13/gocek/blob/master/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2.5rem;">
  <b><a href="#gocek">BACK TO TOP</a></b>
</p>
