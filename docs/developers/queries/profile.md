# Query a profile
This query endpoint allows you to retrieve the details of a single profile having its moniker or address. 

**CLI**
 ```bash
desmoscli query profiles profile [address_or_moniker]

# Example
# desmoscli query profiles profile desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
# desmoscli query profiles profile leonardo
``` 

**REST**
```
/profiles/{address_or_moniker}

# Example
# curl https://morpheus4000.desmos.network/profiles/desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
# curl https://morpheus4000.desmos.network/profiles/leonardo
```