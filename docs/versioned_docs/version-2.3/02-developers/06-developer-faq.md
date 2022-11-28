---
id: faq
title: F.A.Q
sidebar_position: 6
---

# FAQs

## General Concepts
### What is a transaction? 
As you know, the blockchain can be seen as a decentralized state machine that stores a state. A transaction is the method used by clients to tell the chain that some operations must be taken to change such current state.

Every transaction is composed of some common data (such as the account data of the sender, a `memo`, etc.) but the most important parts are [messages](#what-is-a-message). 

### What is a message?
A (transaction) message is the method that allows you to specify what action(s) should be done inside a transaction to change the current chain state. For example, inside Desmos we can use messages to tell the chain to store a new post, add a like to a post, comment on an existing post and many more.

To know more about all the different messages types and the associated actions that allow you to do, please go to the __transactions__ section. 

### How do I send a transaction?
Sending a transaction is pretty straight forward. All what you need to have is access to an instance of an HD wallet associated with a Desmos account having some `desmos` tokens inside. Once you have it, you simply need to: 

1. Create the proper JSON object containing the message(s) that you want to send as well as the account information of the sender. 

2. Sign that JSON using the private key associated with the HD wallet of the transaction sender. 

3. Put the signed JSON inside a bigger JSON object containing the un-signed transaction data. 

4. Send the complete JSON to a full node REST API endpoint. 

Please note that when sending transactions you are required to pay a **fee** so that the chain can sustain economically. To avoid paying a higher fee and wasting the user's funds, you should always **put multiple messages inside the same transactions**. This will also decrease the overall execution of all messages and can allow you to save time and provide the users a better UX overall. 

### How long does transactions take to be executed? 
Unluckily there is no way to know how long a transaction will take before being executed. The time that passes between it being received by a full node and it's actual execution and verification can vary a lot based on how many messages are inside, how complex each operation to perform is as well as how high the paid fees are. 

If you want you can try speeding up the transactions execution by specifying a higher fee to be paid during the execution itself, but this might now change a lot if other users are doing the same.

Generally, however, transactions take not a very long time and most of the times they get executed in less than 10 seconds from when they are sent to the chain.  

### What's the best way to know when a transaction is performed?
Due to the fact that transactions can take up a different time to be executed (see ["How long does transactions take to be executed"](#how-long-does-transactions-take-to-be-executed)), the best way a client has to stay updated on when a transaction will be executed is by using a [Websocket](https://en.wikipedia.org/wiki/WebSocket). Each and every full node exposes a websocket that can be reached at the following URI: 

```
ws://<full-node-host>/websocket
```

If you want to know more about it, please refer to the [websocket page](05-observe-data.md).

## Developing applications
### I wrongly did an operation. Can I revert it?
Unfortunately, due to the nature of the blockchain itself we cannot allow to revert any operations that have been done. For example, once you send a post to Desmos, it will stay there forever and everyone will be able to read it as it appeared when created. 
Even if you edit a post, the original one will always be there an people will be able to see that you made some changes. It's like trying to editing something that is public and written in a stone that cannot be destroyed. 

For this reason, we suggest you to take **all the possible precautions** before sending any transaction to the chain. 

As an example, you might not want to send a transaction for every post that the user creates, but instead store the locally created posts offline for a certain amount of time (i.e. 2 minutes) and later send the transactions. During this time, the user will still be able to delete the posts if he wants, but once synced on the chain you will no longer be able to take them down.