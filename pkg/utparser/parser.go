package utparser

import (
	"bytes"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/gaurishhs/utdocs/pkg/config"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	parser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type ParseData struct {
	Title        string
	Body         string
	PageTitle    string
	SiteLogo     string
	SiteName     string
	CodeTheme    string
	SidebarItems []config.SideBarItem
	Footer       string
}

var (
	DefaultTemplate = template.Must(template.New("template").Parse(Layout))
)

func NewParser() goldmark.Markdown {
	parser := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Linkify,
			meta.Meta,
			emoji.New(
				emoji.WithRenderingMethod(emoji.Twemoji),
			),
			extension.Strikethrough,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	return parser
}

func ParseFile(utparser goldmark.Markdown, file string, layout *template.Template, w io.Writer) {
	config := config.Configuration
	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := utparser.Convert(fileContent, &buf, parser.WithContext(context)); err != nil {
		log.Fatal("Error parsing file: ", err)
	}
	metaData := meta.Get(context)
	title := metaData["title"].(string)
	if title == "" {
		log.Fatal("Title not found in file: ", file)
	}
	data := ParseData{
		Title:        config.SiteName,
		Body:         buf.String(),
		PageTitle:    title,
		SiteName:     config.SiteName,
		CodeTheme:    config.CodeTheme,
		SidebarItems: config.SidebarItems,
		SiteLogo:     config.SiteLogo,
		Footer:       config.Footer,
	}

	if layout == nil {
		layout = DefaultTemplate
	}

	if err := layout.Execute(w, data); err != nil {
		log.Fatal("Error executing template: ", err)
	}
}
