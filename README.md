# design-pattern-go-training

go 初学者が go とデザインパターンを勉強するために作ったレポジトリ。
`src` 以下のそれぞれのディレクトリに `README.md` と `main.go` が実装されている。

- `README.md`: 学習内容のアウトプット
  - 目的
  - 抽象的な説明
  - 具体的な説明（main.go の解説）
  - 使い所
  - 感想（所感）
  - 他のデザインパターンとの違い
- `main.go`: 各デザインパターンを使った具体的な実装

### 学んだ結果、、、

(以下、個人的な見解)

デザインパターンはそれぞれの目的と使い所だけ理解していれば良いと感じた。


各デザインパターンの詳細（登場人物など）を全て記憶しておく必要はない。詳細な実装方法は、実装時に調べればよい。

大切なのは、
- どのデザインパターンを使用するべきかを判断できること
- デザインパターンの概念を一度理解しておき、状況に応じてな必要なエッセンスを抽出できること


デザインパターンは、あるユースケースに対する一つの模範的な回答と捉えると腑に落ちた。

例えば、Iterator pattern が効果的なユースケースは[こちら](https://refactoring.guru/ja/design-patterns/iterator#problem)に例が記載されている。[Refactoring.Guru](https://refactoring.guru/ja/design-patterns) の各デザインパターンの問題を参照してみるとよい。デザインパターンが有効となるユースケースが紹介されている。


デザインパターンごとの違いを理解しておくことは、各デザインパターンの目的と使い所を理解するのに役立つ。
- 名前が違うだけの、ほとんど同じ構成のデザインパターンが存在するため。
- デザインパターンの中で、別のデザインパターンが使われるため。

## go install
[goenv を使ったインストール方法](https://github.com/syndbg/goenv/blob/master/INSTALL.md)

## plantUml
[plantUml](https://plantuml.com/)
[extension](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml)

plantUml server 起動
```sh
docker run -d -p 9000:8080 plantuml/plantuml-server:jetty
```

extension の設定で、
```json
{
  "plantuml.render": "PlantUMLServer",
  "plantuml.server": "http://localhost:9000",
}
```
