# 标签缓存服务

针对标签缓存的CURD,涉及的表label_own_id结构

        {
            "label_id":"标签唯一id",
            "sample_id":["sample_id01","sample_id02"],
            "order_id":["order_id01","order_id02"],
            "update_time":"数据更新时间"
            "create_time":"数据创建时间"
        }


## 新增数据集

* /labeling/v1/label

## 更新数据集

* /labeling/v1/label

## 删除数据集

* /labeling/v1/label

## 查找数据集

* /labeling/v1/label

* /labeling/v1/label/<label_id>

# 基本信息服务

针对基本信息缓存的CURD,涉及的表data_infos结构

        {
            "id":"数据唯一id",
            "sample_id":"样本id",
            "order_id":"订单id",
            "sale_id":"销售代表",
            "doctor":"医生",
            "hospital":"医院",
            "hospital_dept":"科室",
            "school":"学校",
            "school_dept":"院系",
            "product":"产品",
            "project":"项目",
            "create_time":"数据创建时间",
            "update_time":"数据更新时间",
            "label_id":["lab01","lab02","lab03"] #定时从缓存中存入
        }


## 新增基本信息

* /baseing/v1/data

## 更新基本信息

* /baseing/v1/data/

## 删除基本信息

* /baseing/v1/data/

## 查找基本信息

* /baseing/v1/data/

* /baseing/v1/data/<data_id>