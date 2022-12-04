package createtodo

import (
	"mime/multipart"
)

type Input struct {
	UserId      int
	Content     string
	Media       multipart.File
	MediaHeader *multipart.FileHeader
}
