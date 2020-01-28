## follow

### follow&unfollow user

* **URL:** /follow/user/:user_id/mark
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



### Check follow user

- **URL:** /follow/user/:user_id/check

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

