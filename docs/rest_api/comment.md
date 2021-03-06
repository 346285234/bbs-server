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
        "author_id": 1,
        "author_name": "x",
        "author_portrait": "x",
        "content": "x",
        "modify_time": 1111,
        "like_count": 1,
        "sub_comments": [
          {
            "id": 2,
            "author_id": 1,
    				"author_name": "x",
    				"author_portrait": "x",
            "content": "x",
            "modify_time": 1111,
            "like_count": 1,
            "sub_comments": [],
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
  "parent_id": 0, // 0: no parent
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
    "author_id": 1,
    "author_name": "x",
    "author_portrait": "x",
    "content": "x",
    "modify_time": 1111,
    "like_count": 0,
    "sub_comments": [],
  },
}
```



### // Revoke comment

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

