package builder

import (
	"fmt"
	"os"
	"strings"
)

type Builder interface {
	makeTitle(string)
	makeString(string)
	makeItems([]string)
	build()
}

type TextBuilder struct {
	sb *strings.Builder
}

func NewTextBuilder() TextBuilder {
	return TextBuilder{&strings.Builder{}}
}

func (b *TextBuilder) makeTitle(title string) {
	fmt.Fprintf(b.sb, "『 %s 』\n\n", title)
}

func (b *TextBuilder) makeString(str string) {
	fmt.Fprintf(b.sb, "■ %s\n\n", str)
}

func (b *TextBuilder) makeItems(items []string) {
	for _, item := range items {
		fmt.Fprintf(b.sb, "\t・ %s\n", item)
	}
}

func (b *TextBuilder) build() {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s\n", strings.Repeat("=", 20))
	fmt.Fprintf(&sb, "%s\n", b.sb.String())
	fmt.Fprintf(&sb, "%s\n", strings.Repeat("=", 20))
	b.sb = &sb
}

func (b *TextBuilder) GetResult() string {
	return b.sb.String()
}

type HTMLBuilder struct {
	sb *strings.Builder
}

func NewHTMLBuilder() HTMLBuilder {
	return HTMLBuilder{&strings.Builder{}}
}

func (b *HTMLBuilder) makeTitle(title string) {
	fmt.Fprintf(b.sb, "<h1>%s</h1>\n", title)
}

func (b *HTMLBuilder) makeString(str string) {
	fmt.Fprintf(b.sb, "<p>%s</p>\n", str)
}

func (b *HTMLBuilder) makeItems(items []string) {
	fmt.Fprintln(b.sb, "<ul>")
	defer fmt.Fprintln(b.sb, "</ul>")

	for _, item := range items {
		fmt.Fprintf(b.sb, "<li>%s</li>\n", item)
	}
}

func (b *HTMLBuilder) build() {
	var sb strings.Builder
	fmt.Fprintf(&sb, "<html>\n<body>\n")
	fmt.Fprintf(&sb, "%s", b.sb.String())
	fmt.Fprintf(&sb, "</html>\n</body>\n")
	b.sb = &sb
}

func (b *HTMLBuilder) GetResult() string {
	dir := "src/builder/dist"
	os.Mkdir(dir, os.ModePerm)

	path := fmt.Sprintf("%s/index.html", dir)
	f, _ := os.Create(path)
	defer f.Close()

	str := b.sb.String()
	f.WriteString(str)

	return str
}

type Director struct {
	builder Builder
}

func NewDirector() Director {
	return Director{}
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.makeTitle("Greeting")
	d.builder.makeString("朝から昼にかけて")
	d.builder.makeItems([]string{
		"おはようございます。",
		"こんにちは。",
	})
	d.builder.makeString("夜に")
	d.builder.makeItems([]string{
		"こんばんは。",
		"おやすみなさい。",
		"さようなら。",
	})
	d.builder.build()
}

func main() {
	director := NewDirector()

	textBuilder := NewTextBuilder()
	director.SetBuilder(&textBuilder)
	director.Construct()
	fmt.Println(textBuilder.GetResult())

	htmlBuilder := NewHTMLBuilder()
	director.SetBuilder(&htmlBuilder)
	director.Construct()
	fmt.Println(htmlBuilder.GetResult())
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "builder pattern"
}

func (e Executer) Do() {
	main()
}
