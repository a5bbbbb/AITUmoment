module github.com/a5bbbbb/AITUmoment/email_service

go 1.23.4

require (
	github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env/v11 v11.3.1
	github.com/nats-io/nats.go v1.42.0
	github.com/nats-io/nkeys v0.4.11
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.72.0
	gopkg.in/mail.v2 v2.3.1
)

require (
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)

replace github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen => ../common/pkg/grpc/proto/email_service/gen
