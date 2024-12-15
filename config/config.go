package config

import (
	"github.com/urfave/cli/v2"

	"github.com/dapplink-labs/wallet-sign-go/flags"
)

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	LevelDbPath     string
	RpcServer       ServerConfig
	CredentialsFile string
	KeyName         string
	HsmEnable       bool
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		LevelDbPath:     ctx.String(flags.LevelDbPathFlag.Name),
		CredentialsFile: ctx.String(flags.CredentialsFileFlag.Name),
		KeyName:         ctx.String(flags.KeyNameFlag.Name),
		HsmEnable:       ctx.Bool(flags.HsmEnable.Name),
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
