---
id: full-sync
title: Full Sync
sidebar_label: Full Sync
slug: full-sync
---
# Mainnet Full sync
## Software downgrade 

:::note   
You will need to build the first version of the Desmos mainnet in order to be able to sync the chain from scratch.
:::

```bash
# Make sure we are inside the home directory
cd $HOME

# Clone the Desmos software
git clone https://github.com/desmos-labs/desmos.git && cd desmos

# Checkout the correct tag
git checkout tags/v1.0.3

# Build the software
# If you want to use the default database backend run
make install
```
