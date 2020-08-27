## Query user relationships
This query endpoint allows you to retrieve the details of a relationship where the creator is the given `address`.

**CLI**
```bash
desmoscli query relationships user_relationships [address]

# Example
# desmoscli query relationships user_relationships desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

**REST**
```
/relationships/{address}

# Example
# curl http://lcd.morpheus.desmos.network:1317/relationships/desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```