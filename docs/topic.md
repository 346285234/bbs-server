# Topic

### List topics

* **URL:** /topics
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
    "topics": [
      {
        "id": 1,
        "title": "x",
        "tag": "x",
        "category": "x",
        "author": "x",
        "portrait": "x",
        "description": "x",
        "modify_time": 1111,
        "favourites_count": 1,
        "likes_count": 1,
        "click_count": 1,
      }
    ]
  }
}
```

### Get a topic

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
    "title": "x",
    "tag": "x",
    "category": "x",
    "author": "x",
    "portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favourites_count": 1,
    "likes_count": 1,
    "click_count": 1,
  }
}
```

### Add a topic

* **URL:** /topic/add
* **method:** POST
* **head:** user info
* **body:**

```json
{
  "title": "x",
  "category": "x",
  "tag": "x",
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
    "title": "x",
    "tag": "x",
    "category": "x",
    "author": "x",
    "portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favourites_count": 1,
    "likes_count": 1,
    "click_count": 1,
  }
}
```

### Remove a topic

* **URL:** /topic/remove
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

### Modify a topic

* **URL:** /topic/modify
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "id": 1,
  "title": "x",
  "tag": "x",
  "category": "x",
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
    "title": "x",
    "tag": "x",
    "category": "x",
    "author": "x",
    "portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favourites_count": 1,
    "likes_count": 1,
    "click_count": 1,
  }
}
```

### List user favourites topics

* **URL:** /user/favourites/list
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
    "topics": [
      {
        "id": 1,
        "title": "x",
        "tag": "x",
        "category": "x",
        "author": "x",
        "portrait": "x",
        "description": "x",
        "modify_time": 1111,
        "favourites_count": 1,
        "likes_count": 1,
        "click_count": 1,
      }
    ]
  }
}
```

### Mark favourites

* **URL:** /topic/favourites/mark
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

### Unmark favourites

* **URL:** /topic/favourites/unmark
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

### List user like topics

* **URL:** /user/like/list
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
    "topics": [
      {
        "id": 1,
        "title": "x",
        "tag": "x",
        "category": "x",
        "author": "x",
        "portrait": "x",
        "description": "x",
        "modify_time": 1111,
        "favourites_count": 1,
        "likes_count": 1,
        "click_count": 1,
      }
    ]
  }
}
```

### Mark like

* **URL:** /topic/like/mark
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

### Unmark like

* **URL:** /topic/like/unmark
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

### List tag

* **URL:** /topic/tag/list
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
    "tags": [
      "value": "x",
    ]
  }
}
```

### Add a tag

* **URL:** /topic/tag/add
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "value": "x",
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

### List category

* **URL:** /topic/category/list
* **method:** GET
* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "total": 1,
    "category": [
      "value": "x",
    ]
  }
}
```