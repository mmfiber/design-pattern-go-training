# Decorator pattern

## 概要
Decorater pattern はラッパーとなるオブジェクトの中にオブジェクトを配置することで、新たな振る舞い（構造）を与えるデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/decorator/diagram/abstract.svg" />

* Componet
  * 機能追加する際の振る舞いを定めたインタフェース
* ConcreteComponent
  * Component を実装したクラス
  * このクラスに Decorator を使って装飾していく
* Decorator
  * Componet と同じ振る舞いを持つ
  * 装飾する Componet をプロパティーに持つ
* ConcreteDecorator
  * Decorator の実装

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ decorator pattern
```

#### ざっくり仕様
* 文字列を装飾する
* 基本となる装飾方法は2パターン
  * 文字列の両サイドに好きな文字を装飾できる
  * 文字列を囲うように装飾ができる
* 基本となる装飾パターンを自由に組み合わせて文字列の装飾ができる

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/decorator/diagram/impl.svg" />

それぞれの型定義。
```go
type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) (string, bool)
}

type StringDisplay struct {
	value string
}

type Border struct {
	display Display
}

type SideBorder struct {
	*Border
	borderChar string
}

type FullBorder struct {
	*Border
}
```

main 関数を見てみる。

オブジェクトを何層にもラップして新たに振る舞い（構造）を与えているのがわかる。
```go
func main() {
	show := func(display Display) {
		for i := 0; i < display.GetRows(); i++ {
			if text, ok := display.GetRowText(i); ok {
				fmt.Println(text)
			}
		}
	}

	b1 := NewStringDisplay("Hello, world.")

	(...)

	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewFullBorder(b1),
					"*",
				),
			),
		),
		"/",
	)
	show(b4)
	// ↓ 実行結果
	// /+-------------------+/
	// /|+-----------------+|/
	// /||*+-------------+*||/
	// /||*|Hello, world.|*||/
	// /||*+-------------+*||/
	// /|+-----------------+|/
	// /+-------------------+/
}
```

## 使えそうなユースケース
① フロントでユーザーにコンポーネントを作らせる
* decorater pattern で実装したコンポーネントをユーザーに提供する
* ユーザーはそのコンポーネントを動的に自在に組み合わすことができる
* ユーザーが取りうる全てのパターンを網羅したサブクラスを作る必要がない（というより、不可能だと思う）


## 感想
やろうと思えば継承でも同じようなことができる。しかし、以下のような場合では継承の実装が難しい。

* サンプル実装のように、デコレーターを何層にも重ねたい場合
  * デコレーターを重ねるたびにサブクラスを作っていると、夥しい数のサブクラスができるため
* 動的にデコレーターをアタッチしたい場合
  * そもそも、作るべきサブクラスが予想できない
  * 予想できても、上記のようにパターンが多いと管理が辛い

こういった場合に、再起的に結果をラップし、機能を追加してくれるデコレーターパターンが便利。何層にもデコレートしたオブジェクトをコピーできる、Prototype pattern と相性良さそう。逆に言うと、上記以外の場合は、サブクラスを作る継承を使えば良いと思う。

レイヤーの深さ と ConcreteDecorator の数には注意が必要だと思った。

レイヤーが深くなると、
* 呼び出し回数が指数関数的に増えていく可能性がある
* 再起的に処理を追っていくのが辛い

ので、max 3層ぐらいで収めると良いのではと思った。（勘）

ConcreteDecorator の数が増えると、細々としたクラスが増えて可読性が落ちそう。

使い所は考える必要がある。


## 関連するデザインパターンとの違い
### Adapter pattern
Adapter pattern は二つの異なるインタフェースのずれを吸収するデザインパターン。ズレを吸収するために、インタフェースを変更する。

Decorater pattern はラッパーとなるオブジェクトによって、新たな振る舞いを与えるデザインパターン。インタフェースは変更されず、追加される。

### Composite pattern
Composite pattern と Decorater pattern は構造的にとても似ているが、目的が異なる。

Composite pattern は容器と中身を同一視して、再起的な構造を作るデザインパターン。再起構造を作ることが目的。

Decorater pattern はラッパーとなるオブジェクトによって、新たな振る舞いを与えるデザインパターン。外枠を重ねることで、機能を使いしていくことに目的がある。

### Strategy pattern
Strategy pattern は実行時に使うアルゴリズムを切り替える。

Decorater pattern は実行結果に、ラッパーオブジェクトを重ねることで、機能を使いしていく。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/decorator
