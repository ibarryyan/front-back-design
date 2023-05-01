# front-back-design
前后端分离设计Demo，包含：

- Spring Boot + Vue前后端分离
- Gin + Vue前后端分离

### 1 什么是前后端分离

![在这里插入图片描述](https://img-blog.csdnimg.cn/c9af4e146e694d4a949efef38c221acf.png)

前后端分离（Front-end and Back-end Separation）指的是将应用程序的前端和后端分离开来，以便更好地实现业务逻辑和数据处理的解耦。

在传统的应用程序中，前端和后端往往是密不可分的。前端负责处理用户的交互和呈现，而后端负责处理数据的存储和处理。这种模式使得前端和后端之间存在着严重的依赖关系，难以进行独立的开发和维护。

前后端分离的好处是可以使得前端开发人员更加专注于页面设计和用户体验，而后端开发人员则更加专注于业务逻辑的实现。这样可以提高开发效率，减少错误，并且使得应用程序更加灵活和可扩展。

前后端分离通常包括以下几个方面：

1. **数据处理层的分离**：将数据处理从前端分离出来，放到后端专门的服务器上，可以提高数据处理的效率和可靠性。
2. **业务逻辑层的分离**：将业务逻辑从前端分离出来，放到后端专门的服务器上，可以使得不同的业务逻辑被更好地单独处理，提高了代码的复用性和可维护性。

总之，前后端分离是一种架构模式，它可以使得前后端开发团队之间更好地协作和独立开发，提高开发效率，减少错误，并且使得应用程序更加灵活和可扩展。

### 2 运行环境

#### 2.1 Spring Boot【后端】

##### 环境要求：

JDK 1.8

Maven 3.x

MySQL 5.x

##### 运行方式：

```shell
mvn install

mvn spring-boot:run
```

#### 2.2 Vue【前端】

##### 环境要求：

node 16.x

##### 运行方式：

```shell
npm install

npm run serve
```

#### 2.3 Go【后端】

##### 环境要求：

Go 1.17

##### 运行方式：

```shell
go mod tidy

go run main.go
```

### 3 相关博客讲解

#### 3.1 手把手教你搭建Spring Boot+Vue前后端分离

https://blog.csdn.net/Mr_YanMingXin/article/details/121870846

#### 3.2 一文搞懂Go+Vue前后端分离设计实践

https://blog.csdn.net/Mr_YanMingXin/article/details/130451171

### 4 TODO

- 后端返回值需要规范
- 安全性问题
- Go的DB操作换成ORM可节省大量代码
- ......

---

如有问题可邮件1712229564@qq.com

> 公众号【扯编程的淡】：
>
> ![image-20230430204003260](https://img-blog.csdnimg.cn/56b9c71cee5443048fa4b98ba8d98e82.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBATXJfWWFuTWluZ1hpbg==,size_20,color_FFFFFF,t_70,g_se,x_16)

> 感谢您的支持，可扫码进行赞赏<**微信**>:
>
> <img src="https://img-blog.csdnimg.cn/8233dee0e742434e8e1a1684004a7f4e.png" alt="微信图片_20230430202859" style="zoom: 70%;" />

