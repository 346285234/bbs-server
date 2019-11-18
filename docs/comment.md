# Comment

### Get comments

* **URL:** /topic/comments?topic_id=[id]
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
        "user_id": "x",
        "user_name": "x",
        "content": "x",
        "subcomment": [
          {
            "id": 2,
            "user_id": "x",
            "user_name": "x",
            "content": "x",
          }
        ],
      }
    ]
  }
}
```

### Get a comment

* **URL:** /topic/[id]
* **method:** GET
* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": 
  {
    "id": 1,
    "name": "x",
    "author": "x",
    "content": "x"
  }
}
```

### Add comment

* **URL:** /topic/comment/add
* **method:** POST
* **body:**

```json
{
  "topic_id": 1,
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
    "comment_id": 1,
    "content": "x",
    "modify_time": 11111,
  },
}
```

### Reply comment

* **URL:** /topic/comment/reply
* **method:** POST
* **body:**

```json
{
  "topic_id": "x",
  "author": "x",
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
    "id": 1
  }
}
```

### Modify a topic

* **URL:** /topic/modify
* **method:** POST
* **body:**

```json
{
  "name": "x",
  "author": "x",
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
    "id": 1
  }
}
```

### add favourites

* **URL:** /topic/favourites/add
* **method:** POST
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

### remove favourites

* **URL:** /topic/favourites/remove
* **method:** POST
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

### list favourites

* **URL:** /topic/favourites/list

* **method:** GET

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "users": [
      {
        "userID": 1,
        "username": "cq",
      }
    ]
  }
}
```

### add like

* **URL:** /topic/like/add
* **method:** POST
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

### remove like

* **URL:** /topic/like/remove
* **method:** POST
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

### list like

* **URL:** /topic/like/list

* **method:** GET

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "users": [
      {
        "userID": 1,
        "username": "cq",
      }
    ]
  }
}
```

### list tag

* **URL:** /topic/tag/list

* **method:** GET

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "tags": [
      "name": "x",
    ]
  }
}
```

### add like

* **URL:** /topic/tag/add
* **method:** POST
* **body:**

```json
{
  "name": "x",
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