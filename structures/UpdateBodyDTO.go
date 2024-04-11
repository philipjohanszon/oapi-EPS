package structures

type UpdateBodyDTO struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Body   string `json:"body"`

	QueryParameters []QueryParameterEntity `json:"queryParameters"`
	Headers         []HeaderEntity         `json:"headers"`
}
