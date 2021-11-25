go test -coverprofile=coverage.out.tmp `(go list  ./... | grep -v /cmd |  grep -v mocks | grep -v ./internal/models | grep -v ./internal/pkg/films/mocks |grep -v ./internal/pkg/auth/delivery/grpc/generated | grep -v ./internal/pkg/films/delivery/grpc/generated)`
cat coverage.out.tmp | grep -v _mock.go  > coverage.out
go tool cover -func=coverage.out
