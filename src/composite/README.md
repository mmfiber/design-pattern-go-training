# Composite pattern

## 概要
モデルがツリー構造で表現できる場合のみに適用すべきデザインパターン。容器と中身を同一視して、再起的な構造を作る。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/composite/diagram/abstract.svg" />

* Leaf
  * 中身の役
* Composite
  * 箱の役
  * 箱の中身は Leaf か Composite どちらでもよい
* Component
  * Leaf、Componet を同一視するためのインタフェース
* Client
  * Composite pattern の利用者

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ composite pattern
```

#### ざっくり仕様
* ツリー構造の処理を再起的に実行できる

#### 実装解説

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/composite/diagram/impl.svg" />

```go
type Component interface {
	Operation() string
}

type Leaf struct {
	name string
}

type Composite struct {
	children []Component
	name     string
}
```

main 関数で実行されている Composite の Operation をみると、Composite が所持している childeren が Leaf か Composite かを判別せずに実行できているのがわかる。Component インタフェースを使って Leaf と Composite を同一視しているためである。
```go
func (c *Composite) Operation() string {
	var results []string
	for _, c := range c.children {
		results = append(results, c.Operation())
	}

	(...)
}

func main() {
	tree := &Composite{name: "Tree"}

	(...)

	fmt.Println(tree.Operation())

	(...)
}
```

## 使えそうなユースケース
① OSのファイルシステム
② DOM構造

## 感想
このデザインパターンの最大の利点は、ツリーを構成するオブジェクトの具象クラスを気にする必要がないこと。オブジェクトが Leaf または Composite なのか知る必要がない。子要素を持つのか、持っている子要素が Leaf または Composite なのか知らなくて良い。メソッドを呼び出すと、オブジェクト自身がリクエストをツリーの下方まで再起的に渡してくれるから。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/composite

## 関連するデザインパターンとの違い
### Iterator pattern
Iterator pattern を使って Composite pattern で作られたツリー構造を再起的に探索できる。

### Prototype pattern
Composite pattern で作られたオブジェクトをクローンするのに役立つ。
