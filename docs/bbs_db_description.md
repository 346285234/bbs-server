# BBS db description



## Topic

* id: int
* user_id: int
* title: string
* content: string
* intro: string
* group_id: int
* is_paste: true
* edit_time: int,
* edit_type: int, // markdown or...
* tag: string
* category_id: string
* favourite_count: int
* like_count: int
* click_count: int 
* created_at: time
* updated_at: time
* deleted_at: time



## Favourite

* id: int
* topic_id: int
* user_id: int
* create_time: int
* modify_time: int

## Like

* id: int
* topic_id: int
* type: int
* user_id: int

- create_time: int
- modify_time: int
- 

## Comment

* id: int
* user_id: int
* topic_id: int
* content: string
* like_count: int
* parent_id: int
* create_time: int
* modify_time: int



## Tag

* value: string



## Category

* id: int
* value: string