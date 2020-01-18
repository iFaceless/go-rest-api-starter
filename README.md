# 介绍

使用 Go 语言构建 RESTful API 项目模板。该项目以开发算法商品管理接口为例（源自 [aizoo.com](https://aizoo.com) 后端项目），演示了我们在实践中采用的项目结构。由于依赖的某些工具或者框架尚未开源，所以这里使用了一些开源框架作为替代演示了我们的实践想法，也方便理解和应用。

需要声明的是，该项目结构包括实践中用到的一些工具或者想法，仅供参考，相信不同的团队有不同的想法，但本质上都是服务于业务。我们希望能够以一致、清晰的方式来编写代码，并且保证项目结构不被随意破坏，项目的可维护性大于一切。合理地使用业务框架，既有利于简化业务代码编写，又利于理解和维护。而对于某些高并发的接口，也许问题不是出现在框架上，你应该首先考虑是否使用了正确的数据库索引？是否添加了 Redis 缓存？是否可以依靠内存缓存加速？某些操作是否可以异步化？如果到最后，的确是框架引起的问题，那么可以考虑替换框架或者改进框架，但是原则还是要**保证可维护性**，避免提前优化，避免过渡优化，权衡好收益！

# 项目结构

```
├── .env（环境变量配置，资源连接串等）
├── .env_unittest（跑单元测试使用的测试资源配置）
├── LICENSE
├── Makefile
├── README.md
├── bin
│   ├── starter-admin（面向管理后台的 RESTful API 服务器）
│   └── starter-web（面向客户端、PC 等前台的 RESTful API 服务器）
├── cmd（各个服务启动的入口）
│   ├── admin（管理后台）
│   │   └── main.go
│   ├── bee（离线异步任务）
│   │   └── main.go
│   ├── service（RPC 服务）
│   │   └── main.go
│   └── web（客户端、PC、小程序等前台）
│       └── main.go
├── go.mod（依赖包 go modules）
├── pkg
│   ├── admin（管理后台服务）
│   │   ├── handler
│   │   ├── router.go
│   │   ├── schema（和返回的 JSON 数据关联的资源结构体定义）
│   │   └── validator（通用的校验代码）
│   ├── config（资源配置）
│   │   ├── fixture.go
│   │   ├── init.go
│   │   └── mysql.go
│   ├── constant（常量、枚举定义）
│   │   ├── enum.go
│   │   ├── gen_enum.go
│   │   └── macro.go
│   ├── controller（复杂业务逻辑）
│   │   ├── company.go
│   │   ├── company_test.go
│   │   └── product.go
│   ├── job（异步离线任务业务逻辑）
│   │   └── after_product_created.go
│   ├── middleware（可复用的中间件）
│   │   └── cors.go
│   ├── model（Models，可能还会聚合来自 RPC 等数据源，数据模型抽象）
│   │   ├── company.go
│   │   ├── doc.go
│   │   ├── init.go
│   │   └── product.go
│   ├── util（工具集）
│   │   ├── orderby.go
│   │   ├── pic
│   │   ├── rest
│   │   ├── seqgen
│   │   └── toolkit
│   └── web（前台服务）
│       ├── handler
│       ├── router.go
│       └── schema
├── script（一些脚本文件）
│   └── 20200101
└── testdata（单元测试有关测试数据、表结构定义）
    ├── fixtures
    │   ├── company.yml
    │   ├── doc.yml
    │   └── product.yml
    └── schema.sql
```

# 构建 & 测试 & 运行

首先，克隆该项目到本地：`git clone git@github.com:iFaceless/go-rest-api-starter.git`

接下来，运行 `go mod tidy` 安装好依赖包

## 构建

构建生成后台和前台的 Web 服务可执行程序：`make all`，之后可以在 `bin` 目录下找到生成的可执行文件。

```
bin
├── starter-admin
└── starter-web
``` 

## 运行

我们使用 [godotenv](https://github.com/joho/godotenv) 加载配置信息到环境变量中，一些关键的连接串信息、服务配置等都存放在 `.env` 文件中。

在 `.env` 文件中，需要修改 MySQL 数据库连接串，确保对应的数据库名称已经创建。

接下来，启动 Web 服务吧~

```
DOT_ENV_FILE=.env ./bin/starter-admin
DOT_ENV_FILE=.env ./bin/starter-web
```

## 测试

给 `controller/company` 添加了简单的测试，主要是演示如何使用 [fixture](https://github.com/ifaceless/fixture) 工具在测试开始时往测试库中建表并准备测试数据，并在测试结束时清理测试数据。

在准备测试前，需要记得修改 `.env_unittest` 中关于测试库连接串信息，并且创建好 `test_go_starter` 数据库。需要注意的是，测试库必须要以 `test_` 开头，尽可能避免出错！

一切就绪后，使用 `make test` 即可运行单元测试咯。

# 在 Docker 中运行

什么？上面的操作太麻烦了！我们为你提供了 `docker-compose.yml` 文件，只需要借助 `docker-compose` 即可一键启动：
1. `cd docker`
2. 构建 & 运行：`docker-compose run --build`
3. 后台运行：`docker-compose run -d`
4. 终止运行的实例：`docker-compose down -v`

![](https://pic4.zhimg.com/80/v2-7e30fdab43678fc7014c57a143319e27.png)
![](https://pic1.zhimg.com/80/v2-367f61e5a35ae8e4b6c78ffc608463e2.png)
![](https://pic3.zhimg.com/v2-5a71f6795e0975a235d0edcbdc714650.jpg)

# RESTful API 说明
## 管理后台
- 产品列表：`GET /v1/products[?offset=0&limit=10&only=id]`
- 新建产品：`POST /v1/products`
- 查看产品：`GET /v1/products/:product_id`
- 更新产品：`PUT /v1/products/:product_id`
- 删除产品：`DELETE /v1/products/:product_id`

## 前台
- 公司列表：`GET /v1/companies`

# 主要依赖包

- [godotenv](https://github.com/joho/godotenv): 加载 `.env` 配置到环境变量
- [gorm](https://github.com/jinzhu/gorm): 用起来还不错的 ORM 框架
- [cast](https://github.com/spf13/cast): 类型转换帮助工具
- [gin](https://github.com/gin-gonic/gin): HTTP Web 框架，但是在该项目中进行了改造，参见 `util/rest`，强制用户返回 `error`
- [logrus](https://github.com/sirupsen/logrus): 日志工具包
- [testify](https://github.com/stretchr/testify): 单元测试断言框架
- [portal](https://github.com/ifaceless/portal): 偷懒神器，Model 与 Schema 映射工具，核心思想源自 Python [marshmallow](https://github.com/marshmallow-code/marshmallow)，对于一般的场景，尤其是管理后台接口开发场景下，可以提升开发效率，避免繁琐的字段类型转换等机械式编码
- [fixture](https://github.com/ifaceless/fixture): 测试数据管理工具

# 反馈

项目仅供参考，如有任何问题，欢迎在 issue 提出，谢谢~

