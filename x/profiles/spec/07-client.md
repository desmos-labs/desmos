---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI

A user can query and interact with the `profiles` module using the CLI.

### Query

The `query` commands allow users to query the `profiles` state.

```
desmos query profiles --help
```

#### profile 
The `profile` command allows users to query a profile given an address or a DTag

```bash
desmos query profiles profile [address_or_dtag] [flags]
```

Example: 
```bash
desmos query profiles profile jabbey
```

Example Output: 
```yaml
profile:
  '@type': /desmos.profiles.v2.Profile
  account:
    '@type': /cosmos.auth.v1beta1.BaseAccount
    account_number: "203491"
    address: desmos1xwazl8ftks4gn00y5x3c47auquc62ssu3nt23g
    pub_key:
      '@type': /cosmos.crypto.secp256k1.PubKey
      key: AmyLTnelrZ0zgMx4bFl/n237JKlztLUkhPbHCq6uP/vw
    sequence: "10510"
  bio: just another dad in the cosmos
  creation_date: "2021-11-08T17:51:05.275853979Z"
  dtag: jabbey
  nickname: jabbey
  pictures:
    cover: https://ipfs.infura.io/ipfs/QmTNUuMuu5FBVrFRWdwCNKRJj76oZa65n7ECVaKpQ2aFYN
    profile: https://ipfs.infura.io/ipfs/QmWp7mAnj7UFd13675nEKvWzXKD57SqDA1dbJndV1S4PXD
```

#### incoming-dtag-transfer-requests
The `incoming-dtag-transfer-requests` command allows users to query incoming DTag transfer requests, optionally specifying a recipient user.

```bash
desmos query profiles incoming-dtag-transfer-requests [[receiver]] [flags]
```

Example:
```bash
desmos query profiles incoming-dtag-transfer-requests
```

Example Output:
```yaml
pagination:
  next_key: null
  total: "0"
requests:
- dtag_to_trade: Jack
  receiver: desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu
  sender: desmos1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns
```

#### chain-links
The `chain-links` command allows users to query for chain links optionally specifying a user address, a chain name and a target address.

```bash
desmos query profiles chain-links [[user]] [[chain_name]] [[target]] [flags]
```

**Note**
- The `chain_name` parameter will be used only if the `user` one is specified as well
- The `target` parameter will be used only if the `chain_name` one is specified as well

Example: 
```bash
desmos query profiles chain-links desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd cosmos 
```

Example Output: 
```yaml
links:
- address:
    '@type': /desmos.profiles.v2.Bech32Address
    prefix: cosmos
    value: cosmos1jjectp30f5cp29nudaau94v87tkdxu5n2dvsdy
  chain_config:
    name: cosmos
  creation_time: "2021-10-25T16:00:17.028200092Z"
  proof:
    plain_text: 7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a22636f736d6f736875622d34222c22666565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2230222c2264656e6f6d223a2261746f6d227d5d2c22676173223a2231227d2c226d656d6f223a22436861696e204c696e6b2050726f6f662c20444f4e2754204544495420414e595448494e47222c226d736773223a5b5d2c2273657175656e6365223a2230227d
    pub_key:
      '@type': /cosmos.crypto.secp256k1.PubKey
      key: AmSl9tkQOEkT2LvcdReB/tHS1JQESJ6NeFkDEujwWcQz
    signature:
      '@type': /desmos.profiles.v2.SingleSignatureData
      mode: SIGN_MODE_LEGACY_AMINO_JSON
      signature: rt0E4vQpX/gDI4I8OFykuJCsYyuPlVVUqSpFsCzU8FcbG02kDcDQ+AipVvEBmV1LrDV0/U23Jwi6L8AnMdo1Zw==
  user: desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd
pagination:
  next_key: null
  total: "0"
```

#### default-external-addresses
The `default-external-addresses` command allows users to query for default external addresses optionally specifying an owner address and a chain name.

```bash
desmos query profiles default-external-addresses [[owner]] [[chain_name]] [flags]
```

**Notes**
- The `chain_name` paremeter will only be used if the `owner` is specified

Example:
```bash
desmos query profiles default-external-address desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c cosmos
```

Example Output:
```yaml
links:
- address:
    '@type': /desmos.profiles.v3.Bech32Address
    prefix: cosmos
    value: cosmos13n9wek2ktpxhpgfrd39zlaqaeahxuyusxrsfvn
  chain_config:
    name: cosmos
  creation_time: "2022-07-18T10:07:51.899288600Z"
  proof:
    plain_text: 6465736d6f733165766a323072796d66747665636d676e3874307876373030776b6a6c67756377667934663063
    pub_key:
      '@type': /cosmos.crypto.secp256k1.PubKey
      key: AqYZhHKaeBcrYktZEvor/SUDlHCkv5JBplaG2vc2bvfS
    signature:
      '@type': /desmos.profiles.v3.SingleSignatureData
      mode: SIGN_MODE_DIRECT
      signature: 3yCU8/HKv7Vn0sf7HB+AhV/hK37DCBAkQkXdruaFSvMjRZB1XrYkpKWZVi+xnhSenc1p951Q1058rrYjuNCk9g==
  user: desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c
pagination:
  next_key: null
  total: "0"
```

#### app-links
The `app-links` allows users to query application links optionally specifying a user, an application name and a username.

```bash
desmos query profiles app-links [[user]] [[application]] [[username]] [flags]
```

**Notes**
- The `application` parameter will only be used if the `user` one is specified
- The `username` parameter will only be used if the `application` one is specified

Example: 
```bash
desmos query profiles app-links desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd twitter
```

Example Output: 
```yaml
links:
- creation_time: "2021-11-04T10:50:54.794047348Z"
  data:
    application: twitter
    username: lucagraziotti
  oracle_request:
    call_data:
      application: twitter
      call_data: 7b226d6574686f64223a227477656574222c2276616c7565223a2231343536323036353331383830343233343330227d
    client_id: desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd-twitter-lucagraziotti
    id: "537363"
    oracle_script_id: "32"
  result:
    success:
      signature: 3558f0a50c68d48f66f808f7d12f90be07ffecf60bbdb67e514b3b00c0d30322644168a8369d991df65ff2f8e4890ab2ef5677a33a1304a1c1dff8e627296dde
      value: 6c7563616772617a696f747469
  state: APPLICATION_LINK_STATE_VERIFICATION_SUCCESS
  user: desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd
pagination:
  next_key: null
  total: "0"
```

#### parameters
The `parameters` command allows users to get hte currently set parameters. 

```bash
desmos query profiles parameters [flags]
```

Example:
```bash
desmos query profiles parameters
```

Example Output:
```yaml
params:
  bio:
    max_length: "1000"
  dtag:
    max_length: "30"
    min_length: "3"
    reg_ex: ^[A-Za-z0-9_]+$
  nickname:
    max_length: "1000"
    min_length: "2"
  oracle:
    ask_count: "5"
    execute_gas: "200000"
    fee_amount: []
    min_count: "3"
    prepare_gas: "50000"
    script_id: "32"
```

### Transactions 
The `tx` commands allow users to interact with the `profiles` module. 

```bash
desmos tx profiles --help
```

#### save
The `save` command allows users to save their Desmos profile. 

```bash
desmos tx profiles save [flags]
```

Example:
```bash
desmos tx profiles save 
 	--dtag "JohnSnow" \
 	--nickname "John Snow" \
 	--bio "Son of Lyanna Stark and Rhaegar Targaryen" \
 	--profile-pic "https://profilePic.jpg" \
 	--cover-pic "https://profileCover.jpg"
```

#### delete
The `delete` command allows users to delete their Desmos profile. 

```bash
desmos tx profiles delete [flags]
```

Example:
```bash
desmos tx profiles delete
```

#### request-dtag-transfer
The `request-dtag-transfer` allows users to request the ownership transferring of a DTag from another user.

```bash
desmos tx profiles request-dtag-transfer [address] [flags]
```

Example:
```bash
desmos tx profiles request-dtag-transfer desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

#### cancel-dtag-transfer-request
The `cancel-dtag-transfer-request` allows users to cancel the transferring of a DTag ownership made towards the specified user.

```bash
desmos tx profiles cancel-dtag-transfer-request [recipient] [flags]
```

Example: 
```bash
desmos tx profiles cancel-dtag-transfer-request desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

#### accept-dtag-transfer-request
The `accept-dtag-transfer-request` allows users to accept an incoming DTag transfer request specifying the new DTag that they will have after the ownership transferring is completed. 

```bash
desmos tx profiles accept-dtag-transfer-request [DTag] [address] [flags]
```

Example:
```bash
desmos tx profiles accept-dtag-transfer-request "NewJohnSnow" desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

#### refuse-dtag-transfer-request
The `refuse-dtag-transfer-request` allows users to refuse an incoming DTag transfer request. 

```bash
desmos tx profiles refuse-dtag-transfer-request [sender] [flags]
```

Example:
```bash
desmos tx profiles refuse-dtag-transfer-request desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

