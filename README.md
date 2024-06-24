<div id="top"></div>

## 実行方法

とりあえずコピー






env という名前の仮想環境を作ります。
その次のコマンドは仮想環境を有効にするコマンドと無効にするコマンドです。

```
% python -m venv env
% source env/bin/activate
(env) % deactivate
```

Chrome Wevdriver を用意します。
```
(env) % python -m pip install webdriver-manager
```

Selenium Python の用意
```
(env) % python -m pip install selenium
```

実行結果です。
```
(env) % python syuwa_min.py                      
{'title': '図解入門 よくわかる最新物理化学の基本と仕組み （単行本）', 'price': '2200円\n（本体2000円＋税10%）', 'author': '齋藤\u3000勝裕', 'describe': 'あなたは「物理化学」と聞くと、メンドクサイ理屈をこねて中身は数式イッパイで嫌だなと感じるかもしれません。でも、それは間違いです。宇宙や原子のワクワクする話から化学反応の基礎を学び、使用する数式も難しい微分積分は一切なく、足し算・引き算・かけ算・割り算だけです。本書は、大学で学ぶ物理化学の基本と仕組みを図表を交えてわかりやすく解説した入門書です。各章末にはこれだけわかれば単位取得に近づく「演習問題」付き。'}

```

## 参考

サイドバイサイドにプレビューを表示 [Command] + [K] → [V]

https://github.com/miyamotok0105/crawling-sample/tree/main/ch02

参考：[全プロジェクトで重宝されるイケてるREADMEを作成しよう！](https://qiita.com/shun198/items/c983c713452c041ef787)

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <!-- バックエンドの言語一覧 -->
  <img src="https://img.shields.io/badge/-Python-F2C63C.svg?logo=python&style=for-the-badge">
</p>

## 目次

1. [プロジェクトについて](#プロジェクトについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [開発環境構築](#開発環境構築)
5. [トラブルシューティング](#トラブルシューティング)


## プロジェクト名

Pythonスクレイピング&クローリング収集マスタリングハンドブックの syuwa_min.py（改）

<!-- プロジェクトについて -->

## プロジェクトについて

本のままだとエラーが出たので改造しました。

## 環境

<!-- 言語、フレームワーク、ミドルウェア、インフラの一覧とバージョンを記載 -->

| 言語・フレームワーク  | バージョン |
| --------------------- | ---------- |
| Python                | 3.12.2     |

<p align="right">(<a href="#top">トップへ</a>)</p>

## ディレクトリ構成



<p align="right">(<a href="#top">トップへ</a>)</p>

## 開発環境構築


## トラブルシューティング


<p align="right">(<a href="#top">トップへ</a>)</p>