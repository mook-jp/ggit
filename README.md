# Git-like App in Go

## Description
このアプリは `git` の 機能を学習するために `go` でクローンアプリを作成したものです。
現在作成済みのコマンドを下記する。
- `init`: `.mygit` ディレクトリを作成するコマンド
- `hash-object`: ファイルごとの `hash` を生成する
- `cat-file`: `<hash>` と与えられたオプションごとに処理を実行する
- `rev-parse`: レポジトリのルートディレクトリを取得する
- `write-tree`: レポジトリ内のツリー構造を取得して、`.mygit`に書き込む

## Usage
このアプリの使い方
後ほど記載する