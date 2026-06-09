package validator

import (
	"errors"
	"flag"
)

type ReplacerParams struct {
	SourceFilePath string
	TemplateFilePath string
	Fs string
	Es string
}

func GetParamsAndValidate() (*ReplacerParams, error) {
	sourceFilePath := flag.String("sfp", "", "source file")
	templateFilePath := flag.String("tfp", "", "template file")
	fs := flag.String("fs", "⟦⟦", "start symbol")
	es := flag.String("es", "⟧⟧", "end symbol")
	flag.Parse()

	if *sourceFilePath == "" || *templateFilePath == "" {
		return nil, errors.New("source (sfp) file and template (tfp) file path are required")
	}

	return &ReplacerParams{
		SourceFilePath: *sourceFilePath,
		TemplateFilePath: *templateFilePath,
		Fs: *fs,
		Es: *es,
	}, nil
}