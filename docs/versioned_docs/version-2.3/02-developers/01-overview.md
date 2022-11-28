---
id: overview
title: Overview
sidebar_position: 1
---
# Developers Overview

## Introduction
[Desmos](../01-intro.md) aims to provide developers a platform on which they will be able to create decentralized and 
censorship-resistant social network applications. 
To do so, we've implemented a set of transactions that are useful to perform the most common operations related to this world. 
If you want to know more about the base concepts of a blockchain and understand some key points, please take a look a the [FAQ page](06-developer-faq.md). 

## Glossary
Before digging into the available transactions, let's clarify the meaning of some terms that we will be using a lot.

* A **profile** contains a series of (personal) data associated to an account that a user can create on the chain;

* A **post** is a public message that everyone can read on the chain.  
  When creating it you can also specify if it allows
  to be commented on or not;
  
* A **comment** is a post that has been linked to a parent post;

* A **reaction** is the way that allows users to express a feeling on a specific post;

* A **subspace** is a "zone" where a specific app or more apps can live on and share contents;

If you want to know more about how we store the data on-chain and all the chain types, please refer to
the __"Types" section__.

## Performing transactions

If you want to know more about performing transactions to change the current chain state, please go to the __"Transactions" section__.

## Querying data

If you want to know all the GRPC endpoints and CLI commands available to query the existing chain state and the
saved data, please go to the __"Queries" section__. 
