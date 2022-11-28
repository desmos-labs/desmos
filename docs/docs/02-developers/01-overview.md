---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---
# Overview

## Introduction
[Desmos](../01-intro.md) aims to provide developers a protocol with which they will be able to create decentralized and censorship-resistant social enabled apps. Different apps, with different scopes and their own Term of Services will be able to use the features offered by Desmos to customize their user experience in a unique way.
 
If you want to know more about the base concepts of a blockchain and understand some key points, please take a look at the [FAQ page](07-faq.md). 

## Core features
The core features of Desmos are organised in **modules** following the specification of the [Cosmos-SDK](https://docs.cosmos.network/main/building-modules/intro.html).   

Here a brief description of each one of these:

* `Profiles`: Handles the creation and management of a decentralized identity and its own links with both your other chains wallets and centralised applications;

* `Relationships`: Handles the creation and management of mono-directional and bidirectional [relationships] between users' wallets. It also allows managing users blocks lists;

* `Subspaces`: Handles the creation and management of a [subspace] and their [sections] inside Desmos;

* `Posts`: Handles the creation and management of posts and their contents. These contents can include a variety of different attachments such as medias (pics, gifs, videos) and polls. Posts can also be enriched with a variety of [entities].

* `Reactions`: Handles the creation and management of reactions to posts;

* `Reports`: Handle the creation and management of posts' and users' reports.

## Support features
These features are not directly connected to the social-networks scope but serve the network maintainers and
external services.

* `Fees`: Allows setting custom additional fees to modules' messages;
* `Supply`: Allows retrieving information about a particular token total and circulating supply.

If you want to know more about each module, its concepts and how to interact with them check the [modules] section.