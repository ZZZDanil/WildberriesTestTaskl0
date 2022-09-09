package service

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}
type Items struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
type OrderTable struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type DataBaseCache struct {
	orderTableCache map[string]OrderTable
}

func (c *DataBaseCache) LoadCache(db *DataBase) {

	c.orderTableCache = make(map[string]OrderTable)

	rows, err := db.connection.Query("select data from order_table_cache;")
	if err == nil {
		for rows.Next() {
			var data string
			err := rows.Scan(&data)
			if err == nil {
				c.AddToLocalCache(data)
			} else {

				log.Print(err)
			}
		}
		defer rows.Close()
	} else {
		log.Print(err)
	}

}

func (c *DataBaseCache) AddToLocalCache(jsonData string) (string, error) {
	r := strings.NewReplacer(
		"\\n", "\n",
		"\\\"", "\"",
		"\t", "",
	)
	jsonData = strings.Trim(r.Replace(jsonData), "\"")

	var order OrderTable
	err := json.Unmarshal([]byte(jsonData), &order)
	if err != nil {
		log.Print(err)
		return jsonData, err
	}
	log.Print("AddToLocalCache: ", order.OrderUID)
	if c.orderTableCache != nil {
		c.orderTableCache[order.OrderUID] = order
		return jsonData, nil
	} else {
		return jsonData, errors.New("empty orderTableCache")
	}
}
func (c *DataBaseCache) GetCache(id string) OrderTable {
	return c.orderTableCache[id]
}
