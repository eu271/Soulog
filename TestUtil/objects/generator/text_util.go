package generator

const (
	BOLD = iota
	ITALIC
	CURIVE
	TRASH
)

type Text interface {
	Word(lenght int) string
	Sentence(lenghtInWords int) string
	Paragraph(lenghtInSentences int) string
	Text(lenghtInParagraphs int) string
}

type Url interface {
	Domain() string
	Email() string
	Slug(sentece string) string
}

type Markdown interface {
	word(lenght int, mode int) string
	Blockquote(text string) string
	List(elements int) string
	Title(level int) string
	Paragraph(lenghtInSentences int, addModesToWords bool) string
	Text(lenghtInParagraph int, addModesToWords bool) string
	Link(url string) string
	Image(url string) string
}

type Numbers interface {
	Id(lenght int) string
	Int(min, max int) int
}

var wordList = []string{
	"asdasd",
	"adasdasd",
}
