# Visitor pattern

## 概要
Visitor pattern は、データ構造と処理を分離するデザインパターン。データ構造に手を加えられない、違和感がある時に使われる。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/visitor/diagram/abstract.svg" />

* 抽象的な説明
* 登場人物の紹介

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ visitor pattern
```

#### ざっくり仕様
* 図形のデータクラスがある
* 図形のデータクラスでバグが発生する懸念があるので、極力処理を加えたくない
* 面積の算出処理を、データクラスから切り離して実装する

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/visitor/diagram/impl.svg" />

型定義
```go
// Visitor
type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

// ConcreteVisitor
type AreaCalculator struct {
	area float64
}

// Element
type Shape interface {
	GetType() string
	Accept(Visitor)
}

// ConcreteElement
type Square struct {
	side float64
}

// ConcreteElement
type Circle struct {
	radius float64
}

// ConcreteElement
type Rectangle struct {
	l float64
	b float64
}
```

Accept -> VisitXXX の二重ディスパッチによって（オーバーロードでも代替可能）、AreaCalculator（ConcreteVisitor）が具象クラス(ConcreteElement)を知っている。そのため、AreaCalculator は、具象クラス独自のフィールドにアクセスすることができる。その結果、正方形・円・長方形と面積の選出方法が異なる場合でも、同じインタフェースで期待する結果を得ることができる。
```go
func (a *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * s.side
}

func (a *AreaCalculator) VisitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
	a.area = s.radius * s.radius * math.Pi
}

func (a *AreaCalculator) VisitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
	a.area = s.l * s.b
}

func (a *AreaCalculator) CalculatedArea() {
	fmt.Printf("Calculated area is: %.2f\n", a.area)
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.Accept(areaCalculator)
	areaCalculator.CalculatedArea()

	circle.Accept(areaCalculator)
	areaCalculator.CalculatedArea()

	rectangle.Accept(areaCalculator)
	areaCalculator.CalculatedArea()

	// <-- Output -->
	// Calculating area for square
	// Calculated area is: 4.00
	// Calculating area for circle
	// Calculated area is: 28.27
	// Calculating area for rectangle
	// Calculated area is: 6.00
}
```

処理の追加はこんな感じ。

やりたい処理ごとに ConcreteVisitor を作る。
```go
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
```

## 使えそうなユースケース
① 何かしらの理由でデータ構造に手を入れることができない時
* 巨大なアプリケーション開発をしていると、アプリケーション内で担当の領域があったりする
* 担当外のデータを使いたい & データクラスに処理を加えたいとする
* すんなり受け入れてくれたら良いが、リスクを負いたくないので拒否されたりするかもしれない
* そういった時に、データと処理を分けてくれる visitor pattern よさそう
* データクラスは、visitor インタフェースを引数に持ち、内部で自身を引数として、visit method を呼ぶ accept 関数を定義してやるだけで良い
* あとは visitor インタフェースを実装した concrete visitor が好きなように処理してくれる

```go
func (d *Data) Accept(v Visitor) {
  v.VisitForHoge(d)
}

func (v *ConcreteVisitor) VisitForHoge(d *Data) {
  // Data クラスのフィールドを使った処理が可能
}
```

## 感想
以下のような状況で役に立ちそうと思った。

* 複雑なデータ構造を持ち、データクラスの変更が容易ではない
* アプリケーションの規模が大きくデータクラス変更が許可されない

まずは、オーバーロードの利用を検討してみると良いと思った。Accept -> Visit という二重ディスパッチの流れは、処理を複雑にするので可読性が下がる。[オーバーロードが意図したように機能しない場合](https://refactoring.guru/ja/design-patterns/visitor-double-dispatch)に、初めて二重ディスパッチを使った実装を試みるべき。

やりたい処理ごとに ConcreteVisitor ができるので、小さなクラスが乱立して可読性が下がりそう。ファイルの分割を上手いことやれば、対処できそうな問題ではあるが。

これは switch 文を置き換えられるデザインパターンかも、、、

と思ったが、そう単純な話ではない気がしている。考えきれていない。少なくとも、二重ディスパッチの複雑さと、ConcreteVisitor に処理が集約されるという制約があると思っている。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社
* https://refactoring.guru/ja/design-patterns/visitor

## 関連するデザインパターンとの違い
### Decorator pattern
* 振る舞いを与えるという点では似ていると思う
* 振る舞いを与える時に具象クラスを知っているかどうか
  * Decorator: 知らない
  * Visitor: 知っている、データ構造に依存しない
