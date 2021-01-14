# Query a session
This query allows you to retrieve the details of a session having its id. 

**CLI**
```bash
desmosd query sessions session [id]

# Example
# desmosd query sessions session 66
```

**REST**
```
/sessions/{session_id}

# Example
# curl http://lcd.morpheus.desmos.network:1317/sessions/66
```
