# Adapter pattern

## 概要
Adapter pattern は二つの異なるインタフェースのずれを吸収するデザインパターン。

<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/adapter/diagram/abstract.svg" />

* Target
  * クライアントが求める振る舞いを定義したインタフェース
* Client
  * Target で定められた振る舞いを使う人
* Adaptee
  * Target とは異なる振る舞いをする何か
* Adapter
  * Target と Adaptee を繋ぐもの

書籍では、継承を使ったパターンと委譲を使ったパターンが紹介されていた。

それぞれの使い分けは、
* Adaptee と Adapter が親子関係である -> 継承
* Adapter と Adaptee は使役の関係である -> 委譲

で、良いと思う。

## サンプル実装
#### 実行方法
```sh
❯ go run main.go  
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select: 
  ❯ adapter pattern
```

#### ざっくり仕様
* Banner は文字列を装飾する機能を持つ
* Client 文字列の表示に強弱をつけて表示したい
  * それを下記メソッドで表現したい 
    * PrintWeak
    * PrintStrong

#### 実装解説
<img src="https://github.com/mmfiber/design-pattern-go-training/blob/main/src/adapter/diagram/impl.svg" />

PrintBanner は Print インタフェースの要求を Banner を使って満たす実装。
PrintBanner と Banner は使役の関係なので、今回は委譲を使った実装をした。

Print(Target) の要求するインタフェースと Banner(Adaptee) の振る舞いの違いを、PrintBanner(Adapter) が吸収している。
```go
type Printer interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	text string
}

type PrintBanner struct {
	banner *Banner
}

func (pb PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
```

## 使えそうなユースケース
① 外部APIへのリクエスト
* 例えば、
  * ユーザーの行動データを取得する外部サービスを導入したいとする
  * 候補がAとBがあるので、どちらがいいか試したい
  * AとBに期待する機能は同じだが、それぞれのAPIへ渡す値のデータ構造が異なる
* みたいな時に、AとBのような外部サービスと依存性を切り離し、ビジネスロジックとAPIへのリクエストの間を Adapter pattern で埋めてやるとよさそう
* AとBのサービス切り替えてもビジネスロジックを変更する必要がない
* 外部サービスに依存しないということは、使う外部サービスの決定を遅らせることができる
  * 企画目線で言うと、あらゆる外部サービスを試しやすいプロダクトになっている
  * エンジニア目線で言うと、詳細な決定を遅らせられるので、他の開発を進めることができる

② ORM
* ORM はまさに Adapter pattern を実装した機能に思える
* アプリケーション側は dbClient のような ORM を用いてデータを取得する
  ```go
  dbClient.getUser(userid)
  ```
* 実際にDBから値を取得するには SQL を描かないといけない
  ```sql
  SELECT * FROM users AS u WHERE u.id = userid;
  ```
* ORMがアプリケーションからの入力をSQLに変えDBに渡し、DBからの出力結果をアプリケーションで使いやすい形に変換している

## 感想
ビジネスロジックなど不変な箇所と、外部サービスなど変化する箇所に Adapter パターンを適用させるとよさそう。こちらで制御できない外部サービスの変化などに対しても強くなる。
出来上がっているもの同士をつなぎ合わせるみたいな観点でも機能しそう。アーキテクチャレベルの話になるが、BFFとかまさにそれ。

## 参考文献
* 結城浩(2020). 増補改訂版 Java言語で学ぶデザインパターン入門. 東京: SBクリエイディブ株式会社 

## 関連するデザインパターンとの違い
