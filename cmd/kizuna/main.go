package main

import "github.com/halcyon-org/kizuna/internal/di"

func main() {
	controllers, err := di.InitializeControllerSet()
	if err != nil {
		panic(err)
	}

	controllers.BeLifelineController.Serve()
}
