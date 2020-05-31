# 概要
Goでクリーンアーキテクチャで実装するためのサンプルリポジトリ

# 準備
- ローカルでLambdaを動かすためにlocalstackをインストールする
    - インストール方法: https://github.com/localstack/localstack
- sqlite3のオンメモリ機能を使うためにgccを動かせるようにする
    - MACの場合: xcodeのインストール

# 動かし方
同じドメインロジックをフレームワークやデータストア変更しても動かせるようになっている
格フレームワークによって動かし方が異なる

## Lambda
- localstackを起動
- localstack上にLambdaを作成
    - `make build_and_create_lambda MAIN_PATH={動かしたいmainファイル}`
- localstack上のLambdaを動かす
    - `make invoke_lambda ZIP_PATH={動かしたいmainファイル}`

## gin
- ローカルでginサーバを起動する
    - `go run {動かしたいmainファイル}`

# ドメインロジックについて

## create user
- ユーザーを作成する

## get user
- 作成したユーザーを取得する