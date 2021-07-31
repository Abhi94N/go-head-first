package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// func scanDirectory(path string) error {
// 	fmt.Println(path)
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		fmt.Printf("Returining error from ScanDirectory(\"%s\" call\n", path)
// 		return err
// 	}

// 	for _, file := range files {
// 		filePath := filepath.Join(path, file.Name())
// 		if file.IsDir() {
// 			err := scanDirectory(filePath)
// 			if err != nil {
// 				fmt.Printf("Returining error from ScanDirectory(\"%s\" call\n", path)
// 				return err
// 			}
// 		} else {
// 			fmt.Println(filePath)
// 		}
// 	}
// 	return nil
// }

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}

}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
	//since the panic is not an error, print error
	else {
		panic(p)
	}
}

func main() {
	defer reportPanic()
	panic("Panic at the disco")
	scanDirectory("../Arrays")
}
