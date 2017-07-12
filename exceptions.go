package main

const ex_usage = 64
const ex_ioerr = 74

// FileSystemNotFoundException thrown when a file system cannot be found.
type FileSystemNotFoundException struct {
	message string
	exitStatus byte
}
func (e *FileSystemNotFoundException) Error() string {
	return e.message
}
func newFileSystemNotFoundException(message string) *FileSystemNotFoundException {
	return &FileSystemNotFoundException{message, ex_usage}
}

// FileSystemException thrown when a file system operation fail.
// It is the general type for file system exceptions.
type FileSystemException struct {
	message string
	exitStatus byte
}
func (e *FileSystemException) Error() string {
	return e.message
}
func newFileSystemException(message string) *FileSystemException {
	return &FileSystemException{message, ex_ioerr}
}


