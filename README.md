# Template CRUD API Project

### Create author

#### Request
```curl
curl --location 'localhost:8080/api/authors' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "name": "Zhanserik",
    "lastname": "Rakhmet",
    "username": "unknownqazaq",
    "specialization": "Software Engineer"
}'
```

#### Response
```json
{
    "id": 1,
    "name": "Zhanserik",
    "lastname": "Rakhmet",
    "username": "unknownqazaq",
    "specialization": "Software Engineer"
}
```


#### Request

### Get author

#### Request
```curl
curl --location 'localhost:8080/api/authors/1'
```

#### Response
```json
{
    "id": 1,
    "name": "Zhanserik",
    "lastname": "Rakhmet",
    "username": "unknownqazaq",
    "specialization": "Software Engineer"
}
```

### Get all authors

#### Request
```curl
curl --location 'localhost:8080/api/authors'
```

#### Response
```json
{
    "1": {
        "id": 0,
        "name": "",
        "lastname": "",
        "username": "",
        "specialization": ""
    },
    "2": {
        "id": 1,
        "name": "Zhanserik",
        "lastname": "Rakhmet",
        "username": "unknownqazaq",
        "specialization": "Software Engineer"
    }
}
```

### Update author

#### Request
```curl
curl --location 'localhost:8080/api/authors/1' \
--header 'Content-Type: application/json' \
--data '{
    "specialization": "Best Software Engineer"
}'
```

#### Response
```json
{
    "id": 1,
    "specialization": "Best Software Engineer"
}
```

### Delete author

#### Request
```curl
curl --location --request DELETE 'localhost:8080/api/authors/1'
```