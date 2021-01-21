单机版本博客系统
------------
|  目录名 | 作 用|
| :------------: | :------------: |
| cmd  |  mian函数的位置 |
|internal|私有目录代码 |
|pkg|外部项目可以引入的代码 |
|api|OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件。 |
|web|特定于 Web 应用程序的组件:静态 Web 资产、服务器端模板和 SPAs。 |
|configs |配置文件 |
|init |初始化代码 |
|scripts |脚本代码 |
|deploy  |IaaS、PaaS、系统和容器编配部署配置和模板(docker-compose、kubernetes/helm、mesos、terraform、bosh)。注意，在一些存储库中(特别是使用 kubernetes 部署的应用程序)，这个目录被称为 |
|test |额外的外部测试应用程序和测试数据。你可以随时根据需求构造 /test 目录。对于较大的项目，有一个数据子目录是有意义的。例如，你可以使用 /test/data 或 /test/testdata (如果你需要忽略目录中的内容)。请注意，Go 还会忽略以“.”或“_”开头的目录或文件，因此在如何命名测试数据目录方面有更大的灵活性|
|docs  |文档 |
|controller  |控制层 |
|service  | |
|dao  |数据交互 |
|entity  |实体层 |
|global  |全局变量 |