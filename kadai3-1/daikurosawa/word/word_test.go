package word_test

import (
	"github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/word"
	"testing"
)

func TestWordFile_Generate(t *testing.T) {
	word := word.NewExportWordFile("./../testdata/word.txt")
	if err := word.Generate(); err != nil {
		t.Error("failed generate word.", err)
	}
	words := word.ExportWords()
	if len(words) == 0 {
		t.Fatalf("failed word was not append.")
	}
}

func TestWordFile_Generate_FileOpenError(t *testing.T) {
	word := word.NewExportWordFile("no such file")
	if err := word.Generate(); err == nil {
		t.Fatal("file open error is nothing.", err)
	}
}

func TestWordFile_GetWord(t *testing.T) {
	word := word.NewExportWordFile("./../testdata/word.txt")
	if err := word.Generate(); err != nil {
		t.Error("failed generate word.", err)
	}
	words := word.ExportWords()
	if len(words) == 0 {
		t.Error("word was not append.")
	}
	
	actual, err := word.GetWord()
	if err != nil {
		t.Error("failed get word.", err)
	}
	if len(actual) == 0 {
		t.Fatal("get word is empty.")
	}
}

func TestWordFile_GetWord_EmptyWords(t *testing.T) {
	word := word.NewExportWordFile("dummy")
	actual, err := word.GetWord();
	if err == nil {
		t.Error("word empty error is nothing.")
	}
	if err.Error() != "word was not found" {
		t.Fatalf("failed different error message: %s", err.Error())
	}
	if len(actual) != 0 {
		t.Fatal("get word is not empty.")
	}
}
