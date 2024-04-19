package domain

type AddFileRequest struct {
	UserId   int32 `json:"userId"`
	Files    []*File
	OldFiles []string
	Img      *File
	LessonId string `json:"lessonId"`
}

type File struct {
	Name string
	Type string
	Data []byte
	Url  string
}
