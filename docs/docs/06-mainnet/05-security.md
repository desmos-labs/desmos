---
id: security
title: Security
sidebar_position: 2
---

## Validator Security
Each validator candidate is encouraged to run its operations independently, as diverse setups increase the resilience of the network.

## Key Management System & Hardware Security Modules
It is critical that an attacker cannot steal a validator's key. If this is possible, it puts the entire stake delegated to the compromised validator at risk. HSM alongside KMS are an important strategies for mitigating this risk.

You can take a look on how to set up a KMS with or without HSM by reading [here](kms/kms.md).

## Sentry Nodes (DDOS Protection)
Validators are responsible for ensuring that the network can sustain denial of service attacks.

One recommended way to mitigate these risks is for validators to carefully structure their network topology in a so-called sentry node architecture.

Validator nodes should only connect to full-nodes they trust because they operate them themselves or are run by other validators they know socially. A validator node will typically run in a data center. Most data centers provide direct links to the networks of major cloud providers. The validator can use those links to connect to sentry nodes in the cloud. This shifts the burden of denial-of-service from the validator's node directly to its sentry nodes, and may require new sentry nodes be spun up or activated to mitigate attacks on existing ones.

Sentry nodes can be quickly spun up or change their IP addresses. Because the links to the sentry nodes are in private IP space, an internet based attacked cannot disturb them directly. This will ensure validator block proposals and votes always make it to the rest of the network.

We suggest sentry nodes to be set up on multiple cloud providers across different regions. A validator may be offline if the connected sentry nodes are all offline due to the outage of a cloud provider in a specific region. 

To setup your sentry node architecture you can follow the instructions below:

Validator Nodes should edit their config.toml:

```bash
# Comma separated list of nodes to keep persistent connections to
# Do not add private peers to this list if you don't want them advertised
persistent_peers =[list of sentry nodes]

# Set true to enable the peer-exchange reactor
pex = false
```

Sentry Nodes should edit their config.toml:

```bash
# Comma separated list of peer IDs to keep private (will not be gossiped to other peers)
# Example ID: 3e16af0cead27979e1fc3dac57d03df3c7a77acc@3.87.179.235:26656

private_peer_ids = "node_ids_of_private_peers"
```
  