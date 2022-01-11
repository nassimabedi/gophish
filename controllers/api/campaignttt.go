package api

import (
	"encoding/json"
	"net/http"
	// "strconv"

	ctx "github.com/gophish/gophish/context"
	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	"fmt"
)


//Begin By Nassim
func (as *Server) Campaignsttt(w http.ResponseWriter, r *http.Request) {
        switch {
        case r.Method == "GET":
                cs, err := models.GetCampaignsttt(ctx.Get(r, "user_id").(int64))
                if err != nil {
                        log.Error(err)
                }
                JSONResponse(w, cs, http.StatusOK)
        //POST: Create a new campaign and return it as JSON
        case r.Method == "POST":
			log.Infof("POSTCampain111111111111111111111111111111111111111111111")
			fmt.Println("POSTCampain111111111111111111111111111111111111111111111")
                c := models.Campaignttt{}
                // Put the request into a campaign
                err := json.NewDecoder(r.Body).Decode(&c)
                if err != nil {
						fmt.Println("err222222222222222222222222222222")
                        JSONResponse(w, models.Response{Success: false, Message: "Invalid JSON structure"}, http.StatusBadRequest)
                        return
                }
				fmt.Println("================>>>>>>")
				fmt.Println(&c)
				fmt.Println("================<<<<<<")
                err = models.PostCampaignttt(&c, ctx.Get(r, "user_id").(int64))
                if err != nil {
						fmt.Println("err 33333333333333333333333333333")
                        JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
                        return
                }
                // If the campaign is scheduled to launch immediately, send it to the worker.
                // Otherwise, the worker will pick it up at the scheduled time
                // if c.Status == models.CampaignInProgress {
                //         go as.worker.LaunchCampaign(c)
                // }
                JSONResponse(w, c, http.StatusCreated)
        }
}
//End by Nassim


// start by Nassim
// Campaign returns details about the requested campaign. If the campaign is not
// valid, APICampaign returns null.
// func (as *Server) Campaignttt(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.ParseInt(vars["id"], 0, 64)
// 	c, err := models.GetCampaignttt(id, ctx.Get(r, "user_id").(int64))
// 	if err != nil {
// 		log.Error(err)
// 		JSONResponse(w, models.Response{Success: false, Message: "Campaign not found"}, http.StatusNotFound)
// 		return
// 	}
// 	switch {
// 	case r.Method == "GET":
// 		JSONResponse(w, c, http.StatusOK)
// 	case r.Method == "DELETE":
// 		err = models.DeleteCampaign(id)
// 		if err != nil {
// 			JSONResponse(w, models.Response{Success: false, Message: "Error deleting campaign"}, http.StatusInternalServerError)
// 			return
// 		}
// 		JSONResponse(w, models.Response{Success: true, Message: "Campaign deleted successfully!"}, http.StatusOK)
// 	}
// }

// End by Nassim



