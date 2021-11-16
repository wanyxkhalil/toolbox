# toolbox

toolbox: mkpasswd, https-expired, mysql-to-gostruct, ip, md-toc-github, pension ...

Table of Contents
=================

* [Install](#install)
    * [Completion](#completion)
* [Usage](#usage)
    * [mkpasswd](#mkpasswd)
    * [https-expired](#https-expired)
    * [mysql-to-gostruct](#mysql-to-gostruct)
    * [ip](#ip)
        * [Local machine address](#local-machine-address)
        * [Remote host address](#remote-host-address)
    * [md-toc-github](#md-toc-github)
    * [Pension](#pension)

## Install

```shell
go install github.com/wanyxkhalil/toolbox@latest
toolbox -h
```

### Completion

If you use zsh, add this to ~/.zshrc & source. Also support bash, fish, powershell

```shell
if [ $commands[toolbox] ]; then
	source <(toolbox completion zsh)
	compdef _toolbox toolbox
fi
```

## Usage

### mkpasswd

Generate password, Inspired by [gaiustech/MkPasswd](https://github.com/gaiustech/MkPasswd).

```shell
$ toolbox mkpasswd -h # help message
A tool for generating random passwords

Usage:
  toolbox mkpasswd [flags]

Flags:
  -d, --digit uint     Number of digits (default 2)
  -h, --help           help for mkpasswd
  -l, --length uint    Length in chars (default 9)
  -c, --lower uint     Number of lowercase chars (default 2)
  -s, --special uint   Number of special chars
  -C, --upper uint     Number of uppercase chars (default 2)
```

Sample

```shell
toolbox mkpasswd -l 17 # length is 17
toolbox mkpasswd -l 17 -C 4 -d 4 -s 3 # length is 17, include 4 upper char, 4 digit, 3 special char, 6 lower char
```

### https-expired

Show the cert expiration time. Just like

```shell
alias ,https-expired='function _as() {echo | openssl s_client -servername $1 -connect $1:443 2>/dev/null | openssl x509 -noout -dates;};_as'
```

Sample

```shell
toolbox https-expired github.com
```

### mysql-to-gostruct

类型对应主要使用 github.com/go-sql-driver/mysql 中的 fields.go。

| tinyint            | int8                |
| ------------------ | ------------------- |
| smallint           | int16               |
| mediumint          | int32               |
| int                | int                 |
| bigint             | int64               |
| tinyintunsigned    | uint8               |
| smallintunsigned   | uint16              |
| mediumintunsigned  | uint32              |
| intunsigned        | uint                |
| bigintunsigned     | uint64              |
|                    |                     |
| float              | float32             |
| double             | float64             |
| decimaldefaultnull | decimal.NullDecimal |
| decimalnotnull     | decimal.Decimal     |
|                    |                     |
| year               | uint8               |
| time               | time.Time           |
| date               | time.Time           |
| datetime           | time.Time           |
| timestamp          | time.Time           |
|                    |                     |
| char               | string              |
| varchar            | string              |
| tinytext           | string              |
| text               | string              |
| mediumtext         | string              |
| longtext           | string              |
| enum               | string              |
| set                | string              |
|                    |                     |
| bit                | byte                |
| binary             | []byte              |
| varbinary          | []byte              |
| tinyblob           | []byte              |
| blob               | []byte              |
| mediumblob         | []byte              |
| longblob           | []byte              |

> - year 对应 uint8，无法对应 time.Time。
> - decimal 对应 github.com/shopspring/decimal

使用示例：

```shell
toolbox mysql-to-gostruct /Tmp/table.sql /Tmp
```

> - table.sql：使用 `show create table xxx` 生成，支持多个表
> - /Tmp: 结果目录，目录名作为包名

Table.sql 示例：

```sql
CREATE TABLE `user`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255)   NOT NULL COMMENT '用户名',
    `valid`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有效：0_无效，1_有效',
    `dec`        decimal(10, 2)          DEFAULT NULL,
    `udec`       decimal(10, 2) NOT NULL,
    `created_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=17318 DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE `my_time`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `year` year DEFAULT NULL,
    `time`      time     DEFAULT NULL,
    `date`      date     DEFAULT NULL,
    `datetime`  datetime DEFAULT NULL,
    `timestamp` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

结果：/Tmp/user.go, /Tmp/my_time.go

```go
package tmp

import (
	"github.com/shopspring/decimal"
	"time"
)

// User 用户
type User struct {
	Id uint64
	// 用户名
	Name string
	// 是否有效：0_无效，1_有效
	Valid     int8
	Dec       decimal.NullDecimal
	Udec      decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

```go
package tmp

import (
	"time"
)

// MyTime
type MyTime struct {
	Id        uint64
	Year      uint8
	Time      time.Time
	Date      time.Time
	Datetime  time.Time
	Timestamp time.Time
}
```

### ip

Get IP address.

#### Local machine address

```shell
toolbox ip
```

Result:

```shell
IP Address:		180.167.000.000
Country:		China
Intranet:		192.168.3.238
```

#### Remote host address

```shell
toolbox ip
```

Result:

```shell
220.181.000.000
```

### md-toc-github

Generate Markdown TOC for GitHub

Sample:

```shell
## This line for toc test

toolbox md-toc-github README.md
```

Result:

```markdown

Table of Contents
=================

* [toolbox](#toolbox)
    * [Install](#Install)
        * [Completion](#Completion)
    * [Usage](#Usage)
        * [mkpasswd](#mkpasswd)
        * [https-expired](#https-expired)
        * [mysql-to-gostruct](#mysql-to-gostruct)
            * [类型对应](#类型对应)
            * [Sample](#Sample)
        * [ip](#ip)
            * [Local machine address](#Local machine address)
            * [Remote host address](#Remote host address)
        * [md-toc-github](#md-toc-github)

```

### Pension

计算基本养老金

```shell
toolbox pension --avgWages 10338 --wageRatio 0.6 --years 30 --personalWages 6000 --months 150
```
> 上年度月平均工资：10338
> 本人历年缴费指数的平均值：0.6
> 累计缴费：30年
> 本人缴费工资基数：6000
> 计发月数：150