# Iterator pattern

## 概要
Iterator patternは、反復処理の依存性を切り離すデザインパータン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/iterator/diagram/abstract.svg" />

* Iterator
  * 要素を順番にスキャンする役のインタフェース
* Aggregate
  * Iterator を作り出す役のインタフェース
  * 自身の持っている要素のスキャン処理を Iterator 役に委譲する
* ConcreteIterator
  * Iterator の具体的な実装
* ConcreteAggregate
  * Aggregate の具体的な実装

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ iterator pattern
```

#### ざっくり仕様
* 本棚に本を追加できる
* 本棚にある本を羅列できる

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/iterator/diagram/impl.svg" />

インタフェースの定義は以下である。

Aggregator は、 Iterator インタフェースを実装した構造体を返すメソッドを持つ。

```go
type Iterator[T any] interface {
	HasNext bool
	Next (T, error)
}

// go の命名規則により Aggregate ではなく Aggregator としている
type Aggregator[T any] interface {
	Iterator Iterator[T]
}
```

Aggregator を実装した構造体が BookShelf(ConcreteAggregate) である。

フィールドに Book という配列をもち、この配列が反復処理される。Iterator メソッドでは、Iterator インタフェースを実装し、処理ごとに Book を返す BookShelfIterator(ConcreteIterator) を作成する。

```go
type Book struct {
	name string
}

type BookShelf struct {
	books []*Book
	last  int
}

func (bs *BookShelf) Iterator() Iterator[*Book] {
	iterator := NewBookShelfIterator(bs)
	return &iterator
}
```

上記を使用したサンプルコードが以下である。
* NewBookShelf で BookShelf を初期化
* Book を BookShelf に登録
* BookShelf からイテレーターを生成
* 反復処理

```go
func main() {
	var ag Aggregator[*Book]
	bs := NewBookShelf()
	ag = &bs

	titles := [...]string{"A", "B", "C", "D"}
	for _, title := range titles {
		book := NewBook(title)
		bs.AppendBook(&book)
	}

	it := ag.Iterator()
	for it.HasNext() {
		book, err := it.Next()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
		fmt.Printf("%s\n", book.Name())
	}
}
```


## 使えそうなユースケース
① データ構造が変わる可能性がある
* データの形が変われば、反復処理の実装が変わる
  * 例えば、データ構造が slice -> map に変化する
  * 当然データのデータの取り出し方が変わり Iterator 内の処理に修正が入る
* Iterator の依存性を切り離しているので、上流の修正は必要ない
* 異なるデータ構造でも、同じようにイテレートできる（ポリモーフィズム）

② 反復処理のロジックを切り替えたい
* 例えば、ConcreteAggreage のフィールドによって、反復処理を変える
  * sort field
    * asc : asc  iterator 使う
    * desc: desc iterator 使う
* ロジックの切り替えがなく、単純に複雑性を取り除きたい場合は、関数化でよさそう

## 感想
Iterator pattern の肝は、Iterator の依存性を、利用者である Aggregate から切り離したことだと思う。その結果、反復処理するデータの構造・反復処理ロジックへの変更が、上位の関数である main 関数に波及しない実装となった。

Iterator に限らず、以下のような場合は、依存性を切り離すことで保守性の高いプログラムの実現が可能だと思った。
* データ構造の変化が予想される
* ロジックの切り替えをしたい
* 振る舞い（インタフェース）に変化なし

こう言った場合は、具体的な実装ではなく、抽象的なインタフェースに依存したコードにしてやると良さそう。（Dependency Inversion）

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 

## 関連するデザインパターンとの違い
