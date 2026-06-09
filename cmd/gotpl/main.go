package main

import (
	"fmt"

	"github.com/spyatachkov/gotpl/internal/reader"
	"github.com/spyatachkov/gotpl/internal/replacer"
	"github.com/spyatachkov/gotpl/internal/validator"
	"github.com/spyatachkov/gotpl/utils"
)

func main() {
	params, err := validator.GetParamsAndValidate()
	if err != nil {
		fmt.Println(err)
		return
	}

	templateLines, err := reader.ReadFileByLine(params.TemplateFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tplMap, err := utils.CreateMapFromLinesBySeparator(templateLines, ":")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := reader.ReadFileByFull(params.SourceFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	replaced := replacer.Process(tplMap, file, params.Fs, params.Es)
	fmt.Println(replaced)
}
