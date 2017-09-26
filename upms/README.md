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
