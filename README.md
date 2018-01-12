#manword go 后端工程

## 安装执行命令

1、 首先，确保本地安装好go环境 ，参见 https://golang.org/doc/

2、 本工程采用 beego框架，具体可参见 https://beego.me/

3、 保证第一步环境安装后，如果有时间可以参与第二步的文档（完全能解决问题~），如果没有时间，
    跟着下面的命令执行即可（包括但不限于下面命令，因为我可能忘记~,出现包未引用错误 call me）：

   （1）get -u github.com/astaxie/beego
   （2）go get -u github.com/beego/bee
   （3）go get github.com/astaxie/beego/session
   （4）go get -u github.com/go-sql-driver/mysql
   
   
4、运行
   确保以上安装完毕后，执行 bee run 命令即可。验证地址则参见SAMPLE-README.md


## 目前已完成工作

    1、基础框架（可以认为包含了一下工作量~）
    2、session机制
    3、数据库mysql引入（数据库配置已完成，不需要修改任何配置，因为链接的是远端数据库~）
    4、路由控制
    
    Todo: coding业务代码...