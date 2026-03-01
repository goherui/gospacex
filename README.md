gospacex/
├── Pos-service/           # 职位服务（gRPC服务）
│   ├── basic/             # 基础配置
│   │   ├── cmd/           # 命令行入口
│   │   │   └── main.go    # 服务启动文件
│   │   ├── config/        # 配置文件
│   │   │   ├── config.go  # 配置结构
│   │   │   └── global.go  # 全局变量
│   │   └── init/          # 初始化
│   │       ├── init.go    # 初始化入口
│   │       ├── mysql.go   # MySQL初始化
│   │       └── viper.go   # 配置加载
│   ├── handler/           # 处理器
│   │   └── service/       # 服务实现
│   │       └── Pos.go     # 职位服务实现
│   └── model/             # 数据模型
│       └── Position.go    # 职位模型
├── bff/                   # Backend for Frontend
│   ├── basic/             # 基础配置
│   │   ├── cmd/           # 命令行入口
│   │   │   └── main.go    # 服务启动文件
│   │   ├── config/        # 配置文件
│   │   │   └── global.go  # 全局变量
│   │   └── init/          # 初始化
│   │       └── init.go    # 初始化入口
│   ├── handler/           # 处理器
│   │   ├── request/       # 请求结构
│   │   │   └── Pos.go     # 职位请求结构
│   │   ├── response/      # 响应结构
│   │   │   └── Pos.go     # 职位响应结构
│   │   └── service/       # 服务实现
│   │       └── Pos.go     # BFF层职位服务
│   └── router/            # 路由
│       └── router.go      # 路由配置
├── proto/                 # gRPC协议定义
│   ├── Position.pb.go     # 生成的Go代码
│   ├── Position.proto     # 协议定义文件
│   └── Position_grpc.pb.go # 生成的gRPC代码
├── .gitignore             # Git忽略文件
├── config.yaml            # 配置文件
├── go.mod                 # Go模块定义
└── go.sum                 # 依赖校验和

架构设计
分层架构
BFF层（Backend for Frontend）：

提供HTTP RESTful API接口
处理前端请求，调用gRPC服务
负责请求参数验证和响应格式化
gRPC服务层：

实现核心业务逻辑
提供gRPC接口供BFF层调用
与数据库交互
数据层：

基于GORM实现数据模型
与MySQL数据库交互
技术栈
技术/框架	版本	用途
Go	1.25.7	开发语言
gRPC	v1.79.1	服务间通信
Gin	v1.12.0	HTTP框架
GORM	v1.31.1	ORM框架
Viper	v1.21.0	配置管理
MySQL	-	数据库
启动流程
环境要求
Go 1.25.7+
MySQL 5.7+
gRPC 工具链

