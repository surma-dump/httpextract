package main

import (
	"http"
	"log"
	"strings"
	"io"
	"os"
	"time"
)

type TeeWriter struct {
	A, B io.Writer
}

func (t *TeeWriter) Write(data []byte) (n int, e os.Error) {
	na, ea := t.A.Write(data)
	nb, eb := t.B.Write(data)

	if na < nb {
		n = na
	} else {
		n = nb
	}

	if eb != nil {
		e = eb
	}
	if ea != nil {
		e = ea
	}
	return
}


func handler(w http.ResponseWriter, req *http.Request) {
	out := io.Writer(w)
	_ = time.LocalTime()
	_ = strings.HasPrefix("a", "b")
	if strings.HasSuffix(req.URL.Path, "stream.php") {
		log.Print("Found stream.php")
		filename := "stream-"+time.LocalTime().String()+".mp3"
		f, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			log.Print("Could not open: ", filename)
			return
		}
		defer f.Close()

		out = &TeeWriter {
			w,
			f,
		}
	}
	c := http.Client{}
	resp, e := c.Do(req)
	if e != nil {
		log.Print("Could not get site: ", e)
		return
	}
	defer resp.Body.Close()

	for key, vals := range resp.Header {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}

	_, e = io.Copy(out, resp.Body)
	if e != nil {
		log.Print("Error while sending data: ", e)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
