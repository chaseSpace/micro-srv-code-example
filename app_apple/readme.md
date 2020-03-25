
#### /app目录存放这个应用的几乎所有的逻辑代码

---
/app/main.go 作为app的启动文件
/app/APPNAME_config 存放配置代码，如struct，以及外部获取和修改struct的方法, 注意命名包含app名称，防止多个app时config引入混乱造成错误。