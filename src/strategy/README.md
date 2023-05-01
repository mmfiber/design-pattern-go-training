# Strategy pattern

## 概要

実行時にアルゴリズムの切り替えを委譲を使って行うデザインパターン。

共通の処理があって、一部のアルゴリズムのみ切り替えたい時に使われる。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/strategy/diagram/abstract.svg" />

* Strategy
  * 切り替えたいアルゴリズムの振る舞いを定義したインタフェース
* ConcreteStrategy
  * Strategy の実装
  * Context で切り替えて実行したい具体的なアルゴリズムを実装
* Context
  * Strategy を利用する人、クラスでも、メソッドでもよい
  * Strategy を使って、共通のロジックを実装
  * 何かしらの方法で Strategy に ConcreteStrategy を当てはめて使う
    * コンストラクタやメソッドの引数としてもらう
    * factory を実装して、メソッド内で切り替える

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ strategy pattern
```

#### ざっくり仕様
* 距離と時間を教えてくれる Navigator がある
* 歩きか公共交通機関かによって出力結果が変わる

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/strategy/diagram/impl.svg" />

Strategy インタフェースを満たした WalkingStrategy と PublicTransportStrategy （ConcreteStrategy）を付け替えることでアルゴリズムを変更。異なるアウトプットを得ることができる。
```go
func main() {
	var navigator *Navigator

	navigator = &Navigator{&WalkingStrategy{}}
	navigator.Route("Shibuya", "Yoyogi")
	navigator.Route("Yoyogi", "Shibuya")

	navigator = &Navigator{&PublicTransportStrategy{}}
	navigator.Route("Shibuya", "Yoyogi")
	navigator.Route("Yoyogi", "Shibuya")
}
```

```
Walking route from Shibuya to Yoyogi: 4 km, 30 min
Walking route from Yoyogi to Shibuya: 4 km, 30 min
Public transport route from Shibuya to Yoyogi: 3 km, 5 min
Public transport route from Yoyogi to Shibuya: 3 km, 5 min
```
## 使えそうなユースケース
① ログインユーザー、ゲストユーザーの切り替え
* Facade pattern と一緒に使うようなイメージ
* 利用者はシンプルな窓口でログインか、ゲストかを知らずに処理をできる
* 内部では Facade pattern でログインユーザー用アルゴリズムを使うか、ゲストユーザー用アルゴリズムを使うか振り分け
* 共通化できない具体的は処理はそれぞれのアルゴリズムに任せる

## 感想
クラス内でアルゴリズムによって処理を切り替えたい場所が多い場合に便利。

クラス内のあらゆるメソッドで if 文や switch 文でアルゴリズムを切り替えるのは大変だし、見通しが悪い。同じ条件分岐が乱立するため、条件を新たに追加したい際に、全ての条件分岐に追加修正が必要。結果、実装漏れにつながる。クラスとアルゴリズムが密結合な状態。

一方で Strategy pattern では、 interface で振る舞いを定義したアルゴリズムを任意のタイミングて切り替えることができる。そのため、if 文や switch 文による条件分岐が不要になる。アルゴリズムを追加したい際も、アルゴリズム自体に集中すれば良い。アルゴリズムの切り替え方法に注意を払う必要がない。interface を使って抽象に依存させることで、クラスとアルゴリズムが疎結合な状態になっている。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/strategy

## 関連するデザインパターンとの違い

### Template method pattern
Template method pattern も Starategy pattern も共通のロジックがあり、部分的にロジックが異なる際に有効なデザインパターン。どちらも、部分的に異なる具体的な処理を、共通のロジックから切り出して実装することを可能にする。しかし、アプローチ方法が異なる。

* Template method pattern
  * **継承**を使う
  * スーパークラスで共通のロジックを定義
  * スーパークラスでサブクラスに実装してほしい振る舞い（抽象メソッド）を定義
  * サブクラスで振る舞いを実装
* Strategy pattern
  * **委譲**を使う
  * コンテキスト（利用者、クラスだったり、メソッドだったり）で共通のロジックを定義
  * インタフェースで実装してほしい振る舞いを定義
  * インタフェースで定められた振る舞いを実装
