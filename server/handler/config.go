package handler

const (
	StorageRootDirectory = "/C-NetDisk/cloudfile"

	UsernameMaxLength = 15
	UsernameMinLength = 4
	ReadBytesEachTime = 1024
	FilenameMaxLength = 100
	FilenameMinLength = 1
	FilepathMaxLength = 2000
)

var ForbiddenChar = map[byte]bool{
	'.': true, '-': true, '/': true, '\\': true, ':': true, '*': true,
	'?': true, '<': true, '>': true, '|': true, '~': true, '`': true,
	'!': true, '@': true, '#': true, '$': true, '%': true, '^': true,
	'&': true, '(': true, ')': true, '+': true, '[': true, ']': true,
	'{': true, '}': true, ';': true, '\'': true, '"': true, ',': true,
	' ': true,
}
