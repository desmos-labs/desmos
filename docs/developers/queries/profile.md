# Query a profile
This query endpoint allows you to retrieve the details of a single profile having its DTag or address. 

**CLI**
 ```bash
desmoscli query profiles profile [address_or_dtag]

# Example
# desmoscli query profiles profile desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
# desmoscli query profiles profile leonardo
``` 

**REST**
```
/profiles/{address_or_dTag}

# Example
# curl http://lcd.morpheus.desmos.network:1317/profiles/desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
# curl http://lcd.morpheus.desmos.network:1317/profiles/leonardo
```