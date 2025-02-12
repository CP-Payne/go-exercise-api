package api

import "go.uber.org/zap"

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
}
