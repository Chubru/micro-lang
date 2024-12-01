package cmd

import (
	"github.com/Chubru/micro-lang/pkg/interpreter"
	parser2 "github.com/Chubru/micro-lang/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var inputFile string
var outputFile string

func Run(inputFile, outputFile string) {
	if inputFile == "" {
		log.Fatal("Error: --input flag is required. Please specify a MicroLang file.")
	}

	input, err := antlr.NewFileStream(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	lexer := parser2.NewMicroLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser2.NewMicroLangParser(stream)
	tree := p.Prog()

	if outputFile != "" {
		log.Printf("Compiling file: %s to %s (compiler not implemented yet)", inputFile, outputFile)
		/*
			compiler := compiler.NewCompiler()
			compiledCode := compiler.Visit(tree)

			err := saveToFile(outputFile, compiledCode.(string))
			if err != nil {
				log.Fatalf("Error saving to file: %s", err)
			}
		*/
	} else {
		interpreter := interpreter.NewInterpreter()
		interpreter.Visit(tree)
	}
}

var RootCmd = &cobra.Command{
	Use:   "micro-lang",
	Short: "MicroLang - a language for micromashinki/noVax and VAX compatibility",
	Long: `MicroLang is a programming language designed for micromashinki/noVax 
and fully compatible with VAX systems. Use this tool to interpret or compile MicroLang programs.`,
	Run: func(cmd *cobra.Command, args []string) {
		Run(inputFile, outputFile)
	},
}

func saveToFile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func init() {
	RootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input MicroLang file (required)")
	RootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to the output file (optional)")
	RootCmd.MarkFlagRequired("input")
}
