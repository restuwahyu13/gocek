# Gocek

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/gocek?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/gocek)](https://goreportcard.com/report/github.com/restuwahyu13/gocek)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/gocek/blob/master/CONTRIBUTING.md)


**Gocek** is a simple BDD / TDD gocekAssertion library for golang that can be delightfully paired with any golang testing framework.


## Table Of Content

- [Gocek](#gocek)
  - [Table Of Content](#table-of-content)
  - [Installation](#installation)
  - [Testing](#testing)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

```bash
go get github.com/restuwahyu13/gocek
```

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
