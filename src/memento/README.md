# Memento pattern

## 概要
オブジェクトの状態のスナップショットを撮り、以前の状態に戻せるようにしたデザインパターン。

private フィールドを public にせずにカプセル化を実現したまま、スナップショットを作成し、オブジェクトをある時点の状態に復元できる。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/memento/diagram/abstract.svg" />

* Originator
  * 状態を持つオブジェクト
  * 自身の状態を保存した Memento オブジェクトを生成する
  * Memento オブジェクトから状態を復元することが可能
* Memento
  * Originator のある時点での状態を保存したオブジェクト
* Caretaker
  * Memento オブジェクトを保持し、必要な場合に Originator に渡す
  * Originator の状態の履歴を管理する

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ memento pattern
```

#### ざっくり仕様
* エディター機能
* 任意の点でスナップショットを撮れる
* スナップショットの状態に復元できる

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/memento/diagram/impl.svg" />

型定義

```go
// Originator（作成者）の構造体
type Editor struct {
	text string
}

// Memento（記念品）の構造体
type EditorMemento struct {
	text string
}

// Caretaker（世話役）の構造体
type History struct {
	mementos []*EditorMemento
}
```

CreateMemento によって作成された EditorMemento(Memento) が Editor(Originator) の private な text の状態を保持している。そのため、 text フィールドを public にしてカプセル化の破壊をせずとも、ある時点での Editor の復元が可能。
```go
// Mementoを生成するメソッド
func (e *Editor) CreateMemento() *EditorMemento {
	return &EditorMemento{text: e.text}
}

// Mementoから状態を復元するメソッド
func (e *Editor) RestoreMemento(m *EditorMemento) {
	e.text = m.text
}
```

History は EditorMemento を保持し、必要な時に Editor に渡す役割を果たす。

```go
func main() {
	editor := &Editor{}
	history := &History{}

	// テキストを設定し表示
	editor.SetText("Hello, World!")
	editor.ShowText()
	// <-- Output -->
	// Current Text: Hello, World!

	// Mementoを作成しHistoryに保存
	history.AddMemento(editor.CreateMemento())

	// テキストを変更し表示
	editor.SetText("Goodbye, World!")
	editor.ShowText()
	// <-- Output -->
	// Current Text: Goodbye, World!

	// Mementoから状態を復元し表示
	memento := history.GetLastMemento()
	editor.RestoreMemento(memento)
	editor.ShowText()
	// <-- Output -->
	// Current Text: Hello, World!
}
```

## 使えそうなユースケース
① エディターとか

## 感想
スナップショットを作成し、以前の状態に戻れるようにしている。 private フィールドも再現できるところが良いところ。無駄に public にしなくて良いのでカプセル化を破壊しない。 Prototype pattern を使って、 undo を実装するデザインパターンという印象。

Originator を復元するために Memento が Orginator のフィールドの値を知っている必要がある。大抵の場合、Originator と Memento が同じプロパティーを保持するという構造になると考えられる。その結果、 Originator にフィールドを追加したら、 Memento にもフィールドの追加が必要で、ちょっと面倒だなと思う。

## 関連するデザインパターンとの違い
### Protptype pattern
スナップショットを作るという点ではほぼ一緒。 prototype pattern は自身をクローンするが、 memento pattern は memento （状態を表すオブジェクト）を作る。memenot pattern ではさらに caretaker と呼ばれる存在が memento 群を管理し、 任意の memento の状態に戻れるようにしている。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/memento
