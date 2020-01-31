

## Favorite



### Mark&unmark topic favorite

* **URL:** /favorite/topic/:topic_id/mark
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



### Check topic favorite

- **URL:** /favorite/topic/:topic_id/check

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



### Get topic favorite users

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

