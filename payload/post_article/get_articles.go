package postarticle_payload

type GetArticlesPaylod struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
