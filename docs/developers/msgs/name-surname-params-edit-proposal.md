# `EditNameSurnameParamsProposal`
This proposal allows you to request a change of the profiles' name and surname parameters.

## Structure
````json
{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "content": {
      "type": "desmos/EditNameSurnameParamsProposal",
      "value": {
        "title": "<Proposal title>",
        "description": "<Proposal description>",
        "name_surname_params": "<Name/Surname parameters lengths>"
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
| `name_surname_params` | Object | Profile's parameters to be changed |
| `initial_deposit` | Object | Proposal's initial deposit |
| `proposer` | String | Desmos address of the user that is creating the proposal |

## Example
````json
{
  "type": "cosmos-sdk/MsgSubmitProposal",
  "value": {
    "content": {
      "type": "desmos/NameSurnameParamsEditProposal",
      "value": {
        "title": "Edit name and surname params proposal",
        "description": "My awesome proposal",
        "name_surname_params": {
          "min_name_surname_len": "3",
          "max_name_surname_len": "200"
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