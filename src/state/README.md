# State pattern

## 概要
状態ごとにオブジェクトを作り、状態によって振る舞いを変化させるデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/state/diagram/abstract.svg" />

* Context
  * ConcreteState のいずれか一つへ参照を持つ
  * 状態に固有の作業は、 すべて ConcreteState に委譲する
  * 新しい State を渡すための setter を持っている
* State
  * 状態ごとの共通の振る舞いを定義したインタフェース
* ConcreteState
  * State の実装

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ state pattern
```

#### ざっくり仕様
* 自動販売機

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/state/diagram/impl.svg" />

型定義

```go
// State
type State interface {
	InsertCoin()
	SelectProduct()
	Dispense()
}

// ConcreteState
type NoCoinState struct{}

// ConcreteState
type HasCoinState struct{}

// Context
type VendingMachine struct {
	state State
}
```

InsertCoin, Dispense メソッドで state の状態を変化させ、状態に固有の作業は NoCoinState, HasCoinState といった ConcreteState に委譲している。ここで、 state の変更は、 ConcreteState が行ってもよい。 ConcreteState が Context を保持し、 Context のセッターを呼ぶことで実装可能。

```go
func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin()
	vm.state = &HasCoinState{}
}

func (vm *VendingMachine) SelectProduct() {
	vm.state.SelectProduct()
}

func (vm *VendingMachine) Dispense() {
	vm.state.Dispense()
	vm.state = &NoCoinState{}
}

func main() {
	vendingMachine := &VendingMachine{state: &NoCoinState{}}

	vendingMachine.InsertCoin()    // コインが挿入されました。
	vendingMachine.SelectProduct() // 商品が選択されました。
	vendingMachine.Dispense()      // 商品が提供されました。
}
```

## 使えそうなユースケース
① 状態ごとに異なる振る舞いがあり、状態の数が多い時

## 感想
状態数が多く、状態ごとに振る舞いを条件分岐で変えているような時に役立つ。特に複数のメソッドで同様の条件分岐をしている際は。

```go
type Document struct {
  state string
}

func (d *Document) Publish() {
  // 状態によって異なる振る舞いの数が多い
  switch d.state {
    case "draft":
      d.state = "moderation"
        break
    case "moderation":
      if (currentUser.role == "admin")
          d.state = "published"
      break
    case "published":
      break
    default:
      // error
  }
}
```

逆に状態数が2個など、少ない時にこのデザインパターンは過剰。イタズラに複雑さを上げるだけ。現在の状態数が多い、もしくは将来的に多くなることが予想される場合のみ、このデザインパターンの導入を検討するとよさそう。

## 関連するデザインパターンとの違い
### Strategy pattern
インタフェースに沿って、実装したオブジェクトを実行時に切り替えて使うという点では strategy pattern と state pattern は非常に似ている。この二つのデザインパターンの違いは、インタフェースを実装したオブジェクト同士が独立しているか（お互いを知っているか）という点にある。

* strategy pattern: お互いを知らない。独立している。
* state pattern: お互いを知っている。ある状態からある状態への遷移が可能。


## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/state
