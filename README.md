# design-pattern-go-training

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
