# Prototype

## 概要
クラスに依存せず、インスタンスをクローンするデザインパターン。

インスタンスが自分自身を [deep copy](https://en.wikipedia.org/wiki/Object_copying) する方法を実装したデザインパターンといって差し支えないと思う。


<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/prototype/diagram/abstract.svg" />

* Client
  * clone して、コピーしたメソッドを使う
* Prototype
  * インスタンスをコピーするための振る舞いを定める
* ConcretePrototype
  * インスタンスを実際にコピーする
* 登場人物の紹介

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ prototype pattern
```

#### ざっくり仕様
* 簡易的なOSのファイルシステム
* Folder は Inode を実装した File か Folder を子要素に持つ
* deep copy を作成できる

#### 実装解説

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/prototype/diagram/impl.svg" />

Inode(Prototype) でクローンする振る舞いを定義する。 File と Folder がそれぞれ、実装する。
```go
type Inode interface {
	print(string)
	Clone() Inode
}

type File struct {
	name string
}

type Folder struct {
	children []Inode
	name     string
}
```

Folder の Clone メソッドに注目すると以下のことがわかる
* f.name で pkg 外で private なフィールドにアクセスできる
  * 必要に応じて pkg 外で private なメソッドにもアクセスできる
* 再帰的にクローンすることができる
  * 再帰構造に限らず、複雑なクローン処理をここに押し込める
```go
func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"} // private field にアクセス可能
	var tempChildren []Inode
  // 再帰的にクローン
	for _, i := range f.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy) // コピーする具体的な構造体を知らなくてもコピーできる
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
```

main 関数(Client) でクローン処理を実行
```go
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	cloneFolder := folder2.Clone()

	fmt.Println("\nPrinting hierarchy for orginal Folder")
	folder2.print("  ")

	fmt.Println("\nPrinting hierarchy for cloned Folder")
	cloneFolder.print("  ")
}

```

実行結果を見ると、全ての要素 の pointer address が、オリジナルとクローンで異なることがわかる。そのため、クローンした結果の値を書き換えてもオリジナルに影響を与えない。
```
Printing hierarchy for orginal Folder
  Folder2(pointer address: 0xc0000a41e8)
    Folder1(pointer address: 0xc0000a41f0)
        File1(pointer address: 0xc0000a41f8)
    File2(pointer address: 0xc0000a4200)
    File3(pointer address: 0xc0000a4208)

Printing hierarchy for cloned Folder
  Folder2_clone(pointer address: 0xc0000a4210)
    Folder1_clone(pointer address: 0xc0000a4218)
        File1_clone(pointer address: 0xc0000a4220)
    File2_clone(pointer address: 0xc0000a4228)
    File3_clone(pointer address: 0xc0000a4230)
```

## 使えそうなユースケース
① GUIベースのアプリケーション
* 例えば、ブラウザでノーコード簡単にWebページが作れるサービスがあったとする
* HTML要素で作られたコンポーネントを作ったり、コピーできたりする
* HTMLElement は振る舞いは定義されているが、中の詳細な階層構造までわからない
* 自分自身を作り出してくれる処理があったら便利（deep copy）

② private なフィールド・メソッドを持つインスタンス（構造体）を deep copy したい
* pkg 外から private なフィールド・メソッドを持つ構造体を deep copy したいとする
* pkg 外から private なフィールド・メソッドにアクセスできない
* 自分自身を deep copy してくれる prototype pattern が役にたつ
* 自分に生えてる clone メソッド内からのアクセスなので、もちろん private なフィールド・メソッドにアクセス可能
* 場合によっては複雑な clone 処理をカプセル化できる


## 感想
Prototype pattern はインスタンス（構造体）が自分自身を [deep copy](https://en.wikipedia.org/wiki/Object_copying) する方法を実装したデザインパターンだと思う。基本的には言語が持っている deep copy の実装を使った方がいいと思った。Prototype pattern はそのインスタンスに特化したクローン処理を実装するので。以下の時は有効だと思う。

* インスタンスを deep copy する方法がない
* private なフィールド・メソッドも deep copy したい

Porototype という抽象に依存させる（Dependency Inversion）ことによって、具体的な構造体（Class）を知らなくてもクローンできるのは便利だと思った。

書籍や参考サイトを見ていると、初期化方法のみが異なるサブクラスの数を減らせることが、 prototype pattern のメリットとして挙げられている。初期化方法のみが異なる場合、protype pattern に限らず以下の方法で回避できると思った。

① コンストラクタ内で条件分岐

② コンストラクタを静的メソッドにして複数持つ
```go
type Message struct {
  text string
  decorator string
}

func (m *Message) NewMessage(text) Message {
  return Message{text, ""}
}

func (m *Message) NewMessageStrong(text) Message {
  return Message{text, "*"}
}

func (m *Message) NewMessageWeak(text) Message {
  return Message{text, "~"}
}
```

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/prototype

## 関連するデザインパターンとの違い
