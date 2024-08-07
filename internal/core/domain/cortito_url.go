package domain

type CortitoUrl struct {
	Url	string `json:"url" binding:"required"`
	ShortUrl string `json:"shortUrl"`
}

func (url *CortitoUrl) Shorten() {
	url.ShortUrl = url.Url + "/ACORTADA"
}