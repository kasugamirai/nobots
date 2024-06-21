package mtt

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type plainTextRenderer struct{}

// RegisterFuncs implements renderer.NodeRenderer.
func (r *plainTextRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// doc
	reg.Register(ast.KindDocument, r.renderDocument)

	// blocks
	reg.Register(ast.KindHeading, r.renderNewLine)
	reg.Register(ast.KindBlockquote, r.renderNewLine)
	reg.Register(ast.KindCodeBlock, r.renderNewLine)
	reg.Register(ast.KindFencedCodeBlock, r.renderNewLine)
	reg.Register(ast.KindHTMLBlock, r.renderNewLine)
	reg.Register(ast.KindList, r.renderNewLine)
	reg.Register(ast.KindListItem, r.renderEmpty)
	reg.Register(ast.KindParagraph, r.renderNewLine)
	reg.Register(ast.KindTextBlock, r.renderNewLine)
	reg.Register(ast.KindThematicBreak, r.renderNewLine)

	// inlines
	reg.Register(ast.KindAutoLink, r.renderText)
	reg.Register(ast.KindCodeSpan, r.renderText)
	reg.Register(ast.KindEmphasis, r.renderEmpty)
	reg.Register(ast.KindImage, r.renderImage)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindRawHTML, r.renderText)
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderText)
}

const MdTest = `
# Heading 1

a simple block text

This is a *paragraph* of _text_.

- List item 1
- List item 2
- List item 3

> hello test
> world

[Link](https://www.example.com)

![image alt text](https://www.example.com/d.jpg)

- ![image alt *text* 2](https://www.example.com/d.jpg "here is title")

**Bold text**

*Italic text*
`

func (r *plainTextRenderer) renderDocument(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *plainTextRenderer) renderNewLine(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		writer.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *plainTextRenderer) renderEmpty(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *plainTextRenderer) renderText(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		writer.Write(n.Text(source))
	}
	return ast.WalkContinue, nil
}

func (r *plainTextRenderer) renderImage(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		image := n.(*ast.Image)
		writer.WriteString(" ")
		writer.Write(image.Destination)
		writer.WriteString(" ")
	}
	return ast.WalkContinue, nil
}

func (r *plainTextRenderer) renderLink(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		link := n.(*ast.Link)
		writer.WriteString(" ")
		writer.Write(link.Destination)
		writer.WriteString(" ")
	}
	return ast.WalkContinue, nil
}

func newPlainTextRenderer() *plainTextRenderer {
	return &plainTextRenderer{}
}

const isDebug = false

var md = goldmark.New(
	goldmark.WithRendererOptions(
		renderer.WithNodeRenderers(util.Prioritized(newPlainTextRenderer(), 0)),
	),
)

func MarkdownToPlainText(markdownText string) string {
	r, _ := MarkdownToPlainTextEx(markdownText)
	return r
}

func MarkdownToPlainTextEx(markdownText string) (string, error) {
	markdown := []byte(markdownText)

	if isDebug {
		md.Parser().Parse(text.NewReader(markdown)).Dump(markdown, 0)
	}
	var buf bytes.Buffer
	if err := md.Convert(markdown, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
