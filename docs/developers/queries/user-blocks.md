## Query user blocked users
This query endpoint allows you to retrieve the user blocked by the user with the given `address`.

**CLI**
```bash
desmoscli query relationships blocklist [address]

# Example
# desmoscli query relationships blocklist desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

**REST**
```
/blocklist/{address}

# Example
# curl http://lcd.morpheus.desmos.network:1317/blocklist/desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```