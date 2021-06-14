package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

// Makes the errors more reusable and immutable
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Delete(word string) {
	// Go has a built-in function that works on maps
	delete(d, word)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// You can modify maps without passing as an address to it
// This is because a map value is a pointer to a runtime.hmap structure
// Never initialize an empty map variable, instead initialize an empty map
// like we were doing previously, or use the make keyword
// var dictionary = map[string]string{}
// or
// var dictionary = make(map[string]string)
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
