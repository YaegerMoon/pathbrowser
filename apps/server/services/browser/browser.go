package browser

type BrowserService struct {
	curPath string
}

func New(path string) *BrowserService {
	bs := BrowserService{curPath: path}
	return &bs
}

func (bs *BrowserService) List() []string {
	list := []string{}
	return list
}

func (bs *BrowserService) ChangeDir(path string) {

}

func (bs *BrowserService) RemoveDir(path string) {

}

func (bs *BrowserService) RemoveImage(path string) {

}

func (bs *BrowserService) Move(src string, dest string) {

}
