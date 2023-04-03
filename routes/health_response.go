package routes

type HealthResponse struct {
	Uptime UptimeResponse `json:"uptime"`
}

type UptimeResponse struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}
