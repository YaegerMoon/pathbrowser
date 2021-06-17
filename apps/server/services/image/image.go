package image

import "github.com/jammy-dodgers/gophenslide/openslide"

type ImageService struct {
	path   string
	format string
	size   string
	slide  openslide.Slide
}
