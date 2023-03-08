# federation-benchmarks

Comparison of a different federation gateways, e.g. supergraphs implementations

## Benchmark results

| Supergraph Gateway          | Benchmark Duration | Slowest     | Fastest     | Average     | Requests/sec     | Latency distribution 95% | Total response count |
|-----------------------------|--------------------| ----------- | ----------- | ----------- |------------------|--------------------------|----------------------|
| Wundergraph RPC Endpoint    | 10.0016 secs       | 0.4944 secs | 0.0001 secs | 0.0043 secs | 230463 (100%)    | 0.0071 secs              | 23042.64 (100%)      |
| Wundergaph Graphql Endpoint | 10.0200 secs       | 0.0767 secs | 0.0001 secs | 0.0034 secs | 293560 (+27.38%) | 0.0068 secs              | 29297.45 (+27.14%)   |
| Apollo Gateway              | 10.0493 secs       | 0.2480 secs | 0.0553 secs | 0.0920 secs | 10911 (-95.27%)  | 0.1173 secs              | 1085.75 (-95.29%)    |
| Apollo Router               | 10.0162 secs       | 0.1663 secs | 0.0046 secs | 0.0248 secs | 40281 (-82.52%)  | 0.0377 secs              | 4021.6 (-82.55%)     |
| Graphql Mesh                | 10.0700 secs       | 0.5075 secs | 0.0698 secs | 0.2566 secs | 3905 (-98.31%)   | 0.3369 secs              | 387.79 (-98.32%)     |
| Mercurios                   | 10.0842 secs       | 0.1330 secs | 0.0263 secs | 0.0346 secs | 29092 (-87.38%)  | 0.0437 secs              | 2884.91 (-87.48%)    |


Tests were done on the MacBook Pro (16-inch, Apple M1 Pro, 2021) with 16GB of RAM

More detailed results could be found in the [results](./results) directory.

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
