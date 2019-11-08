# Topic

### List topics

* **URL**: /
* **method**: GET
* **produce**: application/json

```json
{
  "statusCode": 200,
  "message": "OK",
  "results": 
  {
    "total": 1,
    "topics": [
      {
        "id": 1,
        "name": "x",
        "author": "x",
        "intro": "x"
      }
    ]
  }
}
```

### Get a topic

* **URL**: /topic?id=xxx
* **method**: GET
* **produce**: application/json

```json
{
    "id": 1,
    "name": "x",
    "author": "x",
    "content": "x"
}
```

### Create a topic

* **URL**: /topic/create

* **method**: POST

* **consume**: application/json

```json
{
    "name": "x",
    "author": "x",
    "content": "x"
}
```

* **produce**: application/json

```json
{
  "statusCode": 200,
  "message": "OK",
  "results":
  {
    "id": 1
  }
}
```
