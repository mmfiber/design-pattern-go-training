# Factory method pattern

## 概要
Template method pattern をインスタンス生成に応用したもの。インスタンス生成のための共通のアルゴリズム（方針）と、具体的なインスタンス生成（詳細）を分けて実装するデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/factorymethod/diagram/abstract.svg" />

* Creator（方針）
  * インスタンス生成をする抽象クラス
  * 共通の処理と振る舞いを定めている
* Product（方針）
  * Creatorによって生成されるインスタンスの振る舞いを定義
  * インタフェース or 抽象クラス（下記の使い分けでいいんじゃないかなと思っている）
    * 共通処理あり: 抽象クラス
    * 共通処理なし: インタフェース
    * （Product と ConcreteProduct に親子関係がある前提）
* ConcreteCreator（詳細）
  * Creator の実装
  * ある製品（ConcreteProduct）の生成に必要な具体的な処理を定義（詳細）
* ConcreteProduct（詳細）
  * Product の実装
  * 製品の具体的な振る舞いをメソッドを実装することで定義

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ factory method pattern
```

#### ざっくり仕様
* 以下の流れのでプロダクト作る工場を作成
  * プロダクトの作成
  * プロダクト所有者の記録
* IDCard というプロダクトを作る工場を作成
* 将来別の同じ作成工程を持ち、別のプロダクトを作る工場を作成するかも

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/factorymethod/diagram/impl.svg" />

Factory(Creator) と Product(Product) を定義

```go
type Product interface {
	Use()
}

// go にはオブジェクト思考の概念がないため、
// 抽象クラスを embedding + reciever を使って表現
// 振る舞い部分のみインタフェース（ProductCreator）を定義

type ProductCreator[T Product] interface {
	createProduct(owner Owner) T
	registerProduct(product T)
}

type Factory[T Product] struct {
	ProductCreator[T]
}

func (f Factory[T]) Create(owner Owner) T {
	product := f.createProduct(owner)
	f.registerProduct(product)
	return product
}
```

Factory に必要な ProductCreator と IDCard(ConcreteProduct)
 を実装
```go
type Owner struct {
	name string
}

type IDCard struct {
	owner *Owner
}

// ProductCreator を実装している
type IDCardCreator struct {
	owners []*Owner
}

func (f *IDCardCreator) createProduct(owner Owner) *IDCard {
	idCard := NewIDCard(owner)
	return &idCard
}

func (f *IDCardCreator) registerProduct(product *IDCard) {
	f.owners = append(f.owners, product.owner)
}
```

IDCard を生成するファクトリー（ConcreteFactory）を生成

```go
func main() {
	idCardCreator := NewIDCardCreator()
	idCardFactory := NewFactory[*IDCard](&idCardCreator) // Factory[*IDCard] 型

  // （省略）

}
```

例えば、テレビを作る工場が欲しいとなったら、以下のようにしてやれば良い。
プロダクトを生成して、オーナーを記録という処理は、抽象化した Factory がやってくれるので再度実装する必要はない。
```go
func main() {
	televisionCreator := NewTelevisionCreator()
	televisionFactory := NewFactory[*Television](&televisionCreator) // Factory[*Television] 型

  // （省略）

}
```


## 使えそうなユースケース
① ビジネス要件がわかっている時
* 将来こうなったら、こう事業展開してきます、みたいなのがわかっている時
* 将来的な変更が十分想定できるので、それに耐えられるようにしておこうとなる
* 上記のような理由がない限り、インスタンスの生成が複雑になってしまうこのデザインパターンは受け入れてもらえないと思う

## 感想
インスタンス生成に手続的な処理が必要で、将来的に生成するインスタンスの種類が増えても耐えうる設計にしたい時に便利。将来的に共通の性質を持つインスタンスが現れるかもわからない状況で、このデザインパターンを採用するのは yagni の原則に反すると思う。そもそもどこを共通化するべきかわからない。すでに「将来的に共通の性質を持つインスタンスを作るという」ことが想定されている場合のみ、当てはめるデザインパターンだと思う。

生成するインスタンスの種類の追加に対しては開いてて、修正に対しては閉じている(Open Closed Principle)。ロジックの共通化・カプセル化もできており、インスタンスの生成方法もポリモーフィズムも満たしている。不確定な将来の変更に対しても十分耐えうる設計になると思う。

ただ、何を共通化するかがとても大切。ここを間違えると、このデザインパターンは破綻する。再利用されずに、複雑にインスタンを生成するだけのコードに成り下がってしまう。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社
* https://refactoring.guru/ja/design-patterns/factory-method

## 関連するデザインパターンとの違い
