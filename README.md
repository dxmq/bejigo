# bejigo

#### 介绍
基于Golang的博客系统

#### 软件架构
基于Beego框架，数据库采用sqlite3，安装简单，无需配置数据库

#### 安装教程
1、安装所需库
- go get github.com/astaxie/beego
- go get github.com/beego/bee
- go get github.com/mattn/go-sqlite3

2、克隆项目到 `$GOPATH/src` 目录
- cd $GOPATH/src/
- git clone https://github.com/dxmq/bejigo.git

3、测试是否有错误
- cd bejigo
- bee run 看是否出错

4、编译
- go build main.go

5、改运行模式为生产模式，也可以不改变
- vim bejigo/conf/app.conf
- runmode = "prod"
> 注意：项目使用了注解路由，如果开启了生产模式，`router/commentsRouter_controllers.go` 将不再自动生成

6、将应用部署在后端
- 在bejigo目录下 `nohup ./main&`

7、`nginx` 部署
参考 https://beego.me/docs/deploy/nginx.md