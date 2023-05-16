# Observer pattern

## 概要
観察対象のオブジェクトのイベントを購読するデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/observer/diagram/abstract.svg" />

* Subject
  * 状態の変化を追跡する被験者オブジェクト
  * 複数のオブザーバーを登録および削除するためのメソッドを提供する
* Observer
  * Subject の状態変化を受け取るオブジェクト
  * 状態の変化に対する通知を受け取るためのメソッドを提供する
* ConcreteSubject
  * Subject の実装
* ConcreteObserver
  * Obserber の実装

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ Observer pattern
```

#### ざっくり仕様
* pub, sub できる構造

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/observer/diagram/impl.svg" />

型定義
```go
type Subject interface {
	RegisterObserver(Observer)
	UnregisterObserver(Observer)
	NotifyObservers()
}

type Observer interface {
	Update(Subject)
}
```

Subject にイベントを購読する Observer を登録する。登録されていない Observer には Subject からイベントを受け取らない。
```go
func main() {
	obs1 := &ConcreteObserver{name: "obserber1"}
	obs2 := &ConcreteObserver{name: "obserber2"}

	sub := &ConcreteSubject{}
	sub.RegisterObserver(obs1)
	sub.RegisterObserver(obs2)

	sub.NotifyObservers()
	// <-- Output -->
	// Observer obserber1 received notification from Subject
	// Observer obserber2 received notification from Subject

	sub.UnregisterObserver(obs1)

	sub.NotifyObservers()
	// <-- Output -->
	// Observer obserber2 received notification from Subject
}
```

## 使えそうなユースケース
① [MutationObserver](https://developer.mozilla.org/ja/docs/Web/API/MutationObserver)

## 感想
pub, sub をするためのデザインパターンという印象。変更に対して reactive な実装をしたい時に有効だと思う。身近なところで言うと、ブラウザで動く UI フレームワークに採用されている。

Observer 自身が Subject の起点になる場合が想定される。その場合は、無限ループになる危険があるので、注意が必要。

## 関連するデザインパターンとの違い
### Meditator pattern 
Observer pattern の SUbject がイベントを発火して、Observer がイベントを受け取るという構造と、Meditator pattern の Meditator がイベントを発火して、Colleague がイベントを受け取るという構造が一致しているため、違いがわかりにくい。違いは、目的にある。
* Observer pattern : pub, sub の構造を作る
* Meditator pattern: 子コンポーネント同士の依存関係を排除する（脱スパゲッティーコード）

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/observer


