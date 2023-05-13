# Facade pattern

## 概要
ライブラリー、 フレームワーク、 クラスなどの複雑な組み合わせた処理に対して、簡素化されたインターフェースを提供するデザインパターンです。
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/facade/diagram/abstract.svg" />

* Facade
  * その他大勢のシンプルな窓口
* その他大勢（classA, classB, classC）
  * Facade によってまとめ上げられるライブラリー、 フレームワーク、 クラスなど

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ facade pattern
```

#### ざっくり仕様
* 注文システム
* 注文が完了するまでにはいくつかの処理が必要で複雑
  * 注文
  * 支払い
  * 発送

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/facade/diagram/impl.svg" />

型定義。

OrderFacade(Facade)がその他大勢（OrderService, PaymentService, ShippingService）をまめあげている
```go
type OrderService struct{}

type PaymentService struct{}

type ShippingService struct{}

type OrderFacade struct {
	orderService    *OrderService
	paymentService  *PaymentService
	shippingService *ShippingService
}
```

注文・支払い・発送という処理をまとめたシンプルなインタフェースを OrderFacade の PlaceOrder メソッドが実装している。
```go
func (of *OrderFacade) PlaceOrder() {
	of.orderService.PlaceOrder()
	of.paymentService.ProcessPayment()
	of.shippingService.ShipProduct()
}

func main() {
	orderFacade := &OrderFacade{
		orderService:    &OrderService{},
		paymentService:  &PaymentService{},
		shippingService: &ShippingService{},
	}

	orderFacade.PlaceOrder()
	// <-- Output -->
	// 注文を受け付けました
	// 支払いを処理しました
	// 商品を発送しました
}

```

## 使えそうなユースケース
① 他のクラスからは隠蔽したい複数のクラスを使った複雑な処理

## 感想
Facade pattern はクライアントがシステムとの対話を単純化するために、システムの一連のインターフェースを提供している。「そのクラスを呼ぶためには、このクラスを呼んでこの処理が必要だよ」みたいなコミュニケーションが生まれたら、Facade pattern の出番という認識。

Facade となるクラスが、なんでも知っている神クラスにならないように注意が必要。

## 関連するデザインパターンとの違い
### Adapter pattern
Facade pattern も Adapter pattern も利用者からすれば、新しい振る舞いを定義してくれるデザインパターン。趣旨が違う。
* Adapter pattern
  * そのままでは使えないインタフェース(adaptee)を既存のインタフェースで使えるよにする
  * 基本的には 1対1 の関係
  * client -> adapter -> adaptee
* Facade pattern
  * 依存関係を追跡し、 正しい順序でメソッドを実行するなど、複雑な処理を隠蔽してシンプルなインタフェースを提供する
  * 基本的には 1対多 の関係
  * client -> facade -> classA, classB, classC

### [Simple factory pattern](https://techracho.bpsinc.jp/hachi8833/2020_12_03/46064#simple-factory)
Facade pattern は複雑な処理を隠蔽してシンプルなインタフェースを提供する。構造は似ているが、以下のようなオブジェクト生成のみを担当する factory class みたいなものは Facade pattern には当てはまらない認識。

```go
type MyInterface interface {
    SomeMethod() string
}

type StructA struct {}

func (a StructA) SomeMethod() string {
    return "StructAのメソッドが呼ばれました"
}

type StructB struct {}

func (b StructB) SomeMethod() string {
    return "StructBのメソッドが呼ばれました"
}

type Facade {
  structA StructA
  structA StructB
}

func (f *Facade) CreateStruct(condition string) MyInterface {
    switch condition {
    case "option1":
        return f.structA
    case "option2":
        return f.structB
    default:
        return nil
    }
}
```

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/facade
