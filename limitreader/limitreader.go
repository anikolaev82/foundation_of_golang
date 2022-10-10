package limitreader

import (
	"io"
	"log"
)

type LimitedReader struct {
	reader io.Reader
	left   int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	if lr.left == 0 {
		return 0, io.EOF
	}

	if lr.left < len(p) {
		p = p[0:lr.left]
	}
	n, e := lr.reader.Read(p)
	lr.left -= n
	if e != nil {
		log.Fatal(e)
	}

	return n, nil
}

/*
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 10)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}*/
