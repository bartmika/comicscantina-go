package controller

import (
    "fmt"
    // "errors"
    "net/http"
    "github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
    "github.com/luchacomics/comicscantina-go/internal/base/service"
)

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.RegistrationRequest{}
	if err := render.Bind(r, data); err != nil {
        fmt.Println(err)
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database.
    user, _ := model_resource.DBNewUser(data.Email, data.Password, data.FirstName, data.LastName)

    // Create the users authentication token.
    tokenString := service.GenerateJWTToken(user.ID)

    // Take our data and serialize it back into a response object to hand
    // back to the user.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewRegistrationResponse(
        tokenString,
        user.ID,
        user.Email,
        user.FirstName,
        user.LastName,
    ))
}
