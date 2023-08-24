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
    |-  manager     manager层
    |-  middleware  中间件
    |-  model       实体类包 
    |-  service     service层
    |-  terrs       tiktok errors的缩写，用于存放自定义的错误信息和错误码等
|- pkg              用于存放公共接口
    |-  logging     日志相关
    |-  util        工具类
```

api中目前设计了⑤个接口:
1. favorite,用于处理与用户对视频点赞行为相关的请求
2. relation社交接口,用于处理用户对用户的行为
3. user用户接口，用户的基础信息相关功能
4. video视频接口，视频相关的基础功能

设计了四个层次
1. api 用于http数据的解析，参数的校验，调用service层的代码，返回service代码执行结果
2. service 用于处理业务逻辑，不允许service之间互相调用，可以调用manager和dao层的代码
3. manager 用于提供简单的业务逻辑封装供service接口复用,如果需要使用事务功能也需要封装在manager层，调用dao层代码
4. dao 用于直接和数据库交互

如果需要运行请对`/configs/conf.ini`文件进行适当修改

目前的错误处理流程：
任何层抛出未预先定义的异常都需要打印异常信息
dao,service层的非预定义异常都要以ErrInternal形式抛出
对于预先定义的异常进行上抛。
对于controller层出现的错误，使用`AbortWithStatusErrJSON()`函数传递给前端即可

```ini
[database]
username = 
password = 
host = # 192.168.157.129:3306
schema = # tiktok_go

[jwt]
signedKey = tik-jwt-tok

[tiktok]
userAvatar = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/5657526a-4e91-4b36-9633-fe3f30f2e281.jpg
videoTitle = untitled
videoCover = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/ffb34b76-a294-4135-80a3-ad189cc61432.jpg
backgroundImage = https://web-lcx-test.oss-cn-beijing.aliyuncs.com/ffb34b76-a294-4135-80a3-ad189cc61432.jpg
feedSize = 30

[oss]
endPoint = 
accessKeyId =
accessKeySecret = 
bucketName = 


```