package main
//!Example url: http://localhost:8081/generate?t=qr&data=434312344367&h=400&w=200
import (
	"net/http"
	"log"
	"strconv"
	"image/png"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/aztec"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/code93"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/twooffive"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/codabar"
)
func qrCode(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    t := query.Get("t") 
    wid := query.Get("w") 
    h := query.Get("h") 
    data := query.Get("data") 
    if len(t) == 0 {
        w.Write( []byte("Please enter a type") )
		w.WriteHeader( 500 )
		return
    } else {
		marksw := wid
		marksh := h
		var width int
		var height int
			width, werr := strconv.Atoi(marksw)
			if werr!=nil && marksw!="" {
				w.Write( []byte(werr.Error()) )
				w.WriteHeader( 500 )
				return
			}
			height, herr := strconv.Atoi(marksh)
			if herr!=nil && marksh!="" {
				w.Write( []byte(herr.Error()) )
				w.WriteHeader( 500 )
				return
			}
			if width == 0 {
				width=height
			}
			if height == 0 {
				height=width
			}
		if marksh=="" && marksw==""{
			width=200
			height=200
		}
		switch  t{
		case "qr":
			qrCode, _ := qr.Encode(data, qr.M, qr.Auto)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return
		case "code128":
			qrCode, _ := code128.EncodeWithoutChecksum(data)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return
		case "aztec":
			qrCode, _ := aztec.Encode([]byte(data), 1, 1)
			qrCode, _ = barcode.Scale(qrCode, width, height)
 			png.Encode(w, qrCode)
			return
		case "codabar":
			//! Url: http://localhost:8081/generate?t=codabar&data=A80186B
			qrCode, _ := codabar.Encode(data)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return
		case "code39":
			qrCode, _ := code39.Encode(data, true, true)
 			png.Encode(w, qrCode)
			return
		case "code93":
			qrCode, _ := code93.Encode(data, true, true)
 			png.Encode(w, qrCode)
			return
		case "datamatrix":
			qrCode, _ := datamatrix.Encode(data)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return
		case "pdf417":
			qrCode, _ := pdf417.Encode(data,1)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return
		case "2of5":
			qrCode, _ := twooffive.Encode(data,false)
			qrCode, _ = barcode.Scale(qrCode, width, height)
			png.Encode(w, qrCode)
			return

		case "ean":
			_, err := strconv.Atoi(data)
			if err==nil && (len(data)==7 || len(data)==12){
				var qrCode barcode.Barcode
				qrCode, _ = ean.Encode(data)
				qrCode, _ = barcode.Scale(qrCode, width, height)
				png.Encode(w, qrCode)
				return
			}
			w.Write( []byte("Please enter a 8-digitnumber") )
			w.WriteHeader( 500 )
			return
		default: 
		w.Write( []byte("Please enter a valid type!") )
		w.WriteHeader( 500 )
		return
	}
	}
}

func handleRequest(){
	http.HandleFunc("/generate",qrCode)
	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main(){
	handleRequest();
}