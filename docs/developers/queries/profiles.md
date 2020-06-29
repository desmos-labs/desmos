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
# curl https://morpheus7000.desmos.network/profiles
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
# curl https://morpheus7000.desmos.network/profiles/bob
``` 

# Query profiles module parameters
This query endpoint returns all the parameters of the profiles module.

**CLI**
 ```bash
desmoscli query profiles params

# Example
# desmoscli query profiles params
``` 
**REST**
```
/profiles/params

# Example
# curl https://morpheus7000.desmos.network/profiles/params
``` 
