module github.com/a5bbbbb/AITUmoment/core_service

go 1.23.4

require (
	github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen v0.0.0-00010101000000-000000000000
	github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env/v11 v11.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.10.0
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/nats-io/nats.go v1.42.0
	github.com/nats-io/nkeys v0.4.11
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/bytedance/sonic v1.12.5 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/arch v0.12.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen => ../common/pkg/grpc/proto/core_service/gen

replace github.com/a5bbbbb/AITUmoment/common/pkg/grpc/grpcconn => ../common/pkg/grpc/grpcconn

replace github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen => ../common/pkg/grpc/proto/email_service/gen
