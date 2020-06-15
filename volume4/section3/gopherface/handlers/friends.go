package handlers

import (
	"net/http"

	"github.com/richardzhang0301/isokit"

	"github.com/EngineerKamesh/gofullstack/volume4/section3/gopherface/common"
)

func FriendsHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
