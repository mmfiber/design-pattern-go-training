# Meditator pattern

## 概要
Meditator pattern はコンポーネント間の複雑に絡み合った依存性を削減するデザインパターン。独立したコンポーネント間でのやりとりを禁止し、Meditator と呼ばれるクラスを介して間接的にやり取りをする。全てのコンポーネントは Meditator にのみ依存する。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/meditator/diagram/abstract.svg" />

* 抽象的な説明
* 登場人物の紹介

## サンプル実装

[過去に書いた記事を参考](https://qiita.com/masachoco/items/bac5bc9d3d8cca642e90#mdeitator-%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E3%82%92%E6%8E%A1%E7%94%A8%E3%81%97%E3%81%9F%E5%AE%9F%E8%A3%85%E4%BE%8B)

実装：https://github.com/mmikami-lifull/meditator-pattern-vue/tree/main


## 使えそうなユースケース
① フロントエンドの実装

## 感想
インタラクティブなインタフェースに対するデザインパターンという印象が強い。いわゆるスパゲッティーコードのリファクタリングに使えそう。また、Meditator にロジックが集約するため、Meditator の実装を読むだけで、表示以外の仕様を把握できることもメリットだと思う。

サンプル実装もそうだが、依存性を一方向に限定する代わりに Meditator と Colleague との値のバケツリレーが面倒だなと思う。Vue に限った話だが、Vue3 の CompositionAPI を使えばバケツリレーが不要になり、Meditator にロジックを集約しつつ、煩雑なバケツリレーが不要になりそう。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/mediator

## 関連するデザインパターンとの違い
