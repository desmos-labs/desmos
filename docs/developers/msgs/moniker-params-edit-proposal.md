# `EditMonikerParamsProposal`
This proposal allows you to request a change of the profiles' moniker's parameters.

## Structure
````json
{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "content": {
      "type": "desmos/EditMonikerParamsProposal",
      "value": {
        "title": "<Proposal title>",
        "description": "<Proposal description>",
        "moniker_params": "<Moniker's parameters lengths>"
      }
    },
    "initial_deposit": ["<Proposal's deposit>"],
    "proposer": "<Desmos address of the proposal's proposer>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `title` | String | Title of the proposal |
| `description` | String | Proposal's description |
| `moniker_params` | Object | Profile's parameters to be changed |
| `initial_deposit` | Object | Proposal's initial deposit |
| `proposer` | String | Desmos address of the user that is creating the proposal |

## Example
````json
{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "content": {
      "type": "desmos/EditMonikerParamsProposal",
      "value": {
        "title": "Edit moniker's params proposal",
        "description": "My awesome proposal",
        "moniker_params": {
          "min_moniker_len": "5",
          "max_moniker_len": "30"
        }
      }
    },
    "initial_deposit": [
      {
        "denom": "desmos",
        "amount": "10"
      }
    ],
    "proposer": "desmos1kmw3pp2825hs3mfca2y9txz72638auqm68ngma"
  }
}
````

## Message action
The action associated to this message is the following:

```
submit_proposal
```