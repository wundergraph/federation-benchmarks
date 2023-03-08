# federation-benchmarks

Comparison of a different federation gateways, e.g. supergraphs implementations

## Benchmarks



## Prerequisites

You need to have install the next tools:

- [Go 1.17+](https://go.dev/doc/install)
- [Hey](https://github.com/rakyll/hey)
- [Taskfile](https://taskfile.dev/)
- Npm
- Curl

## How to run

### Default command

```shell
task
```

This command will run default task from the Taskfile.yml file, which includes setuping every gateway, running benchmarks and generating reports.

### Available commands

You may list all possible tasks with the following command:

```shell
task --list-all
```

Taskfile is splitted into smaller files responsible for different tasks, e.g. `setup`, `benchmark`, `query`, `subgraphs`, `run`

### Reports

The generated reports are located in the `results` folder.

