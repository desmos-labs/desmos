---
id: observe-data
title: Observing data
sidebar_label: Observing data
slug: observe-data
---

# Observing new data

## Introduction
Aside from querying data, you can also observe new data as its inserted inside the chain itself. In this way, you will be notified as soon as a transaction is properly executed without having to constantly polling the chain state by yourself. 

## Websocket  
All the live data observation is done though the usage of a [websocket](https://en.wikipedia.org/wiki/WebSocket). The endpoint of such websocket is the following: 

```
ws://lcd-endpoint/websocket

# Example
# ws://morpheus.desmos.network/websocket
```

### Events
In order to subscribe to specific events, you will need to send one or more messages to the websocket once you opened a connection to it. Such messages need to contain the following JSON object and must be string encoded: 

```json
{
  "jsonrpc": "2.0",
  "method": "subscribe",
  "id": "0",
  "params": {
    "query": "tm.event='eventCategory' AND eventType.eventAttribute='attributeValue'"
  }
}
``` 

The `query` field can have the following values: 

* `tm.event='NewBlock'` if you want to observe each new block that is created (even empty ones);
* `tm.event='Tx'` if you want to subscribe to all new transactions;
* `message.action='<action>'` if you want to subscribe to events emitted when a specific message is sent to the chain. 
  In this case, please refer to the `Message action` section on each transaction message 
  specification page to know what is the type associated to each message.

Please note that if you want to subscribe to multiple events you will need to send multiple query messages upon connecting to the websocket. 

#### Example
```json
{
  "jsonrpc": "2.0",
  "method": "subscribe",
  "id": "0",
  "params": {
    "query": "message.action='save_profile'"
  }
}
```