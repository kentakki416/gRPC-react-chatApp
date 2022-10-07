# gRPC-react-chatApp
 
gRPCを活用したチャットアプリ

※ブラウザから直接gRPCと通信することはできないので、gRPC-Webの仕組み使ってEnvoyプロキシ越しに通信を行う
（メッセージの取得はサーバーストリーミング
メッセージの送信はUnary RPC）

- CreateMessageは、フロントエンドからバックエンドへのメッセージ送信に用いる。
- GetMessageStreamはフロントエンドがバックエンドからメッセージを受信するのに使う  
※google.protobuf.Empty(void型)を受け取り、MessageのStreamを返す。その後は、MessageがどこかでPostされるたびに、それをStreamで随時受け取る
 
# DEMO
 
"hoge"の魅力が直感的に伝えわるデモ動画や図解を載せる
 
# Features
 
"hoge"のセールスポイントや差別化などを説明する
 
# Requirement
 
"hoge"を動かすのに必要なライブラリなどを列挙する
 
* huga 3.5.2
* hogehuga 1.0.2
 
# Installation
 
コードの自動生成コマンド
 
```bash
# protocが入ったイメージをpullする
❯ docker pull namely/protoc-all
# バックエンドのコードを生成（Go）
# -v: defs（definitions）ディレクトリにカレントディレクトリをマウント
# -f: 対象のprotoファイルを指定
# -o: アウトプットフォルダを指定
# -l: 生成する言語を指定
❯ docker run -v $PWD:/defs namely/protoc-all -f protobuf/chat.proto -o ./server/pb -l go
# フロントエンドのコードを生成（TypeScript）
❯ docker run -v $PWD:/defs namely/protoc-all -f protobuf/chat.proto -o ./client/src/pb -l web
```
 
# Usage
 
DEMOの実行方法など、"hoge"の基本的な使い方を説明する
 
```bash
git clone https://github.com/hoge/~
cd examples
python demo.py
```
 
# Note
 
注意点などがあれば書く
 
# Author
 
作成情報を列挙する
 
* 作成者
* 所属
* E-mail
 
# License
ライセンスを明示する
 
"hoge" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 
社内向けなら社外秘であることを明示してる
 
"hoge" is Confidential.