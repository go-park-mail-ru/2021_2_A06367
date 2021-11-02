package delivery

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"net/http"
)

type SearchHandler struct {
	fu films.FilmsUsecase
	pu auth.AuthUsecase
	au actors.ActorsUsecase
}

func NewSearchHandler(fu films.FilmsUsecase, pu auth.AuthUsecase, au actors.ActorsUsecase) *SearchHandler {
	return &SearchHandler{fu: fu, pu: pu, au: au}
}

func (sh *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	// Агрегирующий поиск, собирает информацию в один запрос
	// Использует неточный поиск по полям

	keyword := r.URL.Query().Get("search")
	result := models.SearchResult{}

	films, _ := sh.fu.GetByKeyword(keyword)
	result.Films = films

	profiles, _ := sh.pu.GetByKeyword(keyword)
	result.Profiles = profiles

	util.Response(w, models.Okey, result)
}
