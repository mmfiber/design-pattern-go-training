# Command pattern

## 概要
命令をオブジェクトとして表すデザインパターン。

オブジェクト化することで、リクエストをメソッドの引数として渡したり、リクエストの実行を遅らせたり、キューイングしたり、スタックを使って履歴管理を行なうことが可能。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/command/diagram/abstract.svg" />

* 抽象的な説明
* 登場人物の紹介

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ command pattern
```

#### ざっくり仕様
* エディター
  * copy
  * cut
  * paste
  * undo

#### 実装解説
Application が Invoker と Reciever の両方を担っている
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/command/diagram/impl.svg" />

型定義
```go
type Editor struct {
	Text string
}

type Command interface {
	Execute() string
}

type CopyCommand struct {
	App *Application
}

type CutCommand struct {
	App *Application
}

type PasteCommand struct {
	App *Application
}

type UndoCommand struct {
	App *Application
}
```

Command を実装した ConcreteCommand(CopyCommand, CutCommand, PasteCommand, UndoCommand) が app.ExecuteCommand によって実行される。それぞれのオブジェクトによって処理がカプセル化されている。そのため、同じ実行方法でも出力結果が異なる。


```go
func main() {
	editor := &Editor{Text: "Hello, World!"}
	app := &Application{Editor: editor}

	copyCmd := &CopyCommand{App: app}
	app.ExecuteCommand(copyCmd)
	// Copying text: Hello, World!

	cutCmd := &CutCommand{App: app}
	app.ExecuteCommand(cutCmd)
	// Cutting text: Hello, World!

	pasteCmd := &PasteCommand{App: app}
	app.ExecuteCommand(pasteCmd)
	// Pasted text

	undoCmd := &UndoCommand{App: app}
	app.ExecuteCommand(undoCmd)
	// Undo last command
}
```

実行時に Command をスタックしているので、undo処理が可能な状態。
```go
type UndoCommand struct {
	App *Application
}

func (u *UndoCommand) Execute() string {
	// Undo command logic...
	if len(u.App.CommandHistory) > 0 {
		u.App.CommandHistory = u.App.CommandHistory[:len(u.App.CommandHistory)-1]
		return "Undo last command"
	}
	return "Nothing to undo"
}

func (a *Application) ExecuteCommand(command Command) {
	a.CommandHistory = append(a.CommandHistory, command)
	fmt.Println(command.Execute())
}
```

## 使えそうなユースケース
① mdn のイベントリスナーのコールバック関数
- Event というコマンドがパラメータとしてわたってくる

## 感想
パラメーターをオブジェクト化できる。オブジェクトにする際に Command インタフェースを切り、抽象に依存させるため具体的な処理は ConcreteCommand にカプセル化する。Command インタフェースを実装したオブジェクトはキューイングして扱いやすい。また、Command の履歴をスタックしておけば、undo 処理とかもできる。
Command が複数あったり、Command に対してキューやスタックを使いたいという要望があって初めて実装すべき。複雑度が増すから。

## 関連するデザインパターンとの違い

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社
* https://refactoring.guru/ja/design-patterns/command 

