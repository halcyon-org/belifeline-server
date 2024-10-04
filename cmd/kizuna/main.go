package main

import "github.com/halcyon-org/kizuna/internal/di"

func main() {
	controllers, err := di.InitializeControllerSet()
	if err != nil {
		panic(err)
	}

	err = controllers.BeLifelineController.Serve()
	if err != nil {
		panic(err)
	}
}
