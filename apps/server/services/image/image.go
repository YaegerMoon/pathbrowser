package image

type ImageService struct {
	tileSize int64
}

func New(tileSize int64) *ImageService {
	return &ImageService{tileSize: tileSize}
}

func (service *ImageService) FindAll(path string) {

}

func (service *ImageService) FindOne(file string) {

}

func (service *ImageService) FindDZI(path string) {

}

func (service *ImageService) ReadTile(path string, col int64, row int64, level int32) {

}
