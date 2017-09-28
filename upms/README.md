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

# 数据集服务

针对角色的CURD,涉及的表role_infos结构

	{
			"_id" : "test1",
			"id" : "test1",
			"rule" : "",
			"name" : "",
			"match_field" : [
					{
							"data_type" : 0,
							"src_field" : [ ],
							"dest_field" : ""
					}
			],
			"type" : "",
			"validity" : false,
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-27T13:44:52.611Z"),
			"update_time" : ISODate("2017-09-27T13:44:52.611Z")
	}


## 新增数据集

* /seting/v1/data

## 更新数据集

* /seting/v1/data

## 删除数据集

* /seting/v1/data

## 查找数据集

* /seting/v1/data

* /seting/v1/data/<data_id>


# 数据集服务

针对角色的CURD,涉及的表role_infos结构

        {
            "id":"服务id",
            "parent":"父服务",
            "depend":["服务id01","服务id02"]
            "name":"服务名称",
            "level":"服务级别（公司内部服务，公司外部服务）",
            "path":"服务访问路径",
            "register_time":"注册时间",
            "validity":"是否有效",
            "buildin":"内建",
            "owner":["user_type(医生/教师/个人/员工/企业)"],
            "create_user_id":"创建人ID",
            "create_time":"创建时间",
            "status":"状态"
        }


## 新增数据集

* /servicing/v1/service

## 更新数据集

* /servicing/v1/service

## 删除数据集

* /servicing/v1/service

## 查找数据集

* /servicing/v1/service

* /servicing/v1/service/<service_id>


# 组用户和组角色服务

涉及的表group_own_users_and_roles

	{
			"_id" : "test2",
			"group_id" : "test2",
			"user_id" : [
					"dd1235f",
					"fgg"
			],
			"role_id" : [ ],
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-28T10:28:35.503Z"),
			"update_time" : ISODate("2017-09-28T10:28:35.503Z")
	}

## 新增组中的用户

* /releation/v1/group/user

## 更新组中的用户

* /releation/v1/group/user

## 删除组中用户

* /releation/v1/group/user

删除的body为[{"group_id":"test2","user_id":["dd1235f","fgg"]}]

删除用户时，首先是删除user_id,如果user_id和role_id为空，则删除整条记录

## 查找组中用户

* /releation/v1/group/<group_id>/user

* /releation/v1/group/user

## 新增组中的角色

* /releation/v1/group/role

## 更新组中的角色

* /releation/v1/group/role

## 删除组中角色

* /releation/v1/group/role

删除的body为[{"group_id":"test2","user_id":["dd1235f","fgg21"]}]

删除用户时，首先是删除user_id,如果user_id为空，则删除整条记录

## 查找组中角色

* /releation/v1/group/role

* /releation/v1/group/<group_id>/role

# 用户角色关系服务

涉及的表user_own_roles结构

	{
			"_id" : "test2",
			"user_id" : "test2",
			"role_id" : [
					"gys",
					"hhh"
			],
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-28T13:24:20.181Z"),
			"update_time" : ISODate("2017-09-28T13:25:05.332Z")
	}

## 新增用户中的角色

* /releation/v1/user/role

## 更新用户中的角色

* /releation/v1/user/role

## 删除用户中角色

* /releation/v1/user/role

删除的body为[{"user_id":"test2","role_id":["gys2","hhh2"]}]

删除数据权限时，首先是删除role_id,如果role_id为空，则删除整条记录

## 查找用户中角色

* /releation/v1/user/role

* /releation/v1/user/<user_id>/role

# 用户角色关系服务

涉及的表user_own_roles结构

	{
			"_id" : "test2",
			"role_id" : "test2",
			"authority" : [
					{
							"data_id" : "tag01",
							"action" : [
									"dsd"
							]
					}
			],
			"validity" : "",
			"buildin" : false,
			"create_user_id" : "",
			"create_time" : ISODate("2017-09-28T14:03:29.222Z"),
			"update_time" : ISODate("2017-09-28T14:03:29.222Z")
	}

## 新增用户中的角色

* /releation/v1/role/authority

## 更新用户中的角色

* /releation/v1/role/authority

## 删除用户中角色

* /releation/v1/role/authority

删除的body为[{"role_id":"test2","data_id":["tag01"]}]

删除数据权限时，首先是删除data_id字典,如果data_id为空，则删除整条记录

## 查找用户中角色

* /releation/v1/role/authority

* /releation/v1/role/<role_id>/authority