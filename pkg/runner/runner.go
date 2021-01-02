package runner

import (
	"fmt"
	"go_psum/pkg/log"
)

func Run(nameRaw string, exclude string, showDetail int, verbose int) {
	var log = func(text string) {
		if verbose > 0 {
			log.DefaultLogger.Println(text)
		}
	}

	log(fmt.Sprintf("nameRaw:[%s], exclude:[%s], showDetail:[%d]", nameRaw, exclude, showDetail))
}
