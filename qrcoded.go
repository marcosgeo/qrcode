package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code!")

	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

// GenerateQRCode generates QR codes
func GenerateQRCode(toqrcode string) []byte {
	fmt.Println(toqrcode)

	return nil
}
