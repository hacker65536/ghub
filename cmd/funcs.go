package cmd

import "fmt"

func chkE(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Must(s string, err error) string {
	if err != nil {
		chkE(err)
	}
	return s
}
