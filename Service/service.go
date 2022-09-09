package main

import server "service/server"

func main() {
	Settings := new(server.SettingsOptions)
	Settings.LoadSettings("./settings/settings.json")

	DataBase := new(server.DataBase)
	DataBase.ConnectToDataBase(&Settings.DataBase)

	DataBaseCache := new(server.DataBaseCache)
	DataBaseCache.LoadCache(DataBase)

	MessageBroker := new(server.MessageBroker)
	MessageBroker.ConnectToMessageBroker(&Settings.MessageBroker)
	MessageBroker.SubscribeMessageBroker(DataBaseCache, DataBase)

	Controllers := new(server.Controllers)
	Controllers.ConnectControllers(&Settings.Service, func() {
		Controllers.ShowOrderByID(DataBaseCache)
		Controllers.CreateNewOrder(MessageBroker, &Settings.Service)
	})

}
