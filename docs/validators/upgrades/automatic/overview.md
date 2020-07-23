# Automatic upgrade procedure overview
Starting from version `v0.10.0`, Desmos provides the support for automatic chain upgrades based on on-chain proposals due to the implementation of the [`x/upgrade` Cosmos SDK module](https://github.com/cosmos/cosmos-sdk/tree/master/x/upgrade/spec). 

Thanks to this module, validators can have an easier time when a new version of the Desmos binary is available by ensuring their validator node stays always updated without having to manually perform complex tasks.

If you want to know more about setting up the `upgrade_manager` (the binary responsible for handling automatic upgrades), please refer to the [Setup](./setup.md) page. 

If you want to know how proposal-based upgrades will work, please refer to the [Workflow](./workflow.md) page.  
