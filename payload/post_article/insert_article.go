package postarticle_payload

type Status string

const (
	Publish = "publish"
	Draft   = "draft"
	Thrash  = "thrash"
)

func (s Status) IsValid() bool {
	switch s {
	case Publish, Draft, Thrash:
		return true
	}

	return false
}

type PostArticlePayload struct {
	Title    string `json:"title" binding:"required,gte=20"`
	Content  string `json:"content" binding:"required,gte=100"`
	Category string `json:"category" binding:"required,gte=3"`
	Status   Status `json:"status" binding:"required,enum"`
}
