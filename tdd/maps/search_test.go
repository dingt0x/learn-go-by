package main

import "testing"

func TestSearch(t *testing.T) {

    dictionary := Dictionary{
        "test": "this is just a test",
    }
    t.Run("know word", func(t *testing.T) {
        got, err := dictionary.Search("test")
        want := "this is just a test"

        assertNoError(t, err)
        assertString(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        assertError(t, err, ErrNotFound)
    })

}

func TestAdd(t *testing.T) {

    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        err := dictionary.Add("test", "this is just a test")
        assertNoError(t, err)
        want := "this is just a test"
        got, err := dictionary.Search("test")

        assertNoError(t, err)
        assertString(t, got, want)
    })
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}

        err := dictionary.Add(word, definition)
        assertError(t, err, ErrWordExists)

    })

}

func TestUpdate(t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        newDefinition := "new definition"
        dictionary := Dictionary{word: definition}
        err := dictionary.Update(word, newDefinition)
        assertNoError(t, err)
        got, err := dictionary.Search(word)
        assertNoError(t, err)
        assertString(t, got, newDefinition)
    })

    t.Run("new word", func(t *testing.T) {
        word := "test"
        newDefinition := "new definition"
        dictionary := Dictionary{}
        err := dictionary.Update(word, newDefinition)
        assertError(t, err, ErrWordDoesNotExists)
    })

}

func TestDelete(t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}

        errDel := dictionary.Delete(word)

        assertNoError(t, errDel)

        _, errSearch := dictionary.Search(word)
        assertError(t, errSearch, ErrNotFound)

    })
    t.Run("not existing word", func(t *testing.T) {
        word := "test"
        dictionary := Dictionary{}

        err := dictionary.Delete(word)

        assertError(t, err, ErrWordDoesNotExists)

    })

}
func assertString(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q given, %q", got, want, "test")
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one")
    }
    if want != got {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error but didn't want one:", got)
    }
}
