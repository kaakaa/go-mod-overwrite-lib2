package lib

import "log"

func Call() {
	CallChild()
}

func CallChild() {
	log.Println("This is lib2")
}
