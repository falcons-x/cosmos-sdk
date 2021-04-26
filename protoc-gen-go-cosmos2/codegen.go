/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"strconv"
	"strings"
)

const (
	contextPackage     = protogen.GoImportPath("context")
	grpcPackage        = protogen.GoImportPath("google.golang.org/grpc")
	cosmosTypesPackage = protogen.GoImportPath("github.com/cosmos/cosmos-sdk/types")
)

// generateFile generates a _grpc.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_cosmos.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(file, g)
	return g
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("const _ = ", grpcPackage.Ident("SupportPackageIsVersion7"))
	g.P()
	for _, service := range file.Services {
		genService(file, g, service)
	}
}

func genService(file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	clientName := service.GoName + "Client"

	g.P("// ", clientName, " is the client API for ", service.GoName, " service.")
	g.P("//")
	g.P("// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.")

	// Client interface.
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.Annotate(clientName, service.Location)
	g.P("type ", clientName, " interface {")
	for _, method := range service.Methods {
		g.Annotate(clientName+"."+method.GoName, method.Location)
		if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
			g.P(deprecationComment)
		}
		g.P(method.Comments.Leading,
			clientSignature(g, method))
	}
	g.P("}")
	g.P()

	// Client structure.
	g.P("type ", unexport(clientName), " struct {")
	g.P("cc ", grpcPackage.Ident("ClientConnInterface"))
	for _, method := range service.Methods {
		g.P("_", method.GoName, " ", cosmosTypesPackage.Ident("Invoker"))
	}
	g.P("}")
	g.P()

	// NewClient factory.
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("func New", clientName, " (cc ", grpcPackage.Ident("ClientConnInterface"), ") ", clientName, " {")
	g.P("return &", unexport(clientName), "{cc:cc}")
	g.P("}")
	g.P()

	var methodIndex int
	// Client method implementations.
	for _, method := range service.Methods {
		// Unary RPC method
		genClientMethod(g, method)
		methodIndex++
	}

	// Server interface.
	serverType := service.GoName + "Server"
	g.P("// ", serverType, " is the server API for ", service.GoName, " service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.Annotate(serverType, service.Location)
	g.P("type ", serverType, " interface {")
	for _, method := range service.Methods {
		g.Annotate(serverType+"."+method.GoName, method.Location)
		if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
			g.P(deprecationComment)
		}
		g.P(method.Comments.Leading,
			serverSignature(g, method))
	}
	g.P("}")
	g.P()

	// Server registration.
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	serviceDescVar := service.GoName + "_ServiceDesc"
	g.P("func Register", service.GoName, "Server(s ", grpcPackage.Ident("ServiceRegistrar"), ", srv ", serverType, ") {")
	g.P("s.RegisterService(&", serviceDescVar, `, srv)`)
	g.P("}")
	g.P()

	// Server handler implementations.
	handlerNames := make([]string, 0, len(service.Methods))
	for _, method := range service.Methods {
		hname := genServerMethod(g, method)
		handlerNames = append(handlerNames, hname)
	}

	// Service descriptor.
	g.P("// ", serviceDescVar, " is the ", grpcPackage.Ident("ServiceDesc"), " for ", service.GoName, " service.")
	g.P("// It's only intended for direct use with ", grpcPackage.Ident("RegisterService"), ",")
	g.P("// and not to be introspected or modified (even as a copy)")
	g.P("var ", serviceDescVar, " = ", grpcPackage.Ident("ServiceDesc"), " {")
	g.P("ServiceName: ", strconv.Quote(string(service.Desc.FullName())), ",")
	g.P("HandlerType: (*", serverType, ")(nil),")
	g.P("Methods: []", grpcPackage.Ident("MethodDesc"), "{")
	for i, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		g.P("{")
		g.P("MethodName: ", strconv.Quote(string(method.Desc.Name())), ",")
		g.P("Handler: ", handlerNames[i], ",")
		g.P("},")
	}
	g.P("},")
	g.P("Metadata: \"", file.Desc.Path(), "\",")
	g.P("}")
	g.P()

	g.P("const (")
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		g.P(service.GoName, method.GoName, "Method = ", strconv.Quote(fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())))
	}
	g.P(")")
	g.P()
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return ""
	}

	s := method.GoName + "(ctx " + g.QualifiedGoIdent(contextPackage.Ident("Context"))
	s += ", in *" + g.QualifiedGoIdent(method.Input.GoIdent)
	s += ", opts ..." + g.QualifiedGoIdent(grpcPackage.Ident("CallOption")) + ") ("
	s += "*" + g.QualifiedGoIdent(method.Output.GoIdent)
	s += ", error)"
	return s
}

func genClientMethod(g *protogen.GeneratedFile, method *protogen.Method) {
	service := method.Parent
	sname := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())

	if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("func (c *", unexport(service.GoName), "Client) ", clientSignature(g, method), "{")
	g.P("if invoker := c._", method.GoName, "; invoker != nil {")
	g.P("var out ", method.Output.GoIdent)
	g.P("err := invoker(ctx, in, &out)")
	g.P("return &out, err")
	g.P("}")
	g.P("if invokerConn, ok := c.cc.(", cosmosTypesPackage.Ident("InvokerConn"), "); ok {")
	g.P("var err error")
	g.P("c._", method.GoName, `, err = invokerConn.Invoker("`, sname, `")`)
	g.P("if err != nil {")
	g.P("var out ", method.Output.GoIdent)
	g.P("err = c._", method.GoName, "(ctx, in, &out)")
	g.P("return &out, err")
	g.P("}")
	g.P("}")
	g.P("out := new(", method.Output.GoIdent, ")")
	g.P(`err := c.cc.Invoke(ctx, "`, sname, `", in, out, opts...)`)
	g.P("if err != nil { return nil, err }")
	g.P("return out, nil")
	g.P("}")
	g.P()
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return ""
	}
	var reqArgs []string
	reqArgs = append(reqArgs, g.QualifiedGoIdent(cosmosTypesPackage.Ident("Context")))
	reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(method.Input.GoIdent))
	ret := "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func genServerMethod(g *protogen.GeneratedFile, method *protogen.Method) string {
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return ""
	}

	service := method.Parent
	hname := fmt.Sprintf("_%s_%s_Handler", service.GoName, method.GoName)

	g.P("func ", hname, "(srv interface{}, ctx ", contextPackage.Ident("Context"), ", dec func(interface{}) error, interceptor ", grpcPackage.Ident("UnaryServerInterceptor"), ") (interface{}, error) {")
	g.P("in := new(", method.Input.GoIdent, ")")
	g.P("if err := dec(in); err != nil { return nil, err }")
	g.P("if interceptor == nil { return srv.(", service.GoName, "Server).", method.GoName, "(", cosmosTypesPackage.Ident("UnwrapSDKContext"), "(ctx), in) }")
	g.P("info := &", grpcPackage.Ident("UnaryServerInfo"), "{")
	g.P("Server: srv,")
	g.P("FullMethod: ", strconv.Quote(fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())), ",")
	g.P("}")
	g.P("handler := func(ctx ", contextPackage.Ident("Context"), ", req interface{}) (interface{}, error) {")
	g.P("return srv.(", service.GoName, "Server).", method.GoName, "(", cosmosTypesPackage.Ident("UnwrapSDKContext"), "(ctx), req.(*", method.Input.GoIdent, "))")
	g.P("}")
	g.P("return interceptor(ctx, in, info, handler)")
	g.P("}")
	g.P()

	return hname
}

const deprecationComment = "// Deprecated: Do not use."

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }
