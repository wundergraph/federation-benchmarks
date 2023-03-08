# federation-benchmarks

Comparison of a different federation gateways, e.g. supergraphs implementations

Found out more about Wundergraph Federation Gateway [here](https://wundergraph.com/use-cases/apollo-federation-gateway)

## Benchmark results

Sample command to run benchmarks:

```shell
hey -z 10s -c 100 -m POST -H 'content-type: application/json' -D queries/small.json 'http://localhost:9991/graphql'
```

| Name                        | Slowest     | Fastest     | Average     | Req/sec            | P95  (Latency)     |
|-----------------------------| ----------- | ----------- |-------------|--------------------|--------------------|
| Wundergaph Graphql Endpoint | 0.0767 secs | 0.0001 secs | 0.0034 secs | 29297.45 (100%)    | 0.0068 (100%)      |
| Wundergraph RPC Endpoint    | 0.4944 secs | 0.0001 secs | 0.0043 secs | 23042.64 (-21.35%) | 0.0071 (+4.41%)    |
| Apollo Router               | 0.1663 secs | 0.0046 secs | 0.0248 secs | 4021.6 (-86.27%)   | 0.0377 (+454.41%)  |
| Mercurios                   | 0.1330 secs | 0.0263 secs | 0.0346 secs | 2884.91 (-90.15%)  | 0.0437 (+542.65%)  |
| Apollo Gateway              | 0.2480 secs | 0.0553 secs | 0.0920 secs | 1085.75 (-96.29%)  | 0.1173 (+1625%)    |
| Graphql Mesh                | 0.5075 secs | 0.0698 secs | 0.2566 secs | 387.79 (-98.68%)   | 0.3369 (+4854.41%) |


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
