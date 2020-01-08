# Topic

TODO:
Missing page.

user info: ["userID": 1, ]

### List topics

* **URL:** /topics?hot=true&group_id=1

* **method:** GET

* **query:** 

  | key          | value |
  | ------------ | ----- |
  | hot          | true  |
  | group_id     | 1     |
  | distance     | 1000  |
  | user_id      | 1     |
  | category_id  | 1     |
  | tag          | "x"   |
  | page_size    | 1000  |
  | page         | 1     |
  | key_word     | "x"   |
  | last_post_id | 1     |

  

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": 
  {
    "total": 1,
    "page": 1,
    "page_size": 1000,
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
  "content": "x",
  "category_id": 1,
  "tags": [{"id": 1, "value": "a"},{"value": "b"}],
  "edit_time":1111, // edit use how many time
  "is_paste":true,
  "edit_type":1, // markdown or...
  "group_id": 1,
  
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
    "tags": [{"id": 1, "value": "a"}, {"id": 2, "value":  "b"}],
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

### Update a topic

* **URL:** /topic/update
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

* **URL:** /topic/tags
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
      "id":1,
      "name": "x",
    ]
  }
}
```