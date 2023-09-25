# Remove reflection tool

![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

## 概要

ファイル整理ツールです。2つのディレクトリ内のファイルを比較し、基準となるディレクトリに存在しないファイル名(拡張子は除く)を持つファイルを対象ディレクトリから削除します。

主な使い道として、RAWとJPEGの入ったディレクトリがそれぞれあり、JPEGのディレクトリに入ったを整理した結果をRAWのディレクトリに反映させるといった状況を想定します。


## 使用法

```
remove-relection-tool 対象ディレクトリ 基準ディレクトリ
```

途中で確認が出ます。いきなり削除されることはありません。


## インストール

開発環境はGo1.21です。恐らくGo 1.18以降なら以前のGoでも動作するはずです。

go install で導入できます。

```
go install github.com/okashi-uji/photo-organization-tool@latest
```

## 作者

- おかし(@okashi-uji)
