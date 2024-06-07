package game

type SetNewGameRequest struct {
	GameID string `json:"game_id,omitempty"`
	XSize  int    `json:"x_size,omitempty"`
	YSize  int    `json:"y_size,omitempty"`
}

type SetNewGameResponse struct {
	Status   string      `json:"status,omitempty"`
	Message  string      `json:"message,omitempty"`
	PlayData [][]float64 `json:"play_data,omitempty"`
}

type PlaceStoneRequest struct {
}