#### link-chain 
The `link-chain` command allows users to link a new chain account to their Desmos profile. 

```bash
desmos tx profiles link-chain [data-file] [flags]
```

Example:
```bash
desmos tx profiles link-chain chain-link.json
```

#### unlink-chain
The `unlink-chain` allows users to unlink a previously linked chain account. 

```bash
unlink-chain [chain-name] [address]
```

Example: 
```bash
desmos tx profiles unlink-chain "cosmos" cosmos18xnmlzqrqr6zt526pnczxe65zk3f4xgmndpxn2`
```

#### set-default-external-address
The `set-default-external-address` command allows user to set a default external address.

```bash
set-default-external-address [chain-name] [target]
```

Example:
```bash
desmos tx profiles set-default-external-address "cosmos" cosmos18xnmlzqrqr6zt526pnczxe65zk3f4xgmndpxn2
```

#### link-app
The `link-app` command allows users to link an external app account to their Desmos profile. 

```bash
desmos tx profiles link-app [src-port] [src-channel] [application] [username] [verification-call-data] [flags]
```

Example: 
```bash
desmos tx profiles link-app "profiles" "channel-0" "twitter" "twitter_user" "7B22757365726E616D65223A22526963636172"
```

#### unlink-app
The `unlink-app` command allows users to unlink a previously linked application account.

```bash
desmos tx profiles unlink-app [application] [username] [flags]
```

Example:
```bash
desmos tx profiles unlink-app "twitter" "twitter_user"
```

## gRPC 
A user can query the `profiles` module gRPC endpoints. 

### Profile 
The `Profile` endpoint allows users to query for a profile based on the given address or DTag. 

```bash
desmos.profiles.v2.Query/Profile
```

Example:
```bash
grpcurl -plaintext \
  -d '{"user": "jack"}' localhost:9090 desmos.profiles.v2.Query/Profile
```

Example Output: 
```json
{
  "profile": {
    "@type": "/desmos.profiles.v2.Profile",
    "account": {
      "@type": "/cosmos.auth.v1beta1.BaseAccount",
      "address": "desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu",
      "pubKey": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": "AzAk6eCtT5MEXvWmC7khceZBjNE7CC56e8PtBTEqC0F9"
      },
      "sequence": "3"
    },
    "creationDate": "2022-04-21T08:22:58.743790408Z",
    "dtag": "Jack",
    "pictures": {}
  }
}
```

### IncomingDTagTransferRequests
The `IncomingDTagTransferRequests` endpoint allows users to query for incoming DTag transfer requests. 

```bash
desmos.profiles.v2.Query/IncomingDTagTransferRequests
```

Example: 
```bash
grpcurl -plaintext \
  -d '{"receiver": "desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu"}' localhost:9090 desmos.profiles.v2.Query/IncomingDTagTransferRequests
```

Example Output: 
```json
{
  "requests": [
    {
      "dtagToTrade": "Jack",
      "sender": "desmos1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns",
      "receiver": "desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### ChainLinks
The `ChainLinks` endpoint allows users to query for chain links specifying an optional user, chain name and target. 

```bash
desmos.profiles.v2.Query/ChainLinks
```

**Note**
- The `chain_name` parameter will be used only if the `user` one is specified as well
- The `target` parameter will be used only if the `chain_name` one is specified as well


Example:
```bash
grpcurl -plaintext \
  -d '{"user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd", "chain_name": "osmosis"}' localhost:9090 desmos.profiles.v2.Query/ChainLinks
```

Example Output:
```json
{
  "links": [
    {
      "user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd",
      "address": {
        "@type": "/desmos.profiles.v2.Bech32Address",
        "prefix": "osmo",
        "value": "osmo1kdl9888p5ez2mt7w59anugx7z0ek927gvdvk7p"
      },
      "proof": {
        "pubKey": {
          "@type": "/cosmos.crypto.secp256k1.PubKey",
          "key": "Av/eQ32q7EerWFX7CIxxguthkmlShmR+LIeOmIE+4Qm0"
        },
        "signature": {
          "@type": "/desmos.profiles.v2.SingleSignatureData",
          "mode": "SIGN_MODE_TEXTUAL",
          "signature": "b8EYBePyOGDdLautzdiEXj2CR/0gpWyHMwoQfUizrG4q6qTX2ZivWf/NiKsmx9h2YHZq4OOxIwZV/vgf8J7ZLA=="
        },
        "plainText": "6465736d6f7331366336307938743876726132377a6a673261726c6364353864636b3963776e37703666777464"
      },
      "chainConfig": {
        "name": "osmosis"
      },
      "creationTime": "2021-11-18T13:56:12.238062857Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### ChainLinkOwners
The `ChainLinkOwners` endpoint allows users to query for chain link owners given an optional chain name and target address.

```bash
desmos.profiles.v2.Query/ChainLinkOwners
```

**Note** 
The `target` parameter will be used only if the `chain_name` is specified as well.

Example:
```bash
grpcurl -plaintext \
  -d '{"chain_name": "osmosis", "target": "osmo1kdl9888p5ez2mt7w59anugx7z0ek927gvdvk7p"}' localhost:9090 desmos.profiles.v2.Query/ChainLinkOwners
```

Example Output: 
```json
{
  "owners": [
    {
      "user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd",
      "chain_name": "osmosis",
      "target": "osmo1kdl9888p5ez2mt7w59anugx7z0ek927gvdvk7p"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### DefaultExternalAddresses
The `DefaultExternalAddresses` endpoint allows to query for external default addresses optionally specifying an owner and chain name.

```bash
desmos.profiles.v2.Query/desmos.profiles.v3.Query/DefaultExternalAddresses
```

**Notes**
- The `chainName` parameter will only be used if the `owner` one is specified

Example
```bash
grpcurl -plaintext  \
  -d '{"owner": "desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c", "chainName": "cosmos"}' localhost:9090 desmos.profiles.v3.Query/DefaultExternalAddresses
```

Example Output:
```json
{
  "links": [
    {
      "user": "desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c",
      "address": {"@type":"/desmos.profiles.v3.Bech32Address","prefix":"cosmos","value":"cosmos13n9wek2ktpxhpgfrd39zlaqaeahxuyusxrsfvn"},
      "proof": {
        "pubKey": {"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AqYZhHKaeBcrYktZEvor/SUDlHCkv5JBplaG2vc2bvfS"},
        "signature": {"@type":"/desmos.profiles.v3.SingleSignatureData","mode":"SIGN_MODE_DIRECT","signature":"3yCU8/HKv7Vn0sf7HB+AhV/hK37DCBAkQkXdruaFSvMjRZB1XrYkpKWZVi+xnhSenc1p951Q1058rrYjuNCk9g=="},
        "plainText": "6465736d6f733165766a323072796d66747665636d676e3874307876373030776b6a6c67756377667934663063"
      },
      "chainConfig": {
        "name": "cosmos"
      },
      "creationTime": "2022-07-18T10:07:51.899288600Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### ApplicationLinks
The `ApplicationLinks` endpoint allows users to query for application links optionally specifying a user, application name and username. 

```bash
desmos.profiles.v2.Query/ApplicationLinks
```

**Notes**
- The `application` parameter will only be used if the `user` one is specified
- The `username` parameter will only be used if the `application` one is specified

Example:
```bash
grpcurl -plaintext \
  -d '{"user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd", "application": "twitter"}' localhost:9090 desmos.profiles.v2.Query/ApplicationLinks
```

Example Output: 
```json
{
  "links": [
    {
      "user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd",
      "data": {
        "application": "twitter",
        "username": "lucagraziotti"
      },
      "state": "APPLICATION_LINK_STATE_VERIFICATION_SUCCESS",
      "oracleRequest": {
        "id": "537363",
        "oracleScriptId": "32",
        "callData": {
          "application": "twitter",
          "callData": "7b226d6574686f64223a227477656574222c2276616c7565223a2231343536323036353331383830343233343330227d"
        },
        "clientId": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd-twitter-lucagraziotti"
      },
      "result": {
        "success": {
          "value": "6c7563616772617a696f747469",
          "signature": "3558f0a50c68d48f66f808f7d12f90be07ffecf60bbdb67e514b3b00c0d30322644168a8369d991df65ff2f8e4890ab2ef5677a33a1304a1c1dff8e627296dde"
        }
      },
      "creationTime": "2021-11-04T10:50:54.794047348Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### ApplicationLinkByClientID
The `ApplicationLinkByClientID` endpoint allows users to query an application link given a client id. 

```bash
desmos.profiles.v2.Query/ApplicationLinkByClientID
```

Example:
```bash
grpcurl -plaintext \
  -d '{"client_id": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd-twitter-lucagraziotti"}' localhost:9090 desmos.profiles.v2.Query/ApplicationLinkByClientID
```

Example Output: 
```json
{
  "link": {
    "user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd",
    "data": {
      "application": "twitter",
      "username": "lucagraziotti"
    },
    "state": "APPLICATION_LINK_STATE_VERIFICATION_SUCCESS",
    "oracleRequest": {
      "id": "537363",
      "oracleScriptId": "32",
      "callData": {
        "application": "twitter",
        "callData": "7b226d6574686f64223a227477656574222c2276616c7565223a2231343536323036353331383830343233343330227d"
      },
      "clientId": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd-twitter-lucagraziotti"
    },
    "result": {
      "success": {
        "value": "6c7563616772617a696f747469",
        "signature": "3558f0a50c68d48f66f808f7d12f90be07ffecf60bbdb67e514b3b00c0d30322644168a8369d991df65ff2f8e4890ab2ef5677a33a1304a1c1dff8e627296dde"
      }
    },
    "creationTime": "2021-11-04T10:50:54.794047348Z"
  }
}
```

### ApplicationLinkOwners
The `ApplicationLinkOwners` endpoint allows users to query the application link owners given an optional application name and username. 

```bash
desmos.profiles.v2.Query/ApplicationLinkOwners
```

**Note**
The `user` parameter will be used only if the `application` one is specified as well. 

Example:
```bash
grpcurl -plaintext \
  -d '{"application": "twitter", "user": "lucagraziotti"}' localhost:9090 desmos.profiles.v2.Query/ApplicationLinkOwners
```

Example Output:
```json
{
  "owners": [
    {
      "user": "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd",
      "application": "twitter",
      "username": "lucagraziotti"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### Params
The `Params` endpoint allows users to query the current params of the `profiles` module. 

```bash
desmos.profiles.v2.Query/Params
```

Example:
```bash
grpcurl localhost:9090 desmos.profiles.v2.Query/Params
```

Example Output:
```json
{
  "params": {
    "nickname": {
      "minLength": "Mg==",
      "maxLength": "MTAwMA=="
    },
    "dtag": {
      "regEx": "^[A-Za-z0-9_]+$",
      "minLength": "Mw==",
      "maxLength": "MzA="
    },
    "bio": {
      "maxLength": "MTAwMA=="
    },
    "oracle": {
      "scriptId": "32",
      "askCount": "5",
      "minCount": "3",
      "prepareGas": "50000",
      "executeGas": "200000"
    }
  }
}
```

## REST
A user can query the `profiles` module using REST endpoints. 

### Profile
The `profile` endpoint allows users to query for a Desmos profile using a DTag or an address. 

```
/desmos/profiles/v2/profiles/{DTag or address}
```

### Incoming DTag Transfer Requests
The `dtag-transfer-requests` endpoint allows users to query for incoming DTag transfer requests given an optional user address.

```
/desmos/profiles/v2/dtag-transfer-requests?receiver={address}
```

### Chain Links
The `chain-links` endpoint allows users to query for chain links given an optional user, chain name and target. 

```
/desmos/profiles/v2/chain-links?user={user}&chain_name={chainName}&target={target}
```

**Note**
- The `chain_name` parameter will be used only if the `user` one is specified as well
- The `target` parameter will be used only if the `chain_name` one is specified as well 

### Chain Links Owners
The `chain-links/owners` endpoint allows users to query for chain link owners given an optional chain name and target. 

```
/desmos/profiles/v2/chain-links/owners?chain_name={chainName}&target={target}
```

**Note**
The `target` parameter will be used only if the `chain_name` is specified as well.

### DefaultExternalAddresses
The `default-external-addresses` endpoint allows users to query for default external addresses given an optional owner address and chain name.

```
/desmos/profiles/v3/default-external-addresses?owner={owner}&chain_name={chainName}
```

**Note**
- The `chain_name` parameter will be used only if the `owner` is specified as well.

### Application Links
The `app-links` endpoint allows users to query for application links given an optional user, application name and username. 

```
/desmos/profiles/v2/app-links?user={user}&application={application}&username={username}
```

**Notes**
- The `application` parameter will only be used if the `user` one is specified
- The `username` parameter will only be used if the `application` one is specified  

### Application Links By Client ID
The `app-links/clients` endpoint allows users to get application links given a client id. 

```
/desmos/profiles/v2/app-links/clients/{client_id}
```

### Application Links Owners
The `app-links/owners` endpoint allows users to query for application link owners specifying an optional application and username. 

```
/desmos/profiles/v2/app-links/owners?application={applicationName}&username={username}
```

**Note**
The `user` parameter will be used only if the `application` one is specified as well.

### Params 
The `params` endpoint allows users to query for the module parameters. 

```
/desmos/profiles/v2/params
```