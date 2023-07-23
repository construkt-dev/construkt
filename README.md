# Construkt
Construkt is a toolkit for creating Kubernetes operators. View the docs at [construkt.dev](https://construkt.dev).

> _Disclaimer:_ Construkt is in very early stages of development. It is not ready for production use.
> Feel free to try it out and provide feedback, but be aware that the API is not stable yet and we
> can't give any guarantees about stability and safety.

## Design Goals

Construkt aims to provide a toolkit for creating things on top of Kubernetes. We do this with the following goals in mind:
- **No boilerplate**: We want to provide a simple API that allows you to focus on the business logic of your operator.
- **No Code Generation**: We want to avoid code generation as much as possible. We believe that code generation is a
  powerful tool, but it also comes with a lot of complexity and can be hard to keep up-to-date.
- **Opinionated but Extensible**: We want Construkt to work out of the box for the majority of small-scale operators, but also allow you to extend it to fit your needs.
- **Testable**: We want to make it easy to write tests for your operator. We also want to make it easy to test your operator against a real Kubernetes cluster.
- **Easy To Build With**: We want to make the API surface as easy to work with as possible. We try and provide the necessary syntactic sugar to make your life easier, even if that results in not fully idiomatic Go _(whatever that even means)_.

## Getting Started

### Installation

Construkt is available as a Go module. You can install it with the following command:

```shell
go get github.com/construkt-dev/construkt
```

### Usage

See the [docs at construkt.dev](https://construkt.dev) for detailed guides, API documentation and reference material.

See the [examples](examples) directory for a variety of examples on how to use Construkt.