# Interpreter pattern

## 概要
`<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/interpreter/diagram/abstract.svg" />

* Expression
  * 構文木ノードのインタフェース
* TerminalExpression
  * Expression の実装
  * 終端となる表現
* NonTerminalExpression
  * Expression の実装
  * 終端となる表現
* Context
  * インタプリタが構文解析を行うための情報
* Client
  * TerminalExpression や NonTerminalExpression を使って構文木を組み立てる

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ interpreter pattern
```

#### ざっくり仕様
* 文字列のリスト（トークン）を解析して Expressionを生成
* `AND` または `OR` トークンを見つけると、その後の2つのトークンをそれぞれ左側と右側のExpressionとして解析

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/interpreter/diagram/impl.svg" />

型定義
```go
type Expression interface {
	Interpret() bool
}

type TrueExpression struct{}


type FalseExpression struct{}


type AndExpression struct {
	left, right Expression
}

type OrExpression struct {
	left, right Expression
}

func main() {
	// Context
	tokens := strings.Split("AND TRUE OR FALSE TRUE", " ")
}
```

tokens(Context) を Parse メソッドで構文解析を行う。この時、AndExpression, OrExpression には Composite pattern を採用しており、ツリー構造を有している。`spew.Printf("%#v\n", expression)` で構文解析結果を示す。

Expression の Interpret メソッドを呼び出し、構文解析した内容を Expression の実装で定義したロジックに従って実行する。
```go
func main() {
	// This represents the rule "TRUE AND (FALSE OR TRUE)"
	tokens := strings.Split("AND TRUE OR FALSE TRUE", " ")
	expression := Parse(tokens)
	spew.Printf("%#v\n", expression)
	// Output:
	// (*interpreter.AndExpression){left:(*interpreter.TrueExpression){} right:(*interpreter.OrExpression){
	//	left:(*interpreter.FalseExpression){} right:(*interpreter.TrueExpression){}
	// }}

	fmt.Println(expression.Interpret())
	// Output: true
}
```

## 使えそうなユースケース
① SQL ② CLI

## 感想
[DSL](https://ja.wikipedia.org/wiki/%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E5%9B%BA%E6%9C%89%E8%A8%80%E8%AA%9E)のためのデザインパターン。CLI 開発をする時に参考にできそうだと思った。

## 関連するデザインパターンとの違い
### Composite pattern
NonterminalExpression に Composite pattern が使われる

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
