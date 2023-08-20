# tiktok_go
字节跳动青训营后端项目go语言实现

```
|- api              controller实现
|- bin              用来糊弄go build的（out put directory
|- configs          用来存配置文件，无代码
|- internal         api接口的内部实现
    |-  cache       缓存接口
    |-  conf        用于加载配置信息的包，类似Spring Boot中的Properties
    |-  dao         dao层
    |-  middleware  中间件
    |-  model       实体类包 
    |-  service     service层
    |-  terrs       tiktok errors的缩写，用于存放自定义的错误信息和错误码等
|- pkg              用于存放公共接口
    |-  logging     日志相关
    |-  util        工具类
```

如果需要运行请对`/configs/conf.ini`文件进行适当修改

目前的错误处理流程：
任何层抛出未预先定义的异常都需要打印异常信息。
对于预先定义的异常进行上抛。
对于controller层出现的异常，如果是预先定义的异常则直接以JSON形式返回，如果是未定义的异常则作为InternalError返回。

```ini
[database]
username = mysql用户名
password = mysql密码
#改成自己的ip和端口
host = 192.168.157.129:3306 
schema = 数据库表名

[jwt]
signedKey = tik-jwt-tok

[default]
userAvatar = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/5657526a-4e91-4b36-9633-fe3f30f2e281.jpg
videoTitle = untitled
videoCover = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/ffb34b76-a294-4135-80a3-ad189cc61432.jpg
backgroundImage = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/ffb34b76-a294-4135-80a3-ad189cc61432.jpg


[oss] # 阿里oss配置
endPoint = 
accessKeyId = 
accessKeySecret = 
bucketName = 

```