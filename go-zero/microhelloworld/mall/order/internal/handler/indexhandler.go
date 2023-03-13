package handler

import (
	"net/http"
	"html/template"
	_"github.com/zeromicro/go-zero/rest/httpx"
	_"order/internal/logic"
	"order/internal/svc"
	_"order/internal/types"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(w, nil)
	}
}
