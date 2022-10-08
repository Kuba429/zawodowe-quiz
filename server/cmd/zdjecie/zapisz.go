package zdjecie

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const FolderOgolny = "zdjecia"
const FolderUzytkPrefix = "uzytk"
const Folder = "./" + FolderOgolny + "/" + FolderUzytkPrefix

func Zapisz(r *http.Request) (string, error) {
	plik, fileHeader, err := r.FormFile("obrazek")
	if err != nil {
		return "", nil
	}
	defer plik.Close()

	if err := os.MkdirAll(Folder, os.ModePerm); err != nil {
		return "", err
	}
	sciezka := fmt.Sprintf("%s/%d%s", Folder, time.Now().UnixMilli(), fileHeader.Filename)
	cel, err := os.Create(sciezka)
	if err != nil {
		return "", err
	}
	defer cel.Close()

	io.Copy(cel, plik)

	sciezka = strings.TrimPrefix(sciezka, ".")
	return sciezka, nil
}
