# Flyweight pattern

## 概要

複数のオブジェクトで共通する部分をオブジェクト化し、再利用するデザインパターン。メモリの消費を少なくすることを目的としている。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/flyweight/diagram/abstract.svg" />

* Flyweight
  * 共有する値(Intrinsic State)を持っているオブジェクト
  * 共有しない値(Extrinsic State)は、メソッド実行時に外から与えられるか、Flyweight オブジェクトの利用者が保持する
* FlyweightFactory
  * Flyweight の作成を担当
  * 一度作った Flyweight は保持しておく
  * 再度同じ Flyweight を要求されたら、保持している Flyweight を返す

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ flyweight pattern
```

#### ざっくり仕様
* 共有させる情報を intrinsicState として ConcreteFlyweight が保持
* 共有させない情報は extrinscState として Operation 実行時に渡す

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/flyweight/diagram/impl.svg" />

型定義。

Flyweight を複数定義することも想定される。そのため、Flyweight インタフェースで振る舞いを定義し、ConcreteFlyweight で実装する。

```go
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

type Flyweight interface {
	Operation(extrinsicState string)
}

type ConcreteFlyweight struct {
	intrinsicState string
}
```

GetFlyweight メソッドで、 ConcreteFlyweight の生成を行っている。一度作成した ConcreteFlyweight は map にオブジェクトを保存しておく。再度同じ key(Intrisic State) が要求された際は、保存してあるオブジェクトを返す。

```go
func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
	// すでに存在するFlyweightオブジェクトを再利用
	if flyweight, ok := ff.flyweights[key]; ok {
		return flyweight
	}

	flyweight := &ConcreteFlyweight{intrinsicState: key}
	ff.flyweights[key] = flyweight
	return flyweight
}

func main() {
	factory := FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}

	// Flyweightオブジェクトの取得と操作
	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("state1")

	flyweight2 := factory.GetFlyweight("B")
	flyweight2.Operation("state2")

	flyweight3 := factory.GetFlyweight("A")
	flyweight3.Operation("state3")
}
```

実行結果。

flyweight1 と flyweight3 のアドレスが一致していることがわかる。つまり、Intrinsc State が共通している場合、メモリを共有し、再利用ができている。

```sh
// flyweight1
Flyweight Address: 0xc00014b080, Intrinsic State: A, Extrinsic State: state1

// flyweight2
Flyweight Address: 0xc00014b0b0, Intrinsic State: B, Extrinsic State: state2

// flyweight3
Flyweight Address: 0xc00014b080, Intrinsic State: A, Extrinsic State: state3
```

## 使えそうなユースケース
① ゲーム作成 ② 参照渡し

## 感想
以下の限られた条件下で利用を検討するデザインパターンという認識。
* メモリに収まりきらない膨大な数のオブジェクトが存在する
* メモリの容量を増やすことができず、アプリケーション側が対応する必要がある

サーバサイドなど、メモリの容量を変えることができる場合はそちらを検討した方が良さそう。Intrisic State として共通の不変なオブジェクトを作成すると、予想できない将来的な変更に対応が大変そう。

ブラウザで実行されるアプリケーションだと、アプリケーション側がメモリの容量を勝手に増やすことができない。そのため、メモリを節約する方法を考える必要があり、このデザインパターンが役に立つかも。経験上、Vue などのフレームワークを使用した、リッチな web アプリケーションを作ってもメモリの問題に直面したことはない。

## 関連するデザインパターンとの違い
### Singleton pattern
オブジェクトの持つ値の変更可否
* Singleton: 変更可能
* Flyweight: 変更不可

Extrinsic な情報を持つか
* Singleton: 持たない。共通（Intrinsic）の値のみ
* Flyweight: 持つ。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/flyweight
