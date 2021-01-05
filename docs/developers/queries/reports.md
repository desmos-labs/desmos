# Query the stored reports
This query endpoint allows you to get all the stored reports related to the given
post ID. 

**CLI**
```bash
desmoscli query reports post [id]

# Example
# desmoscli query reports all 301921ac3c8e623d8f35aef1886fea20849e49f08ec8ddfdd9b96feaf0c4fd15
```

**REST**
```
/reports/{post_id}

# Example
# curl http://lcd.morpheus.desmos.network:1317/reports/301921ac3c8e623d8f35aef1886fea20849e49f08ec8ddfdd9b96feaf0c4fd15
```