package memento

import "fmt"

// Originator（作成者）の構造体
type Editor struct {
	text string
}

// Memento（記念品）の構造体
type EditorMemento struct {
	text string
}

// Mementoを生成するメソッド
func (e *Editor) CreateMemento() *EditorMemento {
	return &EditorMemento{text: e.text}
}

// Mementoから状態を復元するメソッド
func (e *Editor) RestoreMemento(m *EditorMemento) {
	e.text = m.text
}

// テキストを設定するメソッド
func (e *Editor) SetText(text string) {
	e.text = text
}

// テキストを表示するメソッド
func (e *Editor) ShowText() {
	fmt.Println("Current Text:", e.text)
}

// Caretaker（世話役）の構造体
type History struct {
	mementos []*EditorMemento
}

// Mementoを追加するメソッド
func (h *History) AddMemento(m *EditorMemento) {
	h.mementos = append(h.mementos, m)
}

// Mementoを取得するメソッド
func (h *History) GetMemento(index int) *EditorMemento {
	return h.mementos[index]
}

// 最後のMementoを取得するメソッド
func (h *History) GetLastMemento() *EditorMemento {
	if len(h.mementos) == 0 {
		return nil
	}
	return h.mementos[len(h.mementos)-1]
}

func main() {
	editor := &Editor{}
	history := &History{}

	// テキストを設定し表示
	editor.SetText("Hello, World!")
	editor.ShowText()

	// Mementoを作成しHistoryに保存
	history.AddMemento(editor.CreateMemento())

	// テキストを変更し表示
	editor.SetText("Goodbye, World!")
	editor.ShowText()

	// Mementoを作成しHistoryに保存
	history.AddMemento(editor.CreateMemento())

	// テキストを変更し表示
	editor.SetText("Gopher")
	editor.ShowText()

	// Mementoから状態を復元し表示
	memento := history.GetLastMemento()
	editor.RestoreMemento(memento)
	editor.ShowText()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "memento pattern"
}

func (e Executer) Do() {
	main()
}
