package lambdas

import "context"

type PDFGeneratorPromt struct {
	HtmlTemplate string
	PageSize     string
	PageRanges   string
}

func Handler(ctx context.Context, prompt PDFGeneratorPromt) {
	panic("not implemented yet")
}
