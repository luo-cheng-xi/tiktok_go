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
