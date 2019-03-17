package main

import "github.com/Farteen/travelfinance/commands"

func main() {
	commands.ResetMongo("users")
	commands.ResetRedis()
}