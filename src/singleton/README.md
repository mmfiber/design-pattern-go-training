# Singleton pattern

## 概要
Singleton pattern はインスタンスが一つしかないことを保証するデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/singleton/diagram/abstract.svg" />

* Singleton
  * インスタンスを生成
  * 唯一のインスタンスを生成するための static メソッドを持っている

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ singleton pattern
```

#### ざっくり仕様
* DBClient インスタンスを作る
* DBClient インスタンスは一つのみ存在（pointer address が一致するはず）

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/singleton/diagram/impl.svg" />

dbClient が nil なら新たに Client 生成。

```go
var (
	dbClient DBClient
	lock     = &sync.Mutex{}
)

func getDBClient(ch chan DBClient) {
	if dbClient == nil {
		lock.Lock()
		defer lock.Unlock()

		fmt.Println("Cleate new instance")
		dbClient = &MySqlDBClient{} // construct
	} else {
		fmt.Println("Instance is already decleared")
	}

	ch <- dbClient
 
}
```

go のようなマルチスレッドの言語では、スレッドセーフな実装になっているか気にする必要がある。どのスレッドから呼ばれても一つしかインスタンスを返さないことを保証するために、[mutex](https://pkg.go.dev/sync#Mutex) でロックしている。

マルチスレッドでも動作するか検証するために、 goroutine を使っている。
```go
func main() {
	ch := make(chan DBClient, 10)
	for i := 0; i < 30; i++ {
		go getDBClient(ch)
	}

	for cilent := range ch {
		fmt.Printf("client address: %p\n", &cilent)
	}
}
```

実行結果を見ると以下のことがわかる
* 一度だけコンストラクタが呼ばれいている
* 返されるインスタンスのアドレスが一致しているので、インスタンスが一つしか作られていない
```sh
Cleate new instance
Instance is already decleared
 ...
Instance is already decleared
client address: 0xc0000a89a0
 ...
client address: 0xc0000a89a0
```

## 使えそうなユースケース
① 上述のような DBClient とか

## 感想
Singleton pattern は必ずインスタンスは一つということを保証してくれる代わりに、以下のことを考慮しないといけない。
* マルチスレッドセーフになっているか
* テスト時の mock class の作成方法
  * テストフレームワークなんかを使ってると、mock class をフレームワークが作ってくれる
  * その時に、継承が用いられることがしばしば
  * コンスラクタが private な Singleton pattern だと mock class を作れなかったりする


なので、インスタンスが必ず一つであって欲しい状況を考えるのが大切だと思った。Singleton pattern は以下の状況の時に使うと良いんじゃないかなと思っている。
* 無駄にメモリを使いたくない
* 複数インスタンスがあると不都合が起きる（この状況になる例が思いつかない、、、）
* 状態を保持したい
* 状態を使った振る舞いを定義したい

状態を持たないのであれば、static method を持ったクラスを作る or 単純に関数群を作れば良いと思う。そちらの方がテストもしやすい。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/singleton

## 関連するデザインパターンとの違い
