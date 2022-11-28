---
id: unlink-chain-account
title: Unlink chain account
sidebar_label: Unlink chain account
slug: unlink-chain-account
---

# `MsgUnlinkChainAccount`
This message allows you to remove a previously linked chain address from your Desmos profile.

## Structure
```json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnlinkChainAccount",
  "target": "<Address of the chain link to unlink>",
  "chain_name": "<Name of the chain to unlink>",
  "owner": "<Desmos address of the profile that should remove the link>"
}
```

### Attributes

| Attribute | Type | Description | Required |
| :-------: | :----: | :-------- | :------- |
| `target` | String | External address that should be unlinked | yes |
| `chain_name` | String | Name of the target external chain to unlink | yes |
| `owner` | String | Desmos address of the profile that should remove the link | yes |

## Example
```json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnlinkChain",
  "target": "cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt",
  "chain_name": "cosmos",
  "owner": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
}
```


## Message action
The action associated to this message is the following:

```
unlink_chain_account
```
