

## Like



### Mark&unmark topic like

* **URL:** /like/topic/:topic_id/mark
* **method:** POST
* **header:** user info
* **body:**

```json
{
  "is_mark": true,
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
	"is_mark": true,
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

* **URL:** /like/topic/:topic_id/check
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

* **URL:** /like/comment/:comment_id/check
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



### Get topic like users

- **URL:** /like/topic/:topic_id

- **method:** GET

- **header:** user info

- **Response:**

  ```
  {
    "success": true,
    "code": 200,
    "message": "OK",
    "data": {
    	users:[
    		{
    			"id": 1,
    			"name": "",
    			"portrait": "",
    		}
    	]
    }
  }
  ```



### Get comment like users

- **URL:** /like/comment/:comment_id

- **method:** GET

- **header:** user info

- **Response:**

  ```
  {
    "success": true,
    "code": 200,
    "message": "OK",
    "data": {
    	users:[
    		{
    			"id": 1,
    			"name": "",
    			"portrait": "",
    		}
    	]
    }
  }
  ```

