# `MsgCreateSession`

This message allows you to create a session that links an existing account on an external blockchain to an account on
the Desmos chain.

## Structure

```json
{
   "@type": "/desmos.magpie.v1beta1.MsgCreateSession",
   "owner": "<Desmos address of the account to connect>",
   "namespace": "<Chain id of the external chain>",
   "external_owner": "<Address on the external chain, Bech32 encoded>",
   "pub_key": "<Arbitrary external reference>",
   "signature": "<Signature of the session, Base64 encoded>"
}
```

### Attributes

| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `owner` | String | Desmos address of the user creating the session |
| `namespace` | String | Chain id of the external chain |
| `external_owner` | String | Bech32 encoded address of the external chain account for which to create the session | 
| `pub_key` | String | Public key associated with the external owner used during the signature verification, Base64 encoded |
| `signature` | String | JSON encoded signature data signed with the private key corresponding to the `pub_key` |

### Creating the signature

The signature of the session must be created as follows.

1. Create a JSON object like the following:
   ```json
   {
     "account_number": "0",
     "chain_id": "<namespace>",
     "fee": {
       "gas": "200000",
       "amount": null
     },
     "memo": "",
     "msgs": [<MsgCreateSession with empty signature>],
     "sequence": 0
   }
   ```

2. Sort all the values alphabetically.

3. Sign the JSON with the user private key.

4. Encode the result using the Base64 algorithm.

## Example

```json
{
   "@type": "/desmos.magpie.v1beta1.MsgCreateSession",
   "owner": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
   "namespace": "cosmoshub-3",
   "external_owner": "cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
   "pub_key": "ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
   "signature": "3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg=="
}
``` 

## Message action

The action associated to this message is the following:

```
create_session
``` 