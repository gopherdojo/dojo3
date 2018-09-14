package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/domain"
	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/omikuji"
)

func OmikujiHandler(w http.ResponseWriter, r *http.Request) {
	nowOmikuji := omikuji.Omikuji{}

	res := &domain.Omikuji{Result: nowOmikuji.Run()}
	json.NewEncoder(w).Encode(res)
}
