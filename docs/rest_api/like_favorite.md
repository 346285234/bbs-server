## Favorite



### Mark&unmark topic favorite

* **URL:** /favorite/topic/:topic_id/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "unmark": false,
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



### Check topic favorite

- **URL:** /favorite/topic/:topic_id

- **method:** GET

- **header:** user info

- **Response:**

  ```
  {
    "success": true,
    "code": 200,
    "message": "OK",
    "data": {
      "is_mark": true,
    }
  }
  ```



## Like



### Mark&unmark topic like

* **URL:** /like/topic/:topic_id/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "type": 1,
  "unmark": false,
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



### Mark & unmark comment like

* **URL:** /like/comment/:comment_id/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
	"unmark": true,
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



### Check topic like

* **URL:** /like/topic/:topic_id
* **method:** GET
* **header:** user info

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "is_mark": true,
  }
}
```



### Check comment like

* **URL:** /like/comment/:comment_id
* **method:** GET
* **header:** user info

* **Response:**

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": {
    "is_mark": true,
  }
}
```



