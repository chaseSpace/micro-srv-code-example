--- 
这个工程演示了一个单体go工程的项目布局，在这个工程内可能包含多个app；
所有的目录都是小写，多语义请加下划线，目录都是可选的，根据需要创建
  
#### 如果想要启动一个新的单体go工程，可直接clone本仓库，然后基于它开发你的应用。


#### 编译
```bash
cd .../go_project_template/app/apple
go build -o apple.exe  # get main.exe

# 代码中从 命令行/ENV 读取项目配置文件
./main --conf-path=../configs/app_apple.toml
```