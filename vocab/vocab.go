package vocab

type Word string

type Lang string

type Entry struct {
	Word   Word
	Remark string
}

type Storage interface {
	Read(Lang, Word) (*Entry, error)
	// ReadChunk ...
	// when no error is returned and empty slice is returned then there are no more entries to read.
	ReadChunk(Lang) ([]Entry, error)
	// ReadRandom returns at max n entries chosen at random from storage.
	ReadRandom(Lang, n int) ([]Entry, error)
	Write(Lang, *Entry) error
}

type Vocabulary struct {
	lang string
	st   Storage
}

func New(l Lang, st Storage) *Vocabulary {
	vc := &Vocabulary{st: st}
	return vc
}

func (vc *Vocabulary) Add(e *Entry) error {
	return nil
}

func (vc *Vocabulary) Delete(w Word) error {
	return nil
}

func (vc *Vocabulary) List(fn func([]Entry, error)) error {
	return nil
}

func (vc *Vocabulary) Random(count int, fn func([]Entry, error)) error {
	return nil
}
