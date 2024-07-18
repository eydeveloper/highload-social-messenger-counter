package entity

type Message struct {
	ReceiverId string `json:"receiver_id"`
	DialogId   string `json:"dialog_id"`
}
