# Proxy pattern
## 概要
他のオブジェクトの代理、 代用を提供するデザインパターン。 プロキシーは、 元のオブジェクトへのアクセスを制御し、 元のオブジェクトへのリクエスト前後で処理を記述できる。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/proxy/diagram/abstract.svg" />

* Subject
  * Proxy と　RealSubject を同一視するためのクラス
* Proxy
  * Subject の実装
  * RealSubject の処理を肩代わりする
  * RealSubject へのアクセスを制御する
  * RealSubject 役が必要になってから、インスタンスを生成する
* RealSubject
  * Subject の実装

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ proxy pattern
```

#### ざっくり仕様
* web server
  * server へのアクセス制御
  * キャッシュによる処理の代理

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/proxy/diagram/impl.svg" />

型定義
```go
// ウェブサーバーのインターフェース
type WebServer interface {
	HandleRequest(request string) string
}

// 実際のウェブサーバーの実装
type RealWebServer struct{}

// プロキシウェブサーバーの実装
type ProxyWebServer struct {
	realWebServer  *RealWebServer
	cachedResponse string
}
```

ProxyWebServer と RealWebServer は共通のインタフェース WebServer を実装している。そのため、ProxyWebServer は RealWebServer に対しての透過的なアクセスを可能にしている。クライアントである main 関数は、実際のレスポンスが ProxyWebServer のキャッシュもしくは、RealWebServer から返されているのかを知る必要がない。

ProxyWebServer は、RealWebServer へのアクセス制御を行いつつ、キャッシュが存在すればレスポンスの作成を代理する。
```go
func (p *ProxyWebServer) HandleRequest(request string) string {
	// キャッシュが存在する場合はキャッシュを返す
	if p.cachedResponse != "" {
		fmt.Println("<-- chaced respose -->")
		return p.cachedResponse
	}

	// realWebServer の処理が必要になった場合のみ、呼び出す
	response := p.realWebServer.HandleRequest(request)
	p.cachedResponse = response

	return response
}

func NewWebSerber() WebServer {
	return &ProxyWebServer{
    // 初期が処理が重い場合は、必要になってからインスタンス化する
		realWebServer:  &RealWebServer{},
		cachedResponse: "",
	}
}

func main() {
	webServer := NewWebSerber()

	// リクエストを処理
	response := webServer.HandleRequest("GET /index.html")
	fmt.Println(response)
	// <html><body><h1>Hello, World!</h1></body></html>


	// 同じリクエストを処理（キャッシュからレスポンスが返される）
	response := webServer.HandleRequest("GET /index.html")
	fmt.Println(response)
	// <-- chaced respose -->
	// <html><body><h1>Hello, World!</h1></body></html>
}
```

## 使えそうなユースケース
① web サーバー（リバースプロキシ）
* アプリケーションの前後で処理を埋め込む
  * 静的ファイルをキャッシュして返す
    * アプリケーションサーバーの処理が必要になって、初めてリクエストを流す
  * ネットワークレベルの処理
    * ロードバランシング
    * アクセスログ

## 感想
コードレベルでも、システムレベルでも利用されるデザインパターンという印象。
コードレベルでは、初期化処理を遅らせたいオブジェクトや、アクセス頻度が高くキャッシュした値を返したい時などで使えそう。ただ、本当に proxy クラスが必要かは要検討。必要になってからリファクタリングでも良いと思う。
システムレベルでは、リバースプロキシなんかが当てはまると思う。

proxy オブジェクトは、
* 元のオブジェクトから責務を分離し、元のオブジェクトの前後で何かしら処理をするオブジェクト
* クライアントからみて、元のオブジェクトに対して透過的なアクセスを可能にする
* 元のオブジェクトへのアクセスへの制御・代理処理をする

という解釈。

## 関連するデザインパターンとの違い
### Decorator pattern
委譲を使って、合成したインスタンス操作の前後に処理を記述するという点では同じ。目的が異なる。
* Decorator pattern: 機能追加が目的
* Proxy pattern: アクセス制御、合成したインスタンスの代わりに処理を肩代わりすることが目的

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 
* https://refactoring.guru/ja/design-patterns/proxy
