# Desmos Docs
This folder contains the documentation of Desmos Network hosted on https://docs.desmos.network.  
Feel free to open issues/PRs to contribute! 

## Installation

```console
yarn install
```

## Local Development

```console
yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live
without having to restart the server.

## Update graphql schemas (if needed)
If a new Desmos release also updated graphql schemas, it's required to also update the related schema files. To do it, run the following command in the project directory:

```console
npx docusaurus graphql-to-doc
```

If you want to read more about the graphql generator please check [here](https://www.npmjs.com/package/@edno/docusaurus2-graphql-doc-generator) the docs of the plugin we are using.

## Build

```console
yarn build
```

```console
docusaurus build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

## Serve

```console
docusaurus serve
```

### Credits 
> Docs powered by [Docusaurus 2](https://docusaurus.io/), a modern static website generator.
