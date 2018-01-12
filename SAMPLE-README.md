本地测试地址，只是验证了路由信息，至于页面上的文字信息，不要深究其意义~

http://localhost:8080/sample

http://localhost:8080/sample/sayhi


http://localhost:8080/download/test.txt

http://localhost:8080/sample/query

http://localhost:8080/sample/query/1


session共享测试，首先关闭浏览器 ，先打开：
http://localhost:8080/sample2/username 查看输出信息，第一次应该
提示未登录，此时打开 http://localhost:8080/sample/username，然后再刷新

http://localhost:8080/sample2/username  会看到提示信息