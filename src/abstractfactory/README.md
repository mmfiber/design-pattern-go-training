# Abstract factory pattern

## 概要
Factory method pattern をさらに抽象化したもの。Factory method pattern の集合と考えて良い。違いは、Factory から作られるインスタンスの詳細な型情報を知っているかどうか。
* Factory method pattern   : 知っている
* Abstract factory pattern : 知らない

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/abstractfactory/diagram/abstract.svg" />

* Product
  * Factory によって作り出されるインスタンスの抽象クラス
* Factory
  * Product を生み出すための抽象クラス
* ConcreteProduct
  * Product を継承したもの
* ConcreteFactory
  * Factory を継承したもの

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ abstract factory patttern
```

#### ざっくり仕様
* UI フレームワーク
* OS ごとに異なるボタンを作る

#### 実装解説
<div style="display: flex">
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/abstractfactory/diagram/impl.windows.svg" />
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/abstractfactory/diagram/impl.mac.svg" />
</div>


Btn(Product) と UIFactory(Factory) で振る舞いを表現する
```go
// go にはオブジェクト思考の概念がないため、
// 振る舞い部分のみインタフェースを代用する

type Btn interface {
	Click()
}

type UIFactory interface {
	CreateBtn() Btn
}
```

Windows 用の実装（ConcreteProduct, ConcreteFactory）

```go
type WindowsBtn struct{}

func (b *WindowsBtn) Click() {
	fmt.Println("windows btn")
}

type WindowsUIFactory struct{}

func NewWindowsUIFactory() UIFactory {
	return &WindowsUIFactory{}
}

func (f *WindowsUIFactory) CreateBtn() Btn {
	return &WindowsBtn{}
}
```

Mac 用の実装（ConcreteProduct, ConcreteFactory）

```go
type MacBtn struct{}

func (b *MacBtn) Click() {
	fmt.Println("Mac btn")
}

type MacUIFactory struct{}

func NewMacUIFactory() UIFactory {
	return &MacUIFactory{}
}

func (f *MacUIFactory) CreateBtn() Btn {
	return &MacBtn{}
}
```

main 関数で注目したいのは作られる 詳細な Btn(Product) の型を知らないこと。

```go
func GetUIFactory(os OS) UIFactory {
	switch os {
	case Windwos:
		return NewWindowsUIFactory()
	case Mac:
		return NewMacUIFactory()
	default:
		panic("invalid os")
	}
}

func main() {
	var factory UIFactory

	factory = GetUIFactory(Windwos)
  // 作られる Product の詳細知らない
  // 実際には WindowsBtn.CLick() を実行しているが、型的には Btn.CLick() を実行
	factory.CreateBtn().Click() 

	factory = GetUIFactory(Mac)
  // 作られる Product の詳細知らない
  // 実際には MacBtn.CLick() を実行しているが、型的には Btn.CLick() を実行
	factory.CreateBtn().Click()
}
```
## 使えそうなユースケース
①[Factory method pattern](https://github.com/mmfiber/design-pattern-go-training/tree/main/src/factorymethod) でどの factory を使っているかも抽象化したい時
* 具体的なユースケースが思いつかない

## 感想
正直、Factory method pattern との違いがあまりないと思った。Factory method pattern は使っている詳細な factory（ConcreteFactory） を知っているが、Abstract factory pattern は抽象的な factory（Abstract Factory）しか知らない。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社
* https://refactoring.guru/ja/design-patterns/abstract-factory

## 関連するデザインパターンとの違い
