# Query a profile
This query endpoint allows you to retrieve the details of a single profile having its moniker. 

**CLI**
 ```bash
desmoscli query profile profile [address]

# Example
# desmoscli query profile profile desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
``` 

**REST**
```
/profiles/{address}

# Example
# curl https://morpheus4000.desmos.network/profile/desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
```