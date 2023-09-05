# simple-cdn-server

# Setup
1. Rename **config.json.example** to **config.json**
2. Set your token and address in config.json file
3. Create **/files** directory
4. Execute **go build cdn-server**

# Endpoints
Upload file:
```http
POST /
```

Find file:
```http
GET /files/uuid.ext
```

#### All uploaded files goes to /files directory
