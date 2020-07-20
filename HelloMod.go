package hellomod

import (
	"errors"
)

// Hello Say hello
func Hello(lang string) error {
	switch lang {
	case "en":
		print("Hello mod new version")
		break
	case "cn":
		print("你好，模块，新版本")
		break
	default:
		return errors.New("Unknown language")
	}
	return nil
}
