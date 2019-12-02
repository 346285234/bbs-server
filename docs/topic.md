# Topic

Missing page.

### List topics

* **URL:** /topics

* **method:** GET

* **query:** 

  |      |      |      |
  | ---- | ---- | ---- |
  |      |      |      |
  |      |      |      |
  |      |      |      |

  

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
        "category_id": 1,
        "category_name": "x",
        "author_id": 1,
        "author_name": "x",
        "portrait": "x",
        "description": "x",
        "image_url": "x",
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
    "category_id": 1,
    "category_name": "x",
    "author_id": 1,
    "author_name": "x",
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
  "category_id": 1,
  "category_name": "x",
  "tag": "x",
  "content": "x",
  "edit_time":1111,
  "is_paste":true,
  "edit_type":1,
  
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
    "category_id": 1,
    "category_name": "x",
    "author_id": 1,
    "author_name": "x",
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
  "category_id": 1,
  "category_name": "x",
  "content": "x",
  "edit_time":"x",
  "is_paste":true,
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
    "category_id":1,
    "category_name":"x",
    "author_id": 1,
    "author_name": "x",
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

### Mark&unmark favourites

* **URL:** /topic/favourites/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "id": 1,
  "type": true,
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

### Mark&unmark like

* **URL:** /topic/like/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "id": 1,
  "type": true,
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
      "id":1,
      "name": "x",
    ]
  }
}
```