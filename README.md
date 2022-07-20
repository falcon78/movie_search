# 実行方法

本アプリケーションを実行するPCにDockerが入ってない場合はDockerをインストールしてください。

## 環境変数の設定

`.env`ファイルが存在しない場合は下記のコマンドで生成してください。

```
cp .example.env .env
```

## データベースを立ち上げる

最初にPostgresを下記のコマンドで立ち上げてください。
```
docker compose up movie_pg -d
```

## データダンプをインポートする
ユーザー名とパスワードは`.env`ファイルの中に書いてあります。
`.example.env`をコーピしてそのまま使っている場合はユーザー名が`movie_pg`でパスワードが`password`になります。
```bash
psql -h localhost -d postgres -U movie_pg -f pg_dump.sql
```

## アプリケーションを立ち上げる

下記のコマンドでアプリケーションを立ち上げてください。
```bash
docker compose up movie_go
```

アプリケーションは`localhost:31415`からアクセスできます。