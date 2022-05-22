package main

import (
	"bizerror/bizerr"
	"errors"
	"fmt"
	"runtime/debug"
)

// Defines error with bizerr.New

var ErrNoPermission = bizerr.New("usecase.bizerror.code.no_permission")
var ErrFolderNameConflict = bizerr.New("usecase.bizerror.code.folder_name_conflict")
var ErrNetworkUnreachable = bizerr.New("usecase.bizerror.code.network_unreachable")

func foo() error {
	return ErrNoPermission.WithParam("param1", "param2")
}

func fooWithStack() error {
	return fmt.Errorf("%w %s",
		ErrNoPermission.WithParam("param1", "param2"),
		debug.Stack())
}

func main() {
	var err error
	err = foo()
	println("1.1--------------------------------------")
	handleErrorWithBranch(err)
	println("1.2--------------------------------------")
	handleErrorWithSwitch(err)
	println("1.3--------------------------------------")
	fmt.Print(err)
	println(`
-----------------------------------------
-----------------------------------------`)
	err = fooWithStack()
	// println(bizerr.ExtractParams(err))
	println("2.1--------------------------------------")
	handleErrorWithBranch(err)
	println("2.2--------------------------------------")
	handleErrorWithSwitch(err)
	println("2.3--------------------------------------")
	fmt.Print(err)
}

func handleErrorWithBranch(err error) {
	if errors.Is(err, ErrNoPermission) {
		fmt.Printf("biz error\ncode: %s\nparams: %+v\n",
			ErrNoPermission, bizerr.ExtractParams(err))
	} else if errors.Is(err, ErrFolderNameConflict) {
		fmt.Printf("biz error: %s\n", err)
	} else if err != nil {
		fmt.Println("internal error")
	}
}

func handleErrorWithSwitch(err error) {
	var bizErr bizerr.Error
	if errors.As(err, &bizErr) {
		switch bizErr {
		case ErrNoPermission:
			fmt.Printf("biz error\ncode: %s\nparams: %+v\n",
				bizErr, bizerr.ExtractParams(err))
		case ErrNetworkUnreachable:
			fmt.Printf("biz error: %s\n", err)
		case ErrFolderNameConflict:
			fmt.Printf("biz error: %s\n", err)
		default:
			fmt.Println("internal error")
		}
	}
}
