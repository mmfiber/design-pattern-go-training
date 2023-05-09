# Chain of responsibirity pattern

## 概要
共通の振る舞いをするハンドラーを、特定の順序で実行していくデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/chainofresponsibility/diagram/abstract.svg" />

* Handler
  * 要求を処理する振る舞いを定める役
  * 抽象クラスで定義して、メソッドを定義しても良い
* ConcreteHandler
  * Handler の実装
* Client
  * ConcreteHandler の利用者

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ chain of resposibility pattern
```

#### ざっくり仕様
* ハンドラーは次に呼ぶハンドラーを持つことができる
* ハンドラーは次に呼び出すハンドラー有無によって処理がことなる
  * あり: 次のハンドラーを呼ぶ
  * なし: 終了

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/chainofresponsibility/diagram/impl.svg" />

型定義。
```go
type Handler interface {
	next() Handler
	setNext(Handler)
	handle()
}

type ConcreteHandler struct {
	id      string
	handler Handler
}
```

ConcreteHandler には必要に応じて自身の次に呼ぶ Handler を setNext 関数によって定義できる。Handler の順序を制御している。

next関数があれば呼び、無ければ処理が終了する。

```go
func (h *ConcreteHandler) handle() {
	next := h.next()
	if next == nil {
		return
	}
	next.handle()
}

func main() {
	h1 := NewConcreteHandler("1")
	h2 := NewConcreteHandler("2")
	h3 := NewConcreteHandler("3")

	h1.setNext(h2)
	h2.setNext(h3)
	h1.handle()
	// <-- Output -->
	// Handler 1
	// Handler 2
	// Handler 3
}
```

## 使えそうなユースケース
① DOM
② [express.js](https://expressjs.com/ja/4x/api.html#app.use)

## 感想
直列に処理を実行できるので、状況に応じて利用者が処理を組み合わせて使うことができる。任意のタイミングで処理を差し込めるので、割と柔軟性が高い。そのため、express や next といったフレームワークにもこのデザインパターンが採用されているのかなと思う。

一方で、直列の処理なので、 Handler 内での時間のかかる処理（http リクエストなど）は向かない。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/chain-of-responsibility

## 関連するデザインパターンとの違い
