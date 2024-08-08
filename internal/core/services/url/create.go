package url

import "github.com/Llambi/cortito/internal/core/domain"

func (s Service) Create(createCortitoUrlRequest domain.CreateCortitoUrlRequest) (domain.CortitoUrlResponse, error) {
	cortitoUrl := createCortitoUrlRequest.Into()

	if cortitoUrl, err := s.Repo.Insert(cortitoUrl); err != nil {
		return domain.CortitoUrlResponse{}, err
	} else {
		return cortitoUrl.IntoCortitoUrlResponse(), nil
	}
}
