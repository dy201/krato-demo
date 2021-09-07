# krato-demo
krato-demo


# 项目概述
目前只完成了两个app功能开发， blog_admin和user

# 启动
1. 分别在两个app中修改 config.yaml，我自己在虚拟机中搭了环境，改了ip。
2. 进入 `doyle-mall\app\user\service\cmd\server` 和 `doyle-blog\app\blog\admin\cmd\server` 目录下运行 `go run .`启动项目

# 问题描述
启动成功后，访问 `http://localhost:8080/admin/user/list` 报错 `{"code":500,"reason":"","message":"parsing field \"id\": strconv.ParseUint: parsing \"list\": invalid syntax","metadata":{}}`

访问其它接口正常，比如`http://localhost:8080/admin/user/2` 返回值 `{"user":{"id":2,"userName":"doyle","password":"123456","nickname":"小良同学","status":0,"email":"doyle@123.com","phone":"18801020304","roleId":0,"age":22,"createAt":null,"updateAt":null}}`

# 数据库
放在deploy/mysql文件夹中, 数据库名`doyle_blog`

# 其它
swagger json 文件，我没更新。

进入blog_admin的proto目录，运行 `protoc --proto_path=.  --proto_path=../../../../third_party  --openapiv2_out .  --openapiv2_opt logtostderr=true blog_admin.proto` 更新即可。