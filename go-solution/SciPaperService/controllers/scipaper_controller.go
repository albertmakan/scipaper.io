package controllers

import (
	"net/http"
	"net/rpc"
	"strings"

	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/helpers"
	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/models"
	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SciPaperController struct {
	sciPaperService *services.SciPaperService
	rpcClient *rpc.Client
}

func NewSciPaperController(sciPaperService *services.SciPaperService) *SciPaperController {
	client, _ := rpc.DialHTTP("tcp", "localhost:4040")
	return &SciPaperController{
		sciPaperService: sciPaperService,
		rpcClient: client,
	}
}

func (spc *SciPaperController) CreateOrUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !spc.isLoggedIn(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		var paper models.Paper
		helpers.ReadJSONBody(r, &paper)
		paper.AuthorID, paper.Author = spc.getName(r)
		var err error
		switch r.Method {
			case http.MethodPost:
				err = spc.sciPaperService.Create(&paper)
			case http.MethodPut:
				err = spc.sciPaperService.Update(&paper)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		helpers.JSONResponse(w, http.StatusOK, nil)
	}
}

func (spc *SciPaperController) GetAllByAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !spc.isLoggedIn(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		un, _ := spc.getName(r)
		helpers.JSONResponse(w, http.StatusOK, spc.sciPaperService.GetAllByAuthorID(un))
	}
}

func (spc *SciPaperController) Publish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !spc.isLoggedIn(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var toPublish struct {PaperID string}
		helpers.ReadJSONBody(r, &toPublish)
		id, _ := primitive.ObjectIDFromHex(toPublish.PaperID)
		authorID, _ := spc.getName(r)
		err := spc.sciPaperService.Publish(id, authorID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		helpers.JSONResponse(w, http.StatusOK, nil)
	}
}

func (spc *SciPaperController) isLoggedIn(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		return false
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	isValid := false
	spc.rpcClient.Call("RPC.IsLoggedIn", tokenString, &isValid)
	return isValid
}

func (spc *SciPaperController) getName(r *http.Request) (username, name string) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		return "", ""
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	name = ""
	spc.rpcClient.Call("RPC.GetName", tokenString, &name)
	f := strings.Fields(name)
	return f[0], strings.Join(f[1:], " ")
}

func (spc *SciPaperController) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.JSONResponse(w, http.StatusOK, "Hello from SciPaperService")
	}
}