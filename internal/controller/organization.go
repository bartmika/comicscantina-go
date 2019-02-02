package controller

import (
	// "strconv"
    "fmt"
	"net/http"
	"github.com/go-chi/render"
	"github.com/luchacomics/comicscantina-go/internal/model"
	"github.com/luchacomics/comicscantina-go/internal/model_resource"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
	// "github.com/luchacomics/comicscantina-data/internal/pkg/database"
    // "github.com/luchacomics/comicscantina-go/internal/base/service"
)


func ListOrganizationsFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("TODO: Implement - List"))
}

func RetrieveOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("TODO: Implement - Retrieve"))
}

func CreateOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user").(*model.User)

    // // Take the user POST data and serialize it.
    data := &serializer.OrganizationRequest{}
	if err := render.Bind(r, data); err != nil {
        fmt.Println(err)
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database.
    organization, _ := model_resource.DBNewOrganization(
        data.Name,
        data.Description,
        data.Email,
        user.ID,
    )

    // Take our data and serialize it back into a response object to hand
    // back to the organization.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewOrganizationResponse(organization))
}
