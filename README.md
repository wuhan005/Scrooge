# 💰 Scrooge ![Go](https://github.com/wuhan005/Scrooge/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/wuhan005/Scrooge)](https://goreportcard.com/report/github.com/wuhan005/Scrooge)

Personal sponsorship site based on [PAYBOB](https://paybob.cn/). / 基于 [PAYBOB](https://paybob.cn/) 的个人赞助收款站点

> “不能再这样下去了——每分钟损失十亿美元！600 年后我非破产不可！” —— [史高治·麦克达克](https://zh.wikipedia.org/wiki/%E5%8F%B2%E9%AB%98%E6%B2%BB%C2%B7%E9%BA%A6%E5%85%8B%E8%80%81%E9%B8%AD)

<img src="assets/Scrooge.JPG" align="right" width="250px"/>

## Requirement

* [PAYBOB](https://paybob.cn/) Account
* [PostgreSQL](https://wiki.postgresql.org/wiki/Detailed_installation_guides) (v12 or higher)

## Start

### 1. Build frontend

```bash
git clone git@github.com:wuhan005/Scrooge.git

cd Scrooge/frontend

yarn install && yarn build
```

### 2. Build backend

```bash
git clone git@github.com:wuhan005/Scrooge.git

cd Scrooge

go build .
```

### 3. Create configuration file

Create your own `./conf/Scrooge.ini` file based on the example [here](./conf/scrooge.ini).

### 4. Set the environment variables

```
PGHOST=<PostgreSQL Host>
PGPORT=<PostgreSQL Port>
PGUSER=<PostgreSQL User>
PGPASSWORD=<PostgreSQL Password>
PGDATABASE=<PostgreSQL Database>
PGSSLMODE=disable

PAYBOB_MCHID=<PAYBOB Mchid>
PAYBOB_KEY=<PAYBOB Key>
```

### 5. Launch!

```bash
./Scrooge web
```

## License

MIT
