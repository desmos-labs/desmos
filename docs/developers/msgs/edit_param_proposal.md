# `EditParamsProposal`
Inside Desmos there are some parameters set for the `profiles` and `posts` modules which can be changed by submitting a proposal through
the `gov` module of the `cosmosSDK`.

## Structure
```json
{
      "type": "cosmos-sdk/MsgSubmitProposal",
      "value": {
        "content": {
          "type": "cosmos-sdk/ParameterChangeProposal",
          "value": {
            "title": "<Proposal's title>",
            "description": "<Proposal's description>",
            "changes": [
              {
                "subspace": "<Module's subspace>",
                "key": "<Parameter's key>",
                "value": "<Parameter's value>"
              }
            ]
          }
        },
        "initial_deposit": "<Proposal's deposit>",
        "proposer": "<Desmos address of the proposer>"
      }
    }
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `title` | String | Title of the proposal |
| `description` | String | Description of the proposal |
| `subspace` | String | Module's subspace |
| `key` | String | Parameter's store key |
| `value` | JSON | Json representation of parameter's object |
| `initial_deposit` | Object | Proposal's initial deposit |
| `proposer` | String | Desmos address of the proposer |

## Example
```json
{
      "type": "cosmos-sdk/MsgSubmitProposal",
      "value": {
        "content": {
          "type": "cosmos-sdk/ParameterChangeProposal",
          "value": {
            "title": "Moniker Param Change",
            "description": "Update moniker lengths",
            "changes": [
              {
                "subspace": "profiles",
                "key": "monikerParams",
                "value": "{\"type\": \"desmos/MonikerParams\",\"value\": {\"min_moniker_len\":\"5\",\"max_moniker_len\":\"40\"}}"
              }
            ]
          }
        },
        "initial_deposit": [
          {
            "denom": "desmos",
            "amount": "10"
          }
        ],
        "proposer": "desmos19yphj7tdpakp8e55t6y8srk943m0ctf0rc3sqe"
      }
    }
```

## Message action
The action associated to this message is the following: 

```
submit_proposal
```