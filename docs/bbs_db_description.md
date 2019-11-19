# BBS db description

## User

* id: int
* name: string
* portrait: string
* favourites: [int]
* like: [int]



## Topic

* id: int

* title: string
* tag: string
* category: string
* user_id: int
* modify_time: int
* description: string
* content: string
* favourites_count: int
* like_count: int
* click_count: int 



## Comment

* id: int
* user_id: int
* topic_id: int
* content: string
* modify_time: int
* like_count: int
* parent_id: int
* sub_id: [int]



## Tag

* value: string



## Category

* value: string