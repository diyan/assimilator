package interfaces

type HTTP struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}
