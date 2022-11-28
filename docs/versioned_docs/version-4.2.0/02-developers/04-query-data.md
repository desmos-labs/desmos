---
id: query-data
title: Query data
sidebar_label: Query data
slug: query-data
---

# Query data
Inside Desmos it is possible to query data in 4 different ways:
1. With the `CLI` directly from terminal.
2. With the `REST` endpoint directly from a full node.
3. With the `gRPC` endpoint directly from a full node.
4. With `GraphQL`.

## CLI
To query data using `CLI`, you can check the following commands of each module:

* [Profiles CLI commands](02-modules/profiles/07-client.md#cli);
* [Relationships CLI commands](02-modules/relationships/06-client.md#cli);
* [Subspaces CLI commands](02-modules/subspaces/07-client.md#cli);
* [Posts CLI commands](02-modules/posts/08-client.md#cli);
* [Reports CLI commands](02-modules/reports/08-client.md#cli);
* [Reactions CLI commands](02-modules/reactions/07-client.md#cli);
* [Supply CLI commands](02-modules/supply/03-client.md#cli);
* [Fees CLI commands](02-modules/fees/06-client.md#cli).

:::info
To be able to perform the above queries, you need to have the desmos daemon installed.
Check the full node [setup section](../03-fullnode/02-setup.md#1-build-the-software) to know how.
:::

## gRPC
To query data using the `gRPC` endpoint, you can use the following endpoints:

1. [Testnet gRPC endpoint](../05-testnet/04-endpoints.md#rest--grpc)
2. [Mainnet gRPC endpoint](../06-mainnet/06-endpoints.md#rest--grpc)

The above endpoints can be combined with the following endpoints to get the desired data:
* [Profiles gRPC endpoints](02-modules/profiles/07-client.md#grpc);
* [Relationships gRPC endpoints](02-modules/relationships/06-client.md#grpc);
* [Subspaces gRPC endpoints](02-modules/subspaces/07-client.md#grpc);
* [Posts gRPC endpoints](02-modules/posts/08-client.md#grpc);
* [Reports gRPC endpoints](02-modules/reports/08-client.md#grpc);
* [Reactions gRPC endpoints](02-modules/reactions/07-client.md#grpc);
* [Supply gRPC endpoints](02-modules/fees/06-client.md#grpc);
* [Fees gRPC endpoints](02-modules/fees/06-client.md#grpc).

## REST
TO query data using the `REST` endpoint, you can use the following endpoints:
1. [Testnet REST endpoint](../05-testnet/04-endpoints.md#rest--grpc)
2. [Mainnet REST endpoint](../06-mainnet/06-endpoints.md#rest--grpc)

The above endpoints can be used with the following endpoints to get the desired data:
* [Profiles REST endpoints](02-modules/profiles/07-client.md#rest);
* [Relationships REST endpoints](02-modules/relationships/06-client.md#rest);
* [Subspaces REST endpoints](02-modules/subspaces/07-client.md#rest);
* [Posts REST endpoints](02-modules/posts/08-client.md#rest);
* [Reports REST endpoints](02-modules/reports/08-client.md#rest);
* [Reactions REST endpoints](02-modules/reactions/07-client.md#rest);
* [Supply REST endpoints](02-modules/fees/06-client.md#rest);
* [Fees REST endpoints](02-modules/fees/06-client.md#rest).

## GQL
Another way to query the Desmos data is GQL. GQL is different from the above methods because it offers high possibilities of customisation for developers based on their needs. It is possible to interact with GQL endpoints in the client you are building by using one of the many libraries available for this kind of interaction. You can check what suites your needs here: [GraphQL resources](https://graphql.org/code/).

The GQL endpoints for Desmos chains are the following:
1. [Testnet GQL endpoint](../05-testnet/04-endpoints.md#gql)
2. [Mainnet GQL endpoint](../06-mainnet/06-endpoints.md#gql)
