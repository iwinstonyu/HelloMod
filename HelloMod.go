package HelloMod

import (
	"errors"
)

func Hello(lang string){
	switch(lang){
	case "en":
		print("Hello mod new version")
	case "cn":
		print("你好，模块，新版本")
	default:
		errors.New("Unknown language")
	}
}
