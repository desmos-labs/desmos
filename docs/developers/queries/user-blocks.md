## Query user blocked users
This query endpoint allows you to retrieve the user blocked by the user with the given `address`.

**CLI**
```bash
desmoscli query relationships user-blocks [address]

# Example
# desmoscli query relationships user-blocks desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

**REST**
```
/userBlocks/{address}

# Example
# curl http://lcd.morpheus.desmos.network:1317/userBlocks/desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```