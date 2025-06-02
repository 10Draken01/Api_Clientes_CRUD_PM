package types


// Respuesta estructurada para APIs
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    *MetaInfo   `json:"meta,omitempty"`
}

type MetaInfo struct {
	Page       int    `json:"page,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Total      int64  `json:"total,omitempty"`
	CacheHit   bool   `json:"cache_hit,omitempty"`
	Source     string `json:"source,omitempty"`
	Timestamp  int64  `json:"timestamp"`
}