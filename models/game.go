package models

type Game struct {
  ID uint `json:"id" gorm:"primary_key"`
  RoomTag string `json:"room_tag"`
  PlayerLimit string `json:"player_limit"`
}