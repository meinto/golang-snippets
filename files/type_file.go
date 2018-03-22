package files


type File struct{
	path string
}

func (f File) ToString() string {
	return f.path
}