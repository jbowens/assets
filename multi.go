package assets

import "io"

func MutliReadCloser(readClosers ...io.ReadCloser) io.ReadCloser {
	readers := make([]io.Reader, len(readClosers))
	for idx, reader := range readClosers {
		readers[idx] = reader
	}

	return &multiReadCloser{
		multiReader: io.MultiReader(readers...),
		readClosers: readClosers,
	}
}

type multiReadCloser struct {
	multiReader io.Reader
	readClosers []io.ReadCloser
}

func (mrc *multiReadCloser) Read(p []byte) (n int, err error) {
	return mrc.multiReader.Read(p)
}

func (mrc *multiReadCloser) Close() error {
	var err error = nil
	for _, closer := range mrc.readClosers {
		e := closer.Close()
		if e != nil {
			err = e
		}
	}
	return err
}
