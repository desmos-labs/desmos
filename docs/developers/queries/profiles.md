# Query the stored profiles
This query endpoint allows you to get all the stored profiles.

**CLI**
 ```bash
desmoscli query profiles all

# Example
# desmoscli query profiles all
``` 

**REST**
```
/profiles

# Example
# curl http://lcd.morpheus.desmos.network:1317/profiles
``` 

# Query a profile with the given moniker
This query endpoint allows you to get the profile related to the given `moniker`.

**CLI**
 ```bash
desmoscli query profiles profile bob

# Example
# desmoscli query profiles profile bob
``` 
**REST**
```
/profiles/{address_or_moniker}

# Example
# curl http://lcd.morpheus.desmos.network:1317/profiles/bob
``` 