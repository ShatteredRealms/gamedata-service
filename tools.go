//go:build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate" //-tags "postgres"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/mitranim/gow"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
	_ "github.com/solo-io/protoc-gen-openapi"
	_ "github.com/spf13/cobra-cli"
	_ "go.uber.org/mock/mockgen"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
