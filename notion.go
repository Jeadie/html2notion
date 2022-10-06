package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/jomei/notionapi"
	"golang.org/x/net/html"
)

// ConvertNode from a html node into a Notion Block representation. Parsers children nodes of the HTML too.
func ConvertDocument(d *goquery.Document) []notionapi.Block {
	nodes := d.Children().Nodes()
	blocks := make([]notionapi.Block, len(n))

	for i, n := range nodes {
		blocks[i] = ConvertNode(n, d)
	}
	return blocks
}

func ConvertNode(n *html.Node, d *goquery.Document) notionapi.Block {
	
}

func Get(type ) {
	switch BlockType(raw["type"].(string)) {
	case BlockTypeParagraph:
		b = &ParagraphBlock{}
	case BlockTypeHeading1:
		b = &Heading1Block{}
	case BlockTypeHeading2:
		b = &Heading2Block{}
	case BlockTypeHeading3:
		b = &Heading3Block{}
	case BlockCallout:
		b = &CalloutBlock{}
	case BlockQuote:
		b = &QuoteBlock{}
	case BlockTypeBulletedListItem:
		b = &BulletedListItemBlock{}
	case BlockTypeNumberedListItem:
		b = &NumberedListItemBlock{}
	case BlockTypeToDo:
		b = &ToDoBlock{}
	case BlockTypeCode:
		b = &CodeBlock{}
	case BlockTypeToggle:
		b = &ToggleBlock{}
	case BlockTypeChildPage:
		b = &ChildPageBlock{}
	case BlockTypeEmbed:
		b = &EmbedBlock{}
	case BlockTypeImage:
		b = &ImageBlock{}
	case BlockTypeVideo:
		b = &VideoBlock{}
	case BlockTypeFile:
		b = &FileBlock{}
	case BlockTypePdf:
		b = &PdfBlock{}
	case BlockTypeBookmark:
		b = &BookmarkBlock{}
	case BlockTypeChildDatabase:
		b = &ChildDatabaseBlock{}
	case BlockTypeTableOfContents:
		b = &TableOfContentsBlock{}
	case BlockTypeDivider:
		b = &DividerBlock{}
	case BlockTypeEquation:
		b = &EquationBlock{}
	case BlockTypeBreadcrumb:
		b = &BreadcrumbBlock{}
	case BlockTypeColumn:
		b = &ColumnBlock{}
	case BlockTypeColumnList:
		b = &ColumnListBlock{}
	case BlockTypeLinkPreview:
		b = &LinkPreviewBlock{}
	case BlockTypeLinkToPage:
		b = &LinkToPageBlock{}
	case BlockTypeTemplate:
		b = &TemplateBlock{}
	case BlockTypeSyncedBlock:
		b = &SyncedBlock{}
	case BlockTypeTableBlock:
		b = &TableBlock{}
	case BlockTypeTableRowBlock:
		b = &TableRowBlock{}

	case BlockTypeUnsupported:
		b = &UnsupportedBlock{}
	default:
		return &UnsupportedBlock{}, nil
	}
}