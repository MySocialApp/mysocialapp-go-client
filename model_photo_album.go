package msa

type PhotoAlbum struct {
	Base
	Name          string     `json:"name"`
	Photos        []Photo    `json:"photos"`
	PreviewPhotos *PhotoBlob `json:"preview_photos"`
}

func (p *PhotoAlbum) AddPhoto(photo Photo) (*Photo, *RestError) {
	// TODO
	return nil, nil
}

func (p *PhotoAlbum) AddPhotoWithFeedResult(photo Photo) (*Feed, *RestError) {
	// TODO
	return nil, nil
}

func (p *PhotoAlbum) Delete() *RestError {
	// TODO
	return nil
}

func (p *PhotoAlbum) Update() *RestError {
	// TODO
	return nil
}


