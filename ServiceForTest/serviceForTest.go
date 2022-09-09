package main

import server "service/server"

func main() {
	Settings := new(server.SettingsOptions)
	Settings.LoadSettings("./settings/settings.json")

	MessageBroker := new(server.MessageBroker)
	MessageBroker.ConnectToMessageBroker(&Settings.MessageBroker)

	Controllers := new(server.Controllers)
	Controllers.ConnectControllers(&Settings.Service, func() {
		Controllers.CreateNewOrder(MessageBroker, &Settings.Service)
	})
}
