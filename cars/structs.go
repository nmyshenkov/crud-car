package cars

// Car - структура с машиной
type Car struct {
	ID         int64  `json:"id,omitempty"`
	Company    string `json:"company,omitempty"`
	Model      string `json:"model,omitempty"`
	Capacity   int    `json:"capacity,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

// Result - структура с успешным результатом запроса
type Result struct {
	Status string `json:"status"`
	Data   []Car  `json:"data,omitempty"`
}

// Error - структура с ошибкой
type Error struct {
	Error string `json:"error"`
}
