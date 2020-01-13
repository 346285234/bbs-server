### List tag

* **URL:** /tags
* **method:** GET
* **header:** user info
* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "tags": ["a", "b"]
  }
}
```

### List category

* **URL:** /categories
* **method:** GET
* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "categories": [
     {
       "id": 1,
       "name": "x"
     }
    ]
  }
}
```

