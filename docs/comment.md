# Comment

### List comments

* **URL:** /topic/[topic_id]/comments
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
        "author": "x",
        "portrait": "x",
        "content": "x",
        "modify_time": 1111,
        "like_count": 1,
        "subComment": [
          {
            "id": 2,
            "author": "x",
            "portrait": "x",
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

### Add a comment

* **URL:** /topic/[topic_id]/comment/add
* **method:** POST
* **header:** user info
* **body:**

```json
{
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
    "author": "x",
    "portrait": "x",
    "content": "x",
    "modify_time": 1111,
    "like_count": 0,
  },
}
```

### Reply a comment

* **URL:** /topic/[topic_id]/comment/reply
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "id": 1,
  "content": "x"
}
```

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data":
  {
    "id": 1,
    "author": "x",
    "portrait": "x",
    "content": "x",
    "modify_time": 1111,
    "like_count": 0,
  }
}
```

### Mark comment like

* **URL:** /topic/[topic_id]/comment/like/mark
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
  "data": {}
}
```

### Unmark comment like

* **URL:** /topic/[topic_id]/comment/like/unmark
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
  "data": {}
}
```
