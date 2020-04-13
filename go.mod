module github.com/trussworks/ecs-service-logs

go 1.14

require (
	github.com/99designs/aws-vault v4.5.1+incompatible
	github.com/99designs/keyring v1.1.4
	github.com/aws/aws-sdk-go v1.29.21
	github.com/go-ini/ini v1.54.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
)

// Update to ignore compiler warnings on macOS catalina
// https://github.com/keybase/go-keychain/pull/55
// https://github.com/99designs/aws-vault/pull/427
replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
