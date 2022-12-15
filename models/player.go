package models

type Player struct {
  ID uint `json:"id" gorm:"primary_key"`
  Username  string `json:"username"`
  Score uint `json:"score"`
	SocketId string `json:"socket_id"`
	GameId string `json:"game_id"`
}