# Builder pattern

## 概要
複雑なオブジェクトを段階的に構築するデザインパターン。

Director が Builder インタフェースを介して、ConcreteBuilder の状態を変更する。Director が ConcreteBuilder のための手続的な処理（アルゴリズム）をカプセル化する。最終的な結果を ConcreteBuilder から取得する。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/builder/diagram/abstract.svg" />

* Builder
  * インスタンス（構造体）生成のためのインタフェース
* ConcreteBuilder
  * Builder インタフェースを実装したインスタンス（構造体）
  * コンストラクト結果を受け取るメソッドを持っている
* Director
  * Builder インタフェースに従ってコンストラクト処理をする役
  * Client にコンストラクト処理をかけるので必須ではないが、再利用の観点である方が望ましい
* Client
  * Director と ConcreteBuilder を使って生成したインスタンスを受け取る役

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ builder pattern
```

#### ざっくり仕様
* 文章をデコレートしてフォーマットする
* フォーマット形式は text, html
* フォーマットした文章をそれぞれ表示

#### 実装解説

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/builder/diagram/impl.svg" />

Builder(Builder) と TextBuilder・HTMLBuilder(ConcreteBuilder) と Director(Director) の定義。
```go
type Builder interface {
	makeTitle(string)
	makeString(string)
	makeItems([]string)
	build()
}

type TextBuilder struct {
	sb *strings.Builder
}

type HTMLBuilder struct {
	sb *strings.Builder
}

type Director struct {
	builder Builder
}
```

Director に注目する。

Director は Builder を具体的な実装（詳細）を知らなくて良い。Builder インタフェースに従った段階的な操作・アルゴリズム（方針）のみを定義している。

```go
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.makeTitle("Greeting")
	d.builder.makeString("朝から昼にかけて")
	d.builder.makeItems([]string{
		"おはようございます。",
		"こんにちは。",
	})
	d.builder.makeString("夜に")
	d.builder.makeItems([]string{
		"こんばんは。",
		"おやすみなさい。",
		"さようなら。",
	})
	d.builder.build()
}
```

Director が 抽象的な Builder に依存しているからこそ、具体的な ConcreteBuilder の付け替えが可能。

```go
func main() {
	director := NewDirector()

	textBuilder := NewTextBuilder()
	director.SetBuilder(&textBuilder) // ConcreteBuilder 付け替え
	director.Construct()
	fmt.Println(textBuilder.GetResult())

	htmlBuilder := NewHTMLBuilder()
	director.SetBuilder(&htmlBuilder) // ConcreteBuilder 付け替え
	director.Construct()
	fmt.Println(htmlBuilder.GetResult())
}
```

実行結果を見ると、同じ Director から異なるアウトプットを得ていることがわかる。

```
====================
『 Greeting 』

■ 朝から昼にかけて

        ・ おはようございます。
        ・ こんにちは。
■ 夜に

        ・ こんばんは。
        ・ おやすみなさい。
        ・ さようなら。

====================

<html>
<body>
<h1>Greeting</h1>
<p>朝から昼にかけて</p>
<ul>
<li>おはようございます。</li>
<li>こんにちは。</li>
</ul>
<p>夜に</p>
<ul>
<li>こんばんは。</li>
<li>おやすみなさい。</li>
<li>さようなら。</li>
</ul>
</html>
</body>
```

Builder pattern では ConcreteBuilder を Director に参照渡することが肝。最終的な結果を Director ではなく ConcreteBuilder から生成するため。Director から生成しない理由は、返り値の型を Director で縛ってしまうと使える状況が限定されるから。Builder に getResult という振る舞いを定義しないのはそのため。ただ、この戻り値に関してだけ言えば、 generics を使えば解決するはず。

ConcreteBuilder を参照渡しするメリットは、上記の Director の代わりに結果を生成できること以外にもある。それは、ConcreteBuilder が持っている状態を使って、 Builder には定義されていない独自のメソッドを実行できることである。

## 使えそうなユースケース
① 同じデータから異なるアウトプットを作る
* 分析結果を、text、excel, csv に吐き出せるようにする

## 感想
Director が抽象に依存しているので(Dependecy Inversion)、Builder の付け替えが可能。また、Director がコンストラクト処理をカプセル化してくれるので、再利用しやすい。

一方で、使う場所を考えないと無駄な知識の分離になる思った。Builder pattern は同じアルゴリズムを使うが、アウトプットの形式が異なる場合に有効だと思う。[過度な抽象化](https://www.linkedin.com/advice/0/how-do-you-balance-between-following-dry-principle)に気をつけないといけない。どのデザインパターンでも言えるが、アルゴリズムの共通化・抽象化がとても難しい。本当に共通化や抽象化が必要になってから分離しても良いのかもしれない。

## 関連するデザインパターンとの違い
### Template method pattern
Template method pattern のスーパークラス（抽象クラス）で実現していたことを、Director と Builder（インタフェース）を使って表現している。 

### Factory method pattern
Factory method pattern とも似ている。
| Factory method | Builder | 
| ------------ |--------------- |
| <img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/factorymethod/diagram/abstract.svg" /> | <img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/builder/diagram/abstract.svg" /> |
| <ul><li>インスタンスを生成する</li><li>抽象クラス(Creator)でロジックを抽象化</li><li>抽象クラスから(Creator)インスタンスを生成するため、生成したクラスが抽象的な型情報(Product)しか持たない（generics で解決する）</li></ul> | <ul><li>生成するものは決まっていない</li><li>コンストラク処理担当クラス(Director)とインタフェース(Builder)でロジックを抽象化</li><li>具象クラス(ConcreteBuilder)から結果を生成するため、生成した結果に具体的な型情報を待たすことができる</li><li>具象クラス(ConcreteBuilder)が状態を持っている</li></ul>

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/builder
