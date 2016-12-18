package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"github.com/nfnt/resize"
	"os"
	"image"
	"image/draw"
)

func main() {
	f, _ := os.Create("qrcode.png")
	defer f.Close()

	qrcode, err :=qr. Encode("hello world", qr.L, qr.Auto)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, err = barcode.Scale(qrcode, 300, 300)

		img1,er:=LoadImage("5.png")
		logo,_:=LoadImage("1.png")
		//logo:=image.NewRGBA(image.Rect(175,175,50,50))
		logo=resize.Resize(60,60,logo,resize.Bicubic)
		fmt.Println(img1,er)

		if err != nil {
			fmt.Println(err)
		} else {
			pt:=image.Pt(0,0)
			logoPt:=image.Pt(0,0)

			dst:=image.NewRGBA(img1.Bounds())
			//im
			draw.Draw(dst,img1.Bounds(),img1,pt,draw.Src);
			draw.Draw(dst,image.Rect(0,0,300,300),qrcode,pt,draw.Src);
			draw.Draw(dst,image.Rect(0,0,50,50),logo,logoPt,draw.Src);
			png.Encode(f, dst)
		}
	}
}

func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}
