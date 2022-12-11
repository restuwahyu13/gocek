# Gocek

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/gocek?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/gocek)](https://goreportcard.com/report/github.com/restuwahyu13/gocek)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/gocek/blob/master/CONTRIBUTING.md)


**Gocek** is a simple BDD / TDD gocekAssertion library for golang that can be delightfully paired with any golang testing framework.


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

| **Method**                 | **Params**             | **Description**                                                                                                        |
| -------------------------- | ---------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **ToBe**                   | **value interface{}**  | ToBe to compare primitive values or to check referential identity of object instances.                                 |
| **ToBeFalsy**              | **None**               | ToBeFalsy when you don't care what a value is and you want to ensure a value is false in a boolean context.            |
| **ToBeGreaterThan**        | **value interface{}**  | ToBeGreaterThan to compare received > expected for number or big integer values                                        |
| **ToBeGreaterThanOrEqual** | **value interface{}**  | ToBeGreaterThan to compare received >= expected for number or big integer values                                       |
| **ToBeInstanceOf**         | **value reflect.Kind** | ToBeInstanceOf(Class) to check that an object is an instance of a class. This matcher uses instanceof underneath       |
| **ToBeLessThan**           | **value interface{}**  | ToBeLessThan to compare received < expected for number or big integer values                                           |
| **ToBeLessThanOrEqual**    | **value interface{}**  | ToBeLessThanOrEqual to compare received <= expected for number or big integer values                                   |
| **ToBeMinus**              | **None**               | ToBeMinus when you don't care what a value is and you want to ensure a value is minus in a float or integer context.   |
| **ToBeNil**                | **None**               | ToBeNil when you don't care what a value is and you want to ensure a value is nil context.                             |
| **ToBeTruthy**             | **None**               | ToBeTruthy when you don't care what a value is and you want to ensure a value is false in a boolean context.           |
| **ToBeZero**               | **None**               | ToBeZero when you don't care what a value is and you want to ensure a value is zero in a integer or floatcontext.      |
| **ToContain**              | **value string**       | ToContain when you want to check that an item is in an string                                                          |
| **ToEqual**                | **value interface{}**  | ToEqual to compare recursively primitive values or all properties of object instances (also known as "deep" equality). |
| **ToHaveLength**           | **value int**          | ToHaveLength to check that an object has a length property and it is set to a certain numeric value.                   |
| **ToHaveReturned**         | **None**               | ToHaveReturned to check function returned value or not                                                                 |
| **ToMatchObject**          | **value interface{}**  | ToMatchObject to check that a slice, map, struct or array matches a subset of the properties of an object.             |

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
