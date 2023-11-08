# Remove reflection tool
![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

## 概要

ファイル整理ツールです。2つのディレクトリ内のファイルを比較し、基準となるディレクトリに存在しないファイル名(拡張子は除く)を持つファイルを対象ディレクトリから削除します。現在のバージョンではディレクトリ直下のファイルだけを対象としていて、再帰的に処理はしません。

主な使い道として、RAWとJPEGの入ったディレクトリがそれぞれあり、JPEGのディレクトリに入った画像を整理した結果をRAWのディレクトリに反映させるといった状況を想定しています。


## 使用法

```
remove-reflection-tool 対象ディレクトリ 基準ディレクトリ
```

途中で確認が出ます。いきなり削除されることはありません。


## インストール

開発環境はGo 1.21です。1.21かそれより新しいバージョンの環境で使用することを推奨しますが、恐らく1.18以降なら動作すると思います。

go install で導入できます。

```
go install github.com/oka4shi/remove-reflection-tool@latest
```

## 作者

- おかし(@okashi-uji)
