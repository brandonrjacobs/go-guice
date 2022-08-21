package config

func NewFileLoader(fileType FileType, filePaths ...string) Loader {
	return &simpleFileLoader{fileType: fileType, filesPaths: filePaths}
}
