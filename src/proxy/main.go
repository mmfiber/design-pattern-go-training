package proxy

import (
	"fmt"
	"strings"
)

// ウェブサーバーのインターフェース
type WebServer interface {
	HandleRequest(request string) string
}

// 実際のウェブサーバーの実装
type RealWebServer struct{}

func (r *RealWebServer) HandleRequest(request string) string {
	return "<html><body><h1>Hello, World!</h1></body></html>"
}

// プロキシウェブサーバーの実装
type ProxyWebServer struct {
	realWebServer  *RealWebServer
	cachedResponse string
}

func (p *ProxyWebServer) HandleRequest(request string) string {
	// キャッシュが存在する場合はキャッシュを返す
	if p.cachedResponse != "" {
		fmt.Println("<-- chaced respose -->")
		return p.cachedResponse
	}

	response := p.realWebServer.HandleRequest(request)
	p.cachedResponse = response

	return response
}

func NewWebSerber() WebServer {
	return &ProxyWebServer{
		realWebServer:  &RealWebServer{},
		cachedResponse: "",
	}
}

func main() {
	// プロキシウェブサーバーを作成
	webServer := NewWebSerber()

	// リクエストを処理
	response := webServer.HandleRequest("GET /index.html")
	fmt.Println(response) // 実際のウェブサーバーへのリクエストが発生し、レスポンスが表示される

	fmt.Println(strings.Repeat("-", 20))

	// 同じリクエストを処理（キャッシュからレスポンスが返される）
	response = webServer.HandleRequest("GET /index.html")
	fmt.Println(response) // キャッシュからレスポンスが表示される
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "proxy pattern"
}

func (e Executer) Do() {
	main()
}
