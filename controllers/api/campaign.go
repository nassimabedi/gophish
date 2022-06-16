package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"
	ctx "github.com/gophish/gophish/context"
	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//api key : 1a91d4d285ecb704e1c66114e736a0766adc3e7e5f590299b5e0cbbd3bc1d923

//[{template1 Group1} {template2 Group1,group2}]
//&{0 0 campaign6 0001-01-01 00:00:00 +0000 UTC 2022-02-19 16:39:00 +0000 +0000 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 {0 0 template2    0001-01-01 00:00:00 +0000 UTC []} 0 {0 0 landingPage1  false false  0001-01-01 00:00:00 +0000 UTC}  [] [{0 0 Group1 0001-01-01 00:00:00 +0000 UTC []} {0 0 group2 0001-01-01 00:00:00 +0000 UTC []}] [] 0 {0 0  Profile1     false [] 0001-01-01 00:00:00 +0000 UTC}  [{template1 Group1} {template2 Group1,group2}]}
// Campaigns returns a list of campaigns if requested via GET.
// If requested via POST, APICampaigns creates a new campaign and returns a reference to it.
func (as *Server) Campaigns(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		cs, err := models.GetCampaigns(ctx.Get(r, "user_id").(int64))
		if err != nil {
			log.Error(err)
		}
		JSONResponse(w, cs, http.StatusOK)
	//POST: Create a new campaign and return it as JSON
	case r.Method == "POST":
		fmt.Println("PPPPPPPPPPPPPPPPPPPPPPPPPPPPP")
		c := models.Campaign{}
		// Put the request into a campaign
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(r.Body)
		fmt.Println(&c)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Invalid JSON structure"}, http.StatusBadRequest)
			return
		}
		err = models.PostCampaignttt(&c, ctx.Get(r, "user_id").(int64))
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// If the campaign is scheduled to launch immediately, send it to the worker.
		// Otherwise, the worker will pick it up at the scheduled time
		if c.Status == models.CampaignInProgress {
			go as.worker.LaunchCampaign(c)
		}
		JSONResponse(w, c, http.StatusCreated)
	}
}

// //Begin By Nassim
// func (as *Server) Campaignsttt(w http.ResponseWriter, r *http.Request) {
//         switch {
//         case r.Method == "GET":
//                 cs, err := models.GetCampaignsttt(ctx.Get(r, "user_id").(int64))
//                 if err != nil {
//                         log.Error(err)
//                 }
//                 JSONResponse(w, cs, http.StatusOK)
//         //POST: Create a new campaign and return it as JSON
//         case r.Method == "POST":
// 			log.Infof("POSTCampain111111111111111111111111111111111111111111111")
// 		fmt.Println("POSTCampain111111111111111111111111111111111111111111111")
//                 c := models.Campaignttt{}
//                 // Put the request into a campaign
//                 err := json.NewDecoder(r.Body).Decode(&c)
//                 if err != nil {
// 			fmt.Println("err222222222222222222222222222222")
//                         JSONResponse(w, models.Response{Success: false, Message: "Invalid JSON structure"}, http.StatusBadRequest)
//                         return
//                 }
// 				fmt.Println("================>>>>>>")
// 				fmt.Println(&c)
// 				fmt.Println("================<<<<<<")
//                 err = models.PostCampaignttt(&c, ctx.Get(r, "user_id").(int64))
//                 if err != nil {
// 			fmt.Println("err 33333333333333333333333333333")
//                         JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
//                         return
//                 }
//                 // If the campaign is scheduled to launch immediately, send it to the worker.
//                 // Otherwise, the worker will pick it up at the scheduled time
//                 // if c.Status == models.CampaignInProgress {
//                 //         go as.worker.LaunchCampaign(c)
//                 // }
//                 JSONResponse(w, c, http.StatusCreated)
//         }
// }
//End by Nassim

// CampaignsSummary returns the summary for the current user's campaigns
func (as *Server) CampaignsSummary(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		cs, err := models.GetCampaignSummaries(ctx.Get(r, "user_id").(int64))
		if err != nil {
			log.Error(err)
			JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, cs, http.StatusOK)
	}
}

// Campaign returns details about the requested campaign. If the campaign is not
// valid, APICampaign returns null.
func (as *Server) Campaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	c, err := models.GetCampaign(id, ctx.Get(r, "user_id").(int64))
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Campaign not found"}, http.StatusNotFound)
		return
	}
	switch {
	case r.Method == "GET":
		JSONResponse(w, c, http.StatusOK)
	case r.Method == "DELETE":
		err = models.DeleteCampaign(id)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Error deleting campaign"}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{Success: true, Message: "Campaign deleted successfully!"}, http.StatusOK)
	}
}

// start by Nassim
// Campaign returns details about the requested campaign. If the campaign is not
// valid, APICampaign returns null.
func (as *Server) Campaignttt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	c, err := models.GetCampaignttt(id, ctx.Get(r, "user_id").(int64))
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Campaign not found"}, http.StatusNotFound)
		return
	}
	switch {
	case r.Method == "GET":
		JSONResponse(w, c, http.StatusOK)
	case r.Method == "DELETE":
		err = models.DeleteCampaign(id)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Error deleting campaign"}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{Success: true, Message: "Campaign deleted successfully!"}, http.StatusOK)
	}
}

// End by Nassim

// CampaignResults returns just the results for a given campaign to
// significantly reduce the information returned.
func (as *Server) CampaignResults(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	cr, err := models.GetCampaignResults(id, ctx.Get(r, "user_id").(int64))
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Campaign not found"}, http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		JSONResponse(w, cr, http.StatusOK)
		return
	}
}

// CampaignSummary returns the summary for a given campaign.
func (as *Server) CampaignSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	switch {
	case r.Method == "GET":
		cs, err := models.GetCampaignSummary(id, ctx.Get(r, "user_id").(int64))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				JSONResponse(w, models.Response{Success: false, Message: "Campaign not found"}, http.StatusNotFound)
			} else {
				JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
			}
			log.Error(err)
			return
		}
		JSONResponse(w, cs, http.StatusOK)
	}
}

// CampaignComplete effectively "ends" a campaign.
// Future phishing emails clicked will return a simple "404" page.
func (as *Server) CampaignComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	switch {
	case r.Method == "GET":
		err := models.CompleteCampaign(id, ctx.Get(r, "user_id").(int64))
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Error completing campaign"}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{Success: true, Message: "Campaign completed successfully!"}, http.StatusOK)
	}
}


func (as *Server) CampaignSetting(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		fmt.Println("GGGG")
		cs, err := models.GetCampaignSetting()
		if err != nil {
			log.Error(err)
		}
		JSONResponse(w, cs, http.StatusOK)
	//POST: Create a new campaign and return it as JSON
	case r.Method == "POST":
		fmt.Println("=========================>>>>>>>>>>PPPPPPPPPPPPPPPPPPPPPPPPPPPPP")
		c := models.Campaign{}
		// Put the request into a campaign
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(r.Body)
		fmt.Println(&c)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Invalid JSON structure"}, http.StatusBadRequest)
			return
		}
		/*err = models.PostCampaignttt(&c, ctx.Get(r, "user_id").(int64))
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// If the campaign is scheduled to launch immediately, send it to the worker.
		// Otherwise, the worker will pick it up at the scheduled time
		if c.Status == models.CampaignInProgress {
			go as.worker.LaunchCampaign(c)
		}*/
		JSONResponse(w, c, http.StatusCreated)
	}
}
