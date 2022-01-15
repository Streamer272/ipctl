package handle_error

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		err := fmt.Errorf("%v\n", err)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		os.Exit(1)
	}
}
