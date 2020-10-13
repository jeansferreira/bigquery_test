# bigquery

Project to abstract Google Big Query interactions. To see more about Google Big Query service access this link [https://cloud.google.com/bigquery/docs](https://cloud.google.com/bigquery/docs)

## Why

We create this lib to avoid Big Query coupling and source code duplication. This lib is build on the top of [Google Big Query client](https://godoc.org/cloud.google.com/go/bigquery).
See https://godoc.org/cloud.google.com/go for authentication, timeouts, connection pooling and similar aspects of Google clients.


## Simple usage:

TODO

## Testing

Run:

```
make check
```
Open generated coverage on a browser:

```
make coverage
```
To perform static analysis:

```
make analyze
```

## Releasing

Run:

```
make release version=<version>
```

It will create a git tag with the provided **<version>**
and build and publish a docker image.

## Git Hooks

To install the project githooks run:

```
make githooks
```
