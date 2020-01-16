# Topic

TODO:
Missing page.

user info: ["userID": 1, ]

"image_url": "x",

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
        "tag": ["a", "b"],
        "category_id": 1,
        "category_name": "x",
        "author_id": 1,
        "author_name": "x",
        "author_portrait": "x",
        "description": "x",
        "content": "", // no content
        "modify_time": 1111,
        "favourite_count": 1,
        "like_count": 1,
        "view_count": 1,
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
    "tag": ["a", "b"],
    "category_id": 1,
    "category_name": "x",
    "author_id": 1,
    "author_name": "x",
    "author_portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favorite_count": 1,
    "like_count": 1,
    "view_count": 1,
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
  "tags": ["a", "b"],
  // future use.
  "edit_time": 1111, // edit use how many time
  "is_paste": true,
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
    "tags": ["a", "b"],
    "category_id": 1,
    "category_name": "x",
    "author_id": 1,
    "author_name": "x",
    "author_portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favorite_count": 1,
    "like_count": 1,
    "view_count": 1,
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
  "content": "x",
  "category_id": 1,
  "tag": ["a", "b"],
  // future use.
  "edit_time": 1111, // edit use how many time
  "is_paste": true,
  "edit_type": 1, // markdown or...
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
    "tag": ["a", "b"],
    "category_id": 1,
    "category_name": "x",
    "author_id": 1,
    "author_name": "x",
    "author_portrait": "x",
    "description": "x",
    "content": "x",
    "modify_time": 1111,
    "favorite_count": 1,
    "like_count": 1,
    "view_count": 1,
  }
}
```

