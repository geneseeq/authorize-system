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
            "id":"用户ID",
            "type":"医生/教师/个人/员工/企业",
            "number":"工号",
            "username":"用户名",
            "tele":"电话",
            "gender":"性别",
            "status":"岗位状态"(离职，全职，实习),
            "validity":"是否合法",
            "vip":"是否是vip",
            "buildin":"内建用户",
            "create_user_id":"创建人ID",
            "create_time":"创建时间",
            "update_time":"更新时间",
            "avatar":"http://.../avatar.jpg",
        }

## 新增用户

* /usering/v1/user

## 更新用户

* /usering/v1/user

* /usering/v1/user/<user_id>

## 删除用户

* /usering/v1/user

* /usering/v1/user/<user_id>

## 查找用户组

* /usering/v1/user

* /usering/v1/user/<user_id>


# 角色服务

针对角色的CURD,涉及的表role_infos结构

	{
		"type":"岗位"（主岗，副岗，职务）,
		"name":"名称",
		"alias":"别名",
		"id":"角色ID",
		"buildin":"内建用户组",
		"create_user_id":"创建人ID",
		"create_time":"创建时间",
		"update_time":"更新时间"
	}


## 新增角色

* /roleing/v1/role

## 更新角色

* /roleing/v1/role

* /roleing/v1/role/<role_id>

## 删除角色

* /roleing/v1/role

* /usering/v1/user/<role_id>

## 查找角色

* /usering/v1/user

* /usering/v1/user/<role_id>
