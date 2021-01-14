## Query user relationships
This query endpoint allows you to retrieve the details of a relationship where the creator has the given `address`.

**CLI**
```bash
desmosd query relationships user [address]

# Example
# desmosd query relationships user desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```

**REST**
```
/relationships/{address}

# Example
# curl http://lcd.morpheus.desmos.network:1317/relationships/desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
```