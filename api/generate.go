//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package api

import (
	"fmt"
	"go/types"
	"io"
	"strings"

	"golang.org/x/tools/go/packages"
)

const cwriTypeName string = "ClientWithResponsesInterface"
const ssiTypeName string = "StrictServerInterface"

func GenerateClientWrapper(out io.StringWriter) error {
	src, err := lookupGenerateTypeName(cwriTypeName)
	if err != nil {
		return err
	}
	generator := &wrapperGenerator{
		out: out,
		src: src,
	}
	return generator.Run()
}

func GenerateMock(out io.StringWriter) error {
	src, err := lookupGenerateTypeName(ssiTypeName)
	if err != nil {
		return err
	}
	generator := &mockGenerator{
		out: out,
		src: src,
	}
	return generator.Run()
}

func GenerateAPITest(out io.StringWriter) error {
	src, err := lookupGenerateTypeName(cwriTypeName)
	if err != nil {
		return err
	}
	generator := &apiTestGenerator{
		out: out,
		src: src,
	}
	return generator.Run()
}

func lookupGenerateTypeName(typeName string) (*types.Interface, error) {
	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(config, ".")
	if err != nil {
		return nil, err
	}
	for _, pkg := range pkgs {
		obj := pkg.Types.Scope().Lookup(typeName)
		if obj == nil {
			continue
		}
		src, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			return nil, fmt.Errorf("type %s is not an interface", typeName)
		}
		return src, nil
	}
	return nil, fmt.Errorf("failed to lookup %s", typeName)
}

type wrapperGenerator struct {
	out io.StringWriter
	src *types.Interface
}

func (g *wrapperGenerator) Run() error {
	g.writePreamble(cwriTypeName)
	for method := range g.src.Methods() {
		signature := method.Type().(*types.Signature)
		g.out.WriteString("func (client *Client) " + stripResponseSuffix(method.Name()) + "(")
		g.writeParams(signature)
		g.writeResults(signature)
		g.out.WriteString(" {\n")
		g.writeBody(method, signature)
		g.out.WriteString("}\n\n")
	}
	return nil
}

func (g *wrapperGenerator) writePreamble(typeName string) {
	g.out.WriteString("// generated from api." + typeName + "\n")
	g.out.WriteString("package openmeteo\n")
	g.out.WriteString("import (\n")
	g.out.WriteString("\"context\"\n")
	//	g.out.WriteString("\"io\"\n")
	g.out.WriteString("\n")
	g.out.WriteString("\"github.com/tdrn-org/go-open-meteo/api\"\n")
	g.out.WriteString(")\n")
}

func (g *wrapperGenerator) writeParams(signature *types.Signature) {
	params := signature.Params()
	for i := 0; i < params.Len()-1; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		param := params.At(i)
		g.out.WriteString(param.Name() + " " + stripApiPackage(param.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *wrapperGenerator) writeResults(signature *types.Signature) {
	results := signature.Results()
	resultsLen := results.Len()
	if resultsLen == 0 {
		return
	}
	g.out.WriteString(" (")
	for i := 0; i < resultsLen; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		result := results.At(i)
		g.out.WriteString(result.Name() + " " + stripApiPackage(result.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *wrapperGenerator) writeBody(method *types.Func, signature *types.Signature) {
	methodName := method.Name()
	params := signature.Params()
	g.out.WriteString("response, err := client.apiClient." + methodName + "(")
	for i := 0; i < params.Len()-1; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		param := params.At(i)
		g.out.WriteString(param.Name())
	}
	if strings.HasSuffix(methodName, "WithFormdataBodyWithResponse") {
		g.out.WriteString(", client.requestContentType(\"multipart/form-data; boundary=----test----\")")
	}
	g.out.WriteString(")\n")
	g.out.WriteString("if err != nil {\n")
	g.out.WriteString("return nil,client.wrapSystemError(err)\n")
	g.out.WriteString("}\n")
	g.out.WriteString("err = client.checkAPIResponse(response.HTTPResponse)\n")
	g.out.WriteString("if err != nil {\n")
	g.out.WriteString("return nil,err\n")
	g.out.WriteString("}\n")
	g.out.WriteString("return response,err\n")
}

type mockGenerator struct {
	out io.StringWriter
	src *types.Interface
}

func (g *mockGenerator) Run() error {
	g.writePreamble(ssiTypeName)
	for method := range g.src.Methods() {
		signature := method.Type().(*types.Signature)
		g.out.WriteString("func (s *Server) " + stripResponseSuffix(method.Name()) + "(")
		g.writeParams(signature)
		g.writeResults(signature)
		g.out.WriteString(" {\n")
		g.writeBody(method, signature)
		g.out.WriteString("}\n\n")
	}
	return nil
}

func (g *mockGenerator) writePreamble(typeName string) {
	g.out.WriteString("// generated from api." + typeName + "\n")
	g.out.WriteString("package mock\n")
	g.out.WriteString("import (\n")
	g.out.WriteString("\"context\"\n")
	g.out.WriteString("\n")
	g.out.WriteString("\"github.com/tdrn-org/go-open-meteo/api\"\n")
	g.out.WriteString(")\n")
}

func (g *mockGenerator) writeParams(signature *types.Signature) {
	params := signature.Params()
	for i := 0; i < params.Len(); i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		param := params.At(i)
		g.out.WriteString(param.Name() + " " + stripApiPackage(param.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *mockGenerator) writeResults(signature *types.Signature) {
	results := signature.Results()
	resultsLen := results.Len()
	if resultsLen == 0 {
		return
	}
	g.out.WriteString(" (")
	for i := 0; i < resultsLen; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		result := results.At(i)
		g.out.WriteString(result.Name() + " " + stripApiPackage(result.Type().String()))
	}
	g.out.WriteString(")")
}

func (g *mockGenerator) writeBody(method *types.Func, signature *types.Signature) {
	g.out.WriteString("return nil,nil\n")
	// params := signature.Params()
	// g.out.WriteString("response, err := client.apiClient." + method.Name() + "(")
	// for i := 0; i < params.Len()-1; i++ {
	// 	if i > 0 {
	// 		g.out.WriteString(", ")
	// 	}
	// 	param := params.At(i)
	// 	g.out.WriteString(param.Name())
	// }
	// g.out.WriteString(")\n")
	// g.out.WriteString("if err != nil {\n")
	// g.out.WriteString("return nil,client.wrapSystemError(err)\n")
	// g.out.WriteString("}\n")
	// g.out.WriteString("err = client.checkAPIResponse(response.HTTPResponse)\n")
	// g.out.WriteString("if err != nil {\n")
	// g.out.WriteString("return nil,err\n")
	// g.out.WriteString("}\n")
	// g.out.WriteString("return response,err\n")
}

type apiTestGenerator struct {
	out io.StringWriter
	src *types.Interface
}

func (g *apiTestGenerator) Run() error {
	g.writePreamble(cwriTypeName)
	g.out.WriteString("func TestAPI(t *testing.T) {\n")
	g.out.WriteString("mockServer := mock.Start()\n")
	g.out.WriteString("defer mockServer.Stop(t.Context())\n\n")
	g.out.WriteString("client, err := openmeteo.NewClient(mockServer.APIURL())\n")
	g.out.WriteString("require.NoError(t,err)\n\n")
	for method := range g.src.Methods() {
		g.out.WriteString("{\n")
		signature := method.Type().(*types.Signature)
		g.writeParamVars(method.Name(), signature)
		g.out.WriteString("_, err := client." + stripResponseSuffix(method.Name()) + "(")
		g.writeParams(signature)
		g.out.WriteString(")\n")
		g.out.WriteString("require.NoError(t,err)\n")
		g.out.WriteString("}\n")
	}
	g.out.WriteString("}\n")
	return nil
}

func (g *apiTestGenerator) writePreamble(typeName string) {
	g.out.WriteString("// generated from api." + typeName + "\n")
	g.out.WriteString("package openmeteo_test\n")
	g.out.WriteString("import (\n")
	g.out.WriteString("\"testing\"\n")
	g.out.WriteString("\n")
	g.out.WriteString("\"github.com/stretchr/testify/require\"\n")
	g.out.WriteString("\"github.com/tdrn-org/go-open-meteo\"\n")
	g.out.WriteString("\"github.com/tdrn-org/go-open-meteo/api\"\n")
	g.out.WriteString("\"github.com/tdrn-org/go-open-meteo/mock\"\n")
	g.out.WriteString(")\n")
}

func (g *apiTestGenerator) writeParamVars(methodName string, signature *types.Signature) {
	params := signature.Params()
	for i := 0; i < params.Len()-1; i++ {
		param := params.At(i)
		paramName := param.Name()
		g.out.WriteString(paramName + " := ")
		g.writeVarInitialization(methodName, paramName, param.Type())
	}
}

const applicationJsonContentType string = "application/json"
const multipartFormDataContentType string = "multipart/form-data; boundary=bndry"

const applicationJsonContent string = "{}"
const multipartFormDataContent string = "--bndry\\r\\nContent-Disposition: form-data; name=\\\"document\\\"\\r\\n\\r\\ndocument\\r\\n--bndry--\\r\\n"

var methodContentTypeOverrideMap map[string]string = map[string]string{
	"DocumentsPostDocumentCreateWithBodyWithResponse": multipartFormDataContentType,
	"DocumentsUpdateVersionWithBodyWithResponse":      multipartFormDataContentType,
}

var methodContentOverrideMap map[string]string = map[string]string{
	"DocumentsPostDocumentCreateWithBodyWithResponse": multipartFormDataContent,
	"DocumentsUpdateVersionWithBodyWithResponse":      multipartFormDataContent,
}

func (g *apiTestGenerator) writeVarInitialization(methodName, varName string, varType types.Type) {
	varTypeString := stripApiPackage(varType.String())
	switch varTypeString {
	case "int":
		g.out.WriteString("0\n")
	case "string":
		switch varName {
		case "contentType":
			contentType, ok := methodContentTypeOverrideMap[methodName]
			if !ok {
				contentType = applicationJsonContentType
			}
			g.out.WriteString("\"" + contentType + "\"\n")
		default:
			g.out.WriteString("\"" + varName + "\"\n")
		}
	case "context.Context":
		g.out.WriteString("t.Context()\n")
	case "io.Reader":
		content, ok := methodContentOverrideMap[methodName]
		if !ok {
			content = applicationJsonContent
		}
		g.out.WriteString("strings.NewReader(\"" + content + "\")\n")
	case "api.ConfigUpdateFormdataRequestBody":
		g.out.WriteString("api.ConfigUpdateFormdataRequestBody{ BarcodeTagMapping : \"{}\", UserArgs: \"[]\" }\n")
	default:
		if strings.HasPrefix(varTypeString, "*") {
			g.out.WriteString("&")
			g.out.WriteString(strings.TrimPrefix(varTypeString, "*"))
		} else {
			g.out.WriteString(varTypeString)
		}
		g.out.WriteString("{}\n")
	}
}

func (g *apiTestGenerator) writeParams(signature *types.Signature) {
	params := signature.Params()
	for i := 0; i < params.Len()-1; i++ {
		if i > 0 {
			g.out.WriteString(", ")
		}
		param := params.At(i)
		g.out.WriteString(param.Name())
	}
}

func stripResponseSuffix(s string) string {
	return strings.TrimSuffix(s, "WithResponse")
}
func stripApiPackage(s string) string {
	return strings.Replace(s, "github.com/tdrn-org/go-open-meteo/", "", 1)
}
