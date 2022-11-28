---
id: application-link
title: Application link
sidebar_label: Application link
slug: application-link
---

# Application link
An application link (abbr. _app link_) represents a link to an external (and possibly centralized) application.

## Contained data
Here follows the data of an application link.

### `User` (`string`)
Address of the Desmos profile to which the link is associated. 

### `Data` (`object`)
It contains the link details:

- #### `Application` (`string`)
  Name of the application to which the link refers to (eg. `twitter`, `github`, `reddit`, etc). 

- #### `Username` (`string`)
  Identifier of the application account which the link refers to (eg. Twitter username, GitHub profile, Reddit username, etc).

### `State` (`object`)
Representation of the current state of the link.  
There can be five different states in which a link can be: 

- `APPLICATION_LINK_STATE_INITIALIZED_UNSPECIFIED` if the link has just been created, and it still needs to be processed; 
- `APPLICATION_LINK_STATE_VERIFICATION_STARTED` if the verification process has started; 
- `APPLICATION_LINK_STATE_VERIFICATION_ERROR` if the verification process ended with an error; 
- `APPLICATION_LINK_STATE_VERIFICATION_SUCCESS` if the verification process ended with success;
- `APPLICATION_LINK_STATE_TIMED_OUT` if the verification process expired due to a timeout. 

### `OracleRequest` (`object`)
The `OracleRequest` field contains all the data that has been sent to the oracle script in order to verify the authenticity of the link. 

- #### `ID` (`int64`)
  This is the unique id of the request that has been made to verify the link. 

- #### `OracleScriptID` (`int64`)
  A unique id of the script that has been called to verify the authenticity of the link. 

- #### `CallData` (`object`)
  Contains the details of the data that will be used to call the oracle script. 

  - ##### `Application` (`string`)
    Name of the application for which the link is valid (eg. `twitter`, `github`, `reddit`, etc). 

  - ##### `CallData` (`string`)
    The `CallData` field represents the hex-encoded data that will be given to the data source in order to fetch and verify the validity of the link. 

- #### `ClientID` (`string`)
    ID of the client that has performed the request.

### `Result` (`object`)
The `Result` field contains the effective result of the verification process. This is set only if the link state is either `APPLICATION_LINK_STATE_VERIFICATION_SUCCESS` or `APPLICATION_LINK_STATE_VERIFICATION_ERROR`. 

The `Result` field can be of two types:
- #### `Result_Success` (`object`)
  Represents a successful result. It contains two fields: 
      
  - ##### `Value` (`string`)
    Plain text value that has been signed from the user with their Desmos private key to prove the ownership of the Desmos profile. 

  - ##### `Signature` (`string`)
    Hex-encoded result of the plain text value signature. 
    
- #### `Result_Failed` (`object`)
  Identifies an error during the verification process. It contains only one field:

  - ##### `Error` (`string`)
    Represents the description of the error that has been emitted during the verification process.

### `CreationTime` (`time`)
Contains the time at which the link has been created. 

## Create an application link
Application links can be created by any user having a Desmos profile, and their validity is checked using a multi-step verification process described inside the [_"Themis"_ repository](https://github.com/desmos-labs/themis). 
