package assets

import "io"

// newMultiReadCloser returns an io.ReadCloser that combines all the io.ReadCloser
// arguments supplied to the function.
func newMultiReadCloser(readClosers ...io.ReadCloser) io.ReadCloser {
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

// Read reads, in sequence, from all the ReadClosers contained within
// the MultiReadCloser.
func (mrc *multiReadCloser) Read(p []byte) (n int, err error) {
	return mrc.multiReader.Read(p)
}

// Close closes all ReadClosers contained within the MultiReadCloser.
func (mrc *multiReadCloser) Close() error {
	var err error
	for _, closer := range mrc.readClosers {
		e := closer.Close()
		if e != nil {
			err = e
		}
	}
	return err
}
