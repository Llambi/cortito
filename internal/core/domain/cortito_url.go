package domain

const (
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	aphabetSize = 62
)

type CortitoUrl struct {
	Url      string `json:"url" binding:"required"`
	ShortUrl string `json:"shortUrl"`
}

type FindCortitoUrlRequest struct {
	ShortUrl string `json:"shortUrl" binding:"required"`
}

type CreateCortitoUrlRequest struct {
	Url string `json:"url" binding:"required"`
}

type CortitoUrlResponse struct {
	Url      string `json:"url" binding:"required"`
	ShortUrl string `json:"shortUrl" binding:"required"`
}

func (createRequest *CreateCortitoUrlRequest) Into() CortitoUrl {
	url := createRequest.Url
	cortitoUrl := CortitoUrl{Url: url, ShortUrl: encode([]byte(url))}
	return cortitoUrl
}

func (cortitoUrl *CortitoUrl) IntoCortitoUrlResponse() CortitoUrlResponse {
	return CortitoUrlResponse(*cortitoUrl)
}

func encode(letters []byte) string {
	lettersLen := len(letters) - 1

	if lettersLen < 11 {
		return string(letters)
	}

	center := uint(lettersLen / 2)
	head := letters[:center]
	tail := letters[center:]
	resultLen := max(len(head)-1, len(tail)-1)

	var result []byte
	for i := 0; i <= resultLen; i++ {
		headUInt, tailUInt :=
			getOrDefault(i, head, 0),
			getOrDefault(i, tail, 0)
		letterUInt := headUInt + tailUInt
		result = append(result, alphabet[(letterUInt%aphabetSize)])
	}
	// Volver a empezar
	return encode(result)
}

func getOrDefault(pos int, slice []byte, defaultValue uint) uint {
	sliceLen := len(slice) - 1
	if pos > sliceLen {
		return defaultValue
	} else {
		return uint(slice[pos])
	}
}
