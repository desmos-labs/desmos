## Query user's DTag requests
This query endpoint allows you to retrieve the DTag requests of the user with the given `address`.

**CLI**
```bash
desmoscli query profiles dtag-requests [address]

# Example
# desmoscli query relationships dtag-requests desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

**REST**
```
/profiles/{address}/incoming-dtag-requests

# Example
# curl http://lcd.morpheus.desmos.network:1317/profiles/desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud/incoming-dtag-requests
```