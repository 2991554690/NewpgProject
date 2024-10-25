package main

import "pg-manager/router"

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}

}
