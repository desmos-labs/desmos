# Desmos "Apollo" testnet

## Overview

In order to make sure that validators know what to do when the mainnet will start, we decided to run a copycat chain
called _"Apollo"_. This will be used to test the final Desmos codebase as well as to test the coordination between
different validators and future updates. This testnet will start the same way as the mainnet, using a decentralized
procedure made of three different parts:

1. Creation of genesis transactions by validators.
2. Collection of all genesis transactions.
3. Chain start.

All validators that would like to take part into this testnet from its beginning should participate in this procedure.
Please note that you can follow this procedure only after you
have [created a full node](../fullnode/setup.md#full-node-setup).

:::tip You can join later  
If you do not have time to join our testnet from its beginning you can always join later.  
If you wish to know how please read [here](../testnets/join-public.md).
:::

:::warning Rewards  
This testnet is **not** meant to reward the validators that will take part to it directly. This means that we will
**not** give any DSM to people that will run a validator node. However, we will use this testnet a playground to observe
how different validators behave and reward them via a delegation inside our mainnet if they decide to run a node there
as well.  
:::

## Join "Apollo"

If you are a validator that would like to take part in the genesis of the "Apollo" testnet, all you have to do is:

1. Download the current genesis file.
2. Create a genesis transaction.
3. Submit that genesis transaction to our GitHub repository.

### 1. Download the current genesis file

In order to perform a genesis transaction, you first have to download the current genesis file for the "Apollo"
testnet. To do so, you have to:

1. initialize your node;
2. download the genesis file.

#### Initialize your node

In order to initialize your node appropriately, you can use the `desmos init` command:

```shell
desmos init <Your moniker>
```

:::tip Generate a deterministic node  
We suggest you using the `--recover` flag while running the `desmos init` command. This will ask you for a seed phrase
that will be used to generate the node private key. Doing so will allow you to recover your node private key if
something goes wrong, instead of having to backup the `priv_key.json` file by hand.  
:::

#### Download the genesis file

Once you have initialized your node, it is now time to download the current genesis file for the "Apollo" testnet. To do
so, you can run the following command:

```shell
rm ~/.desmos/config/genesis.json
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/morpheus-apollo-1/genesis.json > ~/.desmos/config/genesis.json
```

### 2. Creating a genesis transaction

Genesis transactions are used to create the validators nodes that will need to be present in order to start the chain.
This is done by using the `desmos gentx` command that allows you to specify the details of your node.

In order to use this command, however, you need to first have an account inside the genesis file itself. To create such
account you can use the `desmos add-genesis-account` command:

```bash
desmos add-genesis-account <your_key> 10000000udaric 
```

:::warning Amount value  
Please do not try giving yourself more than `10000000udaric` (equivalent to `10 Daric`). If you do so, your genesis
transaction submission will not be accepted, and you will not be able to take part in the genesis procedure.  
:::

After you have inserted your account inside the genesis file, you are now ready to create your validator node using a
genesis transaction:

```bash
desmos gentx <your_key> 1000000udaric --chain-id=morpheus-apollo-1 \
    --moniker="<Your validator moniker>" \
    --commission-max-change-rate=0.01 \
    --commission-max-rate=1.0 \
    --commission-rate=0.10 \
    --details="..." \
    --security-contact="..." \
    --website="..."
```

Once you run this command, the output is going to look something like this:

```
Genesis transaction written to "/home/user/.desmos/config/gentx/gentx-6b1fe44615aa1ac9b0dfc637d1a33fd63de2a05e.json"
```

We're going to refer to that file location as `/path/to/gentx.json`.

### 3. Submit your genesis transaction

Now that everything is ready, the last step you have to perform is to commit the changed genesis file and the new
genesis transaction to our repo.

First thing first, you need to [fork](https://docs.github.com/en/github/getting-started-with-github/fork-a-repo) our
[testnet repository](https://github.com/desmos-labs/morpheus). Once you have done so, you can clone your fork with the
following command:

```shell
git clone https://github.com/<Your GitHub username>/morpheus.git 
```

Now, copy the updated genesis file and the new genesis transaction file inside the proper paths:

```shell
cd morpheus
cp ~/.desmos/config/genesis.json morpheus-apollo-1/ 
cp /path/to/gentx.json moprheus-apollo-1/
```

Next, add, commit and push the new files to your fork repository:

```shell
git add . 
git commit -m "Your commit message"
git push
```

Finally, [create a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
so that your changes can be merged inside our testnet repository.

:::warning Make sure you pass the automatic checks  
When you create a pull request inside our repository, we will perform some automatic checks. These include but are not
limited to:

- validating the genesis account amount
- validating the genesis transaction signature

If you do not pass all of these checks, we will not merge your pull request. This means you will not be able to have
your node inside the final genesis file that will be used to start the chain.  
:::

### Final remarks

Congratulations, you successfully joined the Morpheus "Apollo" testnet as a genesis validator. All you have to do now is
wait until the final genesis file is created. Once that's done, you will need to download it inside your node and then
start the `desmos` process.

We highly suggest you to join [our Discord server](https://discord.gg/yxPRGdq) so that you can keep updated with all the
new information and coordinate with other validator nodes properly.