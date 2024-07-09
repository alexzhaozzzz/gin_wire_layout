### github.com/alexzhaozzzz/gin_wire_layout

布局参考[project-layout](https://github.com/golang-standards/project-layout)，该项目非Go官方标准，但是已经是行业主流。

### 环境配置
1. Google wire
    ```shell
    $ go install github.com/google/wire/cmd/wire@latest
    ```
2. go模块
    ```shell
    $ go env -w GO111MODULE=on
    $ go env -w GOPROXY=https://goproxy.cn,direct
    ```
3. 参考: [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
