# Upgrade procedure workflow
As seen inside the [overview](./overview.md), Desmos now supports automatic proposal-based upgrades. Following you will find the description of the upgrading procedure that we will adopt from now on inside our chain. 

## Upgrade proposals definition
In order to make sure each validator upgrades its node in the correct moment in time, upgrades wil be proposed directly on-chain through the usage of **upgrade proposals**.

Upgrade proposals are on-chain proposals that can be created by anyone from within the chain itself using the [`x/gov` module from the Cosmos SDK](https://github.com/cosmos/cosmos-sdk/tree/master/x/gov/spec). Particularly, this kind of proposals contain an upgrade **plan**.

An upgrade plan contains the following data: 

```go
// Plan specifies information about a planned upgrade and when it should occur
type Plan struct {
	// Sets the name for the upgrade. This name will be used by the upgraded version of the software to apply any
	// special "on-upgrade" commands during the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no upgrade handler with this name has been
	// set in the software, it will be assumed that the software is out-of-date when the upgrade Time or Height
	// is reached and the software will exit.
	Name string `json:"name,omitempty"`

	// The time after which the upgrade must be performed.
	// Leave set to its zero value to use a pre-defined Height instead.
	Time time.Time `json:"time,omitempty"`

	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `json:"height,omitempty"`

	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `json:"info,omitempty"`
}
```

In order to create a simple upgrade proposal, the command that can be execute is the following: 

```shell
desmos tx gov submit-proposal software-upgrade "<Upgrade name>" \
  --title "<Proposal title>" \
  --description "<Proposal description>" \
  --deposit <Initial deposit> \
  --upgrade-height=<Upgrade height>
```

While creating an upgrade proposal, the following data are **required**: 

- the `title` of the proposal
- the `description` of the proposal
- the `upgrade-height` **or** `upgrade-time`

Optionally, an initial `deposit` can be specified. 

## Upgrade proposal life cycle
Once that an upgrade proposal has been created, it needs to go through a series of steps before being accepted and finally trigger the update. 

The first step it needs to go through is the **deposit** period. During this time, users can deposit their tokens into the proposal to make sure it will get into the next phase (voting period). If the proposal gathers enough tokens (defined by the `x/gov` module params), then the deposited tokens will be returned to the depositor. Otherwise, if not enough tokens are deposited before the deposit time ends, all the deposited tokens will be burned.  

Once that enough tokens have been deposited into the proposal, it then goes into the **voting period**.  This period has a predefined length, which is established by the `x/gov` params. Durign this period, any chain user can vote either "Yes", "No" or "No with veto" to the proposal. In order to pass, more than 75% of voting power need to vote into the proposal, and the 51% of votings must be "Yes". A 33% of "No with veto" will invalidate the proposal. 

If the proposal passes, then the upgrade plan is put into action. 
 
## Validators responsibility
As a validator, if a proposal passes you will need to perform the following operations: 

First of all, you need to checkout the proper Desmos version. The upgrade proposal will contain the details of what Desmos version should be used during the upgrade process. 
   
In order to make sure the `upgrade_manager` utility properly handles the upgrade, you need then to build the Desmos binary and put inside the correct folder on your machine. In order to do so run:

```shell
cd ~/desmos
make build
mkdir -p ~/.desmos/upgrade_manager/upgrades/<upgrade-name>/bin
cp build/desmos ~/.desmos/upgrade_manager/upgrades/<upgrade-name>/bin
``` 

Please note that the `<upgrade-name>` placeholder should be replaced with the name of the upgrade that is put inside the upgrade proposal. 

### Scripts
In order to make validators lives easier, during our testnet phase we will provide scripts that you can run in order to perform these tasks easily. We will also provide `.tar.gz` files that can be downloaded and contain all the necessary data to get though upgrades.

### Automatic binaries download
When creating upgrade proposals, from time to time we will also try to specify the binaries that should be automatically downloaded during the upgrade. This will be done following the [`cosmosd` specification](https://github.com/regen-network/cosmosd#auto-download). 

If you want, you can enable the `DAEMON_ALLOW_DOWNLOAD_BINARIES` environmental variable during the [setup](./setup.md) in order to allow your node to auto download them and perform all the procedure by itself. 

However, please keep in mind that while this is fine to do in a test environment (i.e. the during the testnet phase), it is less OK to be done inside a production environment. You should always be caution with automatic operations and always be ready to manually override them if needed.
