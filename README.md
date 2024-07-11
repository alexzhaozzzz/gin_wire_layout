### github.com/alexzhaozzzz/gin_wire_layout

布局参考[project-layout](https://github.com/golang-standards/project-layout)，该项目非Go官方标准，但是已经是行业主流。

### 环境配置
* Google wire
    ```shell
    $ go install github.com/google/wire/cmd/wire@latest
    ```
* go模块
    ```shell
    $ go env -w GO111MODULE=on
    $ go env -w GOPROXY=https://goproxy.cn,direct
    ```
* 参考: [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* swag:
     ```
     go get -u github.com/swaggo/swag/cmd/swag
     // 引入依赖
     go get -u github.com/swaggo/gin-swagger
     go get -u github.com/swaggo/files
     ```

### 运行项目
* 初始化依赖
   ```
   $ go mod tidy
   ```
* 命令
   ```
   # 使用make
   # 打包（Linux/MacOS 下），在项目目录下执行make命令
   $ make 
   
   # 删除已打的包
   $ make clean
   ```

### 项目结构

   ```
   .
   ├── bin
   │   └── configs
   ├── cmd
   │   └── gin_wire_layout
   ├── configs
   ├── docs
   ├── internal
   │   ├── model
   │   │   ├── mysql
   │   │   └── po
   │   ├── repo
   │   ├── router
   │   └── service
   ├── pkg
   │   ├── bootstrap
   │   ├── colorx
   │   ├── connect
   │   │   ├── mysqlx
   │   │   └── redisx
   │   ├── jwt
   │   ├── security
   │   ├── serverx
   │   ├── util
   │   └── version
   └── scripts
   ```

