# bulletin-board-sample

クリーンアーキテクチャについて解説するためのアプリケーション

メッセージの保存、取得をすることができる

# Usage

## Get Messages

```bash
$ curl http://localhost:1323/messages | jq
{
  "Len": 5,
  "Messages": [
    "first message",
    "second message",
    "third message",
    "append message",
    "append message"
  ]
}
```

## Post Message

```bash
$ curl -X POST http://localhost:1323/message -H 'Content-Type: application/json' -d '{"message": "append message"}'
```

# Installation

## Used text and echo

```bash
$ https://github.com/higuruchi/bulletin-board-sample.git
$ cd bulletin-board-sample
$ go build cmd/text_echo/main.go
```

## Used MySQL and echo

```bash
$ https://github.com/higuruchi/bulletin-board-sample.git
$ cd bulletin-board-sample
$ go build cmd/mysql_echo/main.go
```