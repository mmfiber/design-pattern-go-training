# Template method pattern

## 概要
抽象クラス（親クラス）でアルゴリズムを共通化し、具象クラス（子クラス）で具体的な実装をするデザインパターン。方針と詳細を分けている。

* 抽象クラスを使って継承して実装するのが Template method pattern。
* インタフェースと委譲を使って実装するのが Strategy pattern。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/templatemethod/diagram/abstract.svg" />

* Abstract
  * 抽象クラス
  * 共通のロジック（方針）を実装
* Concrete
  * 具象クラス
  * 具体的な（詳細）を実装

## サンプル実装
~~go では抽象クラスを表現できないので、~~（[embedding](https://go.dev/doc/effective_go#embedding)とレシーバーを持つメソッドを使えば表現できる）
委譲を使ったサンプル実装は Starategy pattern に譲る。

## 使えそうなユースケース
① フロントのコンポーネント
* Vuetify のようなフレームワークはこの考えを使っていそう
* 例えば、[Vue slot](https://v2.ja.vuejs.org/v2/guide/components-slots) を使って、[text-filed コンポーネントラベルをカスタマイズ](https://vuetifyjs.com/en/components/text-fields/#label)している
  * 以下のように依存関係が分離されている
    * 方針: text-field と言うコンポーネントがラベルを持つという方針を Vuetify がコンポーネント（Abstract Class）として定義
    * 詳細: ラベルがどんな dom（Concrete Class）を持つかという詳細はユーザーがカスタマイズできる
  * 結果的にカスタマイズできて、汎用性が高いコンポーネントになっている


② ロジックの抽象化
* 例えば、自転車と車という例で考える
  * Abstract Class: Vehicle
  * Concrete Class, Bycycle, Car
* どちらの Concrete Class もエネルギーを受け取り、走るという方針は共通である
* しかし、エネルギーの供給の仕方が異なる
  * Bycycle: 自転車を漕ぐ
  * Car    : アクセルを踏む
* getEnergy みたいな詳細なメソッドだけ、 Concrete Class に定義してやればよさそう

## 感想
Template method pattern は抽象クラスを使った実装方法。大切なのは、共通のロジック（方針）と具体的な要件（詳細）を分離すること。分離ができたら、それぞれを Abstract Class と Concrete Class に分けて実装してやればいい。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 

## 関連するデザインパターンとの違い
