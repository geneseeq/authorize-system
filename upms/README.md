# 用户组服务

针对用户组的CURD,涉及的表group_infos结构

	{
			"_id" : "af",
			"parent" : "",
			"id" : "af",
			"type" : 0,
			"name" : "",
			"code" : "",
			"alias" : "sdfsdf",
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-26T15:40:59.552Z"),
			"update_time" : ISODate("2017-09-26T15:40:59.552Z")
	}


## 新增用户组

* /grouping/v1/group

请求body

	[{
		"id":"af",
		"alias":"sdfsdf",
		"type":0
	}]

返回body

	{
		"status": 200,
		"content": "add user sucessed",
		"sucessedid": [
			"af1"
		],
		"failedid": [
			"af"
		]
	}	


## 更新用户组

* /grouping/v1/group

请求body和返回body同post接口

* /grouping/v1/group/<group_id>

请求body同POST请求

返回body

	{
		"status": 200,
		"content": "update group sucessed"
	}

## 删除用户组

* /grouping/v1/group

请求body

	["af", "af1"]

返回body

	{
		"status": 200,
		"content": "delete mutli group sucessed",
		"sucessedid": [
			"af1"
		],
		"failedid": [
			"af"
		]
	}

* /grouping/v1/group/<group_id>

请求body同DELETE请求

返回body

	{
		"status": 200,
		"content": "delete group sucessed"
	}

## 查找用户组

* /grouping/v1/group

* /grouping/v1/group/<group_id>

返回body

	{
		"content": [
			{
				"id": "af",
				"type": 0,
				"parent": "",
				"name": "",
				"code": "",
				"alias": "sdfsdf",
				"buildin": false,
				"create_user_id": "",
				"create_time": "2017-09-26T23:57:49.269+08:00",
				"update_time": "2017-09-26T23:57:49.269+08:00"
			}
		]
	}

# 用户服务

针对用户的CURD,涉及的表user_infos结构

	{
			"_id" : "af",
			"id" : "af",
			"type" : 0,
			"number" : "",
			"username" : "",
			"tele" : "",
			"gender" : false,
			"status" : 0,
			"validity" : false,
			"vip" : false,
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-26T17:00:16.351Z"),
			"update_time" : ISODate("2017-09-26T17:00:16.351Z"),
			"avatar" : ""
	}



## 新增用户

* /usering/v1/user

请求body

	[{
		"id":"af1"
	},{
		"id":"af1"
	}]

返回body

	{
		"status": 200,
		"content": "add user sucessed",
		"sucessedid": [
			"af1"
		],
		"failedid": [
			"af1"
		]
	}


## 更新用户

* /usering/v1/user

请求body和返回body同post接口

* /usering/v1/user/<user_id>

请求body同POST请求

返回body

	{
		"status": 200,
		"content": "update user sucessed"
	}

## 删除用户

* /usering/v1/user

请求body

	["af", "af1"]

返回body

	{
		"status": 200,
		"content": "delete mutli group sucessed",
		"sucessedid": [
			"af1"
		],
		"failedid": [
			"af"
		]
	}

* /usering/v1/user/<user_id>

请求body同DELETE请求

返回body

	{
		"status": 200,
		"content": "delete user sucessed"
	}

## 查找用户组

* /usering/v1/user

* /usering/v1/user/<user_id>

返回body

	{
		"content": [
			{
				"id": "af",
				"type": 0,
				"parent": "",
				"name": "",
				"code": "",
				"alias": "sdfsdf",
				"buildin": false,
				"create_user_id": "",
				"create_time": "2017-09-26T23:57:49.269+08:00",
				"update_time": "2017-09-26T23:57:49.269+08:00"
			}
		]
	}
