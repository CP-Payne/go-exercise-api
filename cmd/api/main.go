package api

import (
	"github.com/CP-Payne/exercise/internal/env"
	"go.uber.org/zap"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		env:  env.GetString("ENV", "development"),
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
}
