**简介**

类似linux的alias命令，用来简化命令调用。

**如何使用**

1. 初始化配置文件

```
xwc init
```

2. 初始化配置文件以后，会在项目根目录生成xwc.yml, 具体文件内容如下：

```
enviroment:

- your_enviroment: your_enviroment # 环境变量

command:

  your_command: your_command # 需要简写的命令

```

3. 配置文件完整案例

```
enviroment:
  - UC_MIGRATE_DATABASE: mysql://user:password@tcp(host:port)/dbname?x-migrations-table=migrations
command:
  http: gin -p 8080 -a 8080 --immediate -b uc run http
  doc: swag init
  migrate-create : migrate -database $UC_MIGRATE_DATABASE create -ext sql -dir migrations
  migrate: migrate -database $UC_MIGRATE_DATABASE -path migrations

```

4. 运行 `xic -h`

```
Usage:
  xwc [command]

Available Commands:
  doc            Exec: swag init
  help           Help about any command
  http           Exec: gin -p 8080 -a 8080 --immediate -b uc run http
  init           Init config file
  migrate        Exec: migrate -database $UC_MIGRATE_DATABASE -path migrations
  migrate-create Exec: migrate -database $UC_MIGRATE_DATABASE create -ext sql -dir migrations

Flags:
  -h, --help   help for xwc

Use "xwc [command] --help" for more information about a command.
```

5. `xwc http` 会执行 `gin -p 8080 -a 8080 --immediate -b uc run http`