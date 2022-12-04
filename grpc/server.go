package grpc

import (
	db "simple_bank/db/sqlc"
	"time"
)


type Server struct {
	config Config
	store db.Store
	// tokenMaker Maker
}

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// type Maker interface {
// 	// CreateToken creates a new token for a specific username and duration
// 	CreateToken(username string, duration time.Duration) (string, *Payload, error)

// 	// VerifyToken checks if the token is valid or not
// 	VerifyToken(token string) (*Payload, error)
// }
