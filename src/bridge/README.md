# Bridge pattern

## 概要
巨大なクラスや密接に関連したクラスの集まりを、 抽象と実装の二つの階層に分離し、それぞれが独立して開発できるようにしたデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/bridge/diagram/abstract.svg" />

* Abstraction
  * 抽象のクラス階層の最上位クラス **（抽象クラスではない）**
* RefindeAbstraction
  * Abstractionに機能追加をしたクラス
  * 必須ではない
* Implementor
  * 実装クラスの最上位クラス **（インタフェース）**
* ConcreteImplementor
  * Implementor を実装したクラス

抽象（Abstraction）と実装（Implementor）を繋ぐ Bridege の部分は継承で表現されている。

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ bridge pattern
```

#### ざっくり仕様
* プリント可能なコンピュータがある
* 各コンピューターはプリンターを選べる
* プリンターはファイルをプリントできる

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/bridge/diagram/impl.svg" />

AbstractPrintableComputer（Abstraction）と Printer(Implementation)
```go
type AbstractPrintableComputer struct {
	printer Printer
}

type Printer interface {
	PrintFile()
}
```

Windows, Mac（RefinedAbstraction）と Epson, Hp(ConcreteImplementation)
```go
type Windows struct {
	*AbstractPrintableComputer
}

type Mac struct {
	*AbstractPrintableComputer
}

type Epson struct{}

type Hp struct{}
```

main の実行を見てみる。

PrintableComputer(RefinedAbstraction) に Epson(ConcretePrinter) を渡して実行している。AbstractPrintableComputer（Abstraction）と Printer（Implementation）で、委譲を使った関係性の橋渡しがされているので、 windows/mac それぞれで epson/hp それぞれのプリンタを利用することができる。
```go
// RefinedAbstraction 用インフェース
type PrintableComputer interface {
	Print()
	SetPrinter(Printer)
}

func main() {
	var printableComputer PrintableComputer...

	...

	printableComputer.SetPrinter(&Epson{})
	printableComputer.Print()

	...

}
```

## 使えそうなユースケース
① 似たような機能のサブクラスが乱立した時
* [refacotoring guru](https://refactoring.guru/ja/design-patterns/bridge)で紹介されている例を使う
* 形と色に関心があるクラスがある
  * 形: 丸、三角、四角、、、（N個）
  * 色: 赤、青、黄色、、、（N個）
* 形と色の組み合わせごとにサブクラスを作っていたらNの2乗個サブクラスができる
* 形を抽象、色を実装と分けてあげると必要なクラスは2N個になる
* 似たようなサブクラスが乱立した時は検討してみるとよさそう

## 感想
似たようなサブクラスができてから検討するとよさそう。いきなりやるには、無駄に抽象度を上げてしまうのでよくない。複数の機能を持った巨大なクラスの分割にも効果的、単一責任の原則を満たすように変更できる。

## 関連するデザインパターンとの違い
### Starategy pattern
構造的には非常に Startegy pattern と似ている。RefinedAbstraction がなければ同じ構造。でも、Bridge pattern と Startegy pattern は決して同じデザインパターンではない。それは、目的が異なるから。
* Bridge pattern は抽象と実装を分離しそれぞれを独立して開発することが目的
* Strategy pattern は振る舞いを利用時に変更可能にすることが目的

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/bridge
