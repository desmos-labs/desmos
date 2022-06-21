---
id: parameters
title: Parameters
sidebar_label: Parameters
slug: parameters
---

# Parameters

The profiles module contains the following parameters: 

| Key            | Type           | Example                                                                                                                  |
|----------------|----------------|--------------------------------------------------------------------------------------------------------------------------|
| NicknameParams | NicknameParams | `{ "min_length": "2", "max_length": "1000" }`                                                                            |
| DTagParams     | DTagParams     | `{ "reg_ex": "^[A-Za-z0-9_]+$", "min_length": "3", "max_length": "30" }`                                                 |
| MaxBioLen      | BioParams      | `{ "max_length": "1000" }`                                                                                               |
| OracleParams   | OracleParams   | `{ "script_id":"32", "ask_count":"5", "min_count":"3", "prepare_gas":"50000", "execute_gas":"200000", "fee_amount":[] }` |

## OracleParams
The oracle parameter contains the details about the Band Protocol oracle script that needs to be called when verifying application links.