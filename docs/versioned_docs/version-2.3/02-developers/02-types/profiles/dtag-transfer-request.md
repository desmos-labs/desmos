---
id: dtag-transfer-request
title: DTag transfer request
sidebar_label: DTag transfer request
slug: dtag-transfer-request
---

# DTag transfer request
A DTag transfer request represents the request made from a user to get the DTag of another one.
 
## Contained data
Here follows the data of a DTag transfer request. 

### `DTagToTrade` (`string`)
The `DTag` contains the value of the `DTag` that should be transferred from the receiver of the request to the sender.

### `Sender` (`string`)
Sender represents the address of the account that sent the `DTag transfer request`.

### `Receiver` (`string`)
Receiver represents the receiver of the request that, if accepted, will
give to the sender their `DTag`.

