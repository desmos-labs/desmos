# Query a profile
This query endpoint allows you to retrieve the details of a single profile having its moniker. 

**CLI**
 ```bash
desmoscli query profile profile [moniker]

# Example
# desmoscli query profile profile mrCake
``` 

**REST**
```
/profile/{moniker}

# Example
# curl https://morpheus4000.desmos.network/profile/mrCake
```