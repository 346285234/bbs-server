# Comment

### List comments

* **URL:** /comments/:topic_id
* **method:** GET
* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": 
  {
    "total": 1,
    "comments": [
      {
        "id": 1,
        "author": {"id": 1, "name": "x", "portrait": "x"},
        "content": "x",
        "modify_time": 1111,
        "like_count": 1,
        "subComment": [
          {
            "id": 2,
            "author": {"id": 1, "name": "x", "portrait": "x"},
            "content": "x",
            "modify_time": 1111,
            "like_count": 1,
          }
        ],
      }
    ]
  }
}
```

### Reply comment

* **URL:** /comment/:topic_id/reply
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "parent_id": -1,
  "content": "x",
}
```

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "id": 1,
    "author": {"id": 1, "name": "x", "portrait": "x"},
    "content": "x",
    "modify_time": 1111,
    "like_count": 0,
  },
}
```



### Revoke comment

* **URL:** /comment/:topic_id/revoke
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "id": 1,
}
```

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {},
}
```

