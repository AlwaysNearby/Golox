package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Printf("Usage: generate_ast <output directory>\n")
		os.Exit(64)
	}

	outputDir := os.Args[1]
	defineAst(outputDir, "Expr", []string{
		"Binary:Left Expr, Operator Token, Right Expr",
		"Grouping:Expression Expr",
		"Literal:Value interface{}",
		"Unary:Operator Token, Right Expr",
	})
}

func defineAst(outputDir, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName[0:1]) + baseName[1:] + ".go"
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("unable create file %s, because: %s", path, err.Error())
		os.Exit(73)
	}

	buffer := bufio.NewWriter(file)
	buffer.WriteString("package golox\n")
	buffer.WriteString("\n")

	defineVisitor(buffer, baseName, types)

	buffer.WriteString("type " + baseName + " interface {\n")
	buffer.WriteString("	Accept(" + "visitor " + "Visitor" + baseName + ")" + " any\n")
	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	for i, t := range types {
		structType := strings.Split(t, ":")[0]
		fields := strings.Split(t, ":")[1]
		defineType(buffer, structType, baseName, fields, i != len(types)-1)
	}
	buffer.Flush()
	file.Close()
}

func defineType(buffer *bufio.Writer, structType, baseTypeName, filedsList string, newLine bool) {
	buffer.WriteString("type " + structType + " struct {\n")

	fileds := strings.Split(filedsList, ", ")
	for _, filed := range fileds {
		buffer.WriteString("	" + filed + "\n")
	}

	buffer.WriteString("}\n\n")
	buffer.WriteString("func (" + strings.ToLower(structType[0:1]) + " " + structType + ") " + "Accept(visitor Visitor" + baseTypeName + ") any {\n")
	buffer.WriteString("	return visitor.Visit" + structType + "(" + strings.ToLower(structType[0:1]) + ")\n")
	buffer.WriteString("}\n")

	if newLine {
		buffer.WriteString("\n")
	}
}

func defineVisitor(buffer *bufio.Writer, baseName string, types []string) {
	buffer.WriteString("type Visitor" + baseName + " interface {\n")

	for _, t := range types {
		typeName := strings.Split(t, ":")[0]
		buffer.WriteString("	Visit" + typeName + "(" + strings.ToLower(baseName) + " " + typeName + ")" + " any\n")
	}

	buffer.WriteString("}\n\n")
}
