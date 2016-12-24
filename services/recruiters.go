/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package services

import (
        "fmt"
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const RecruiterURL string = "http://localhost:4567"

const RecruiterToken string = ""

//Service Data Type
const RecruiterFormat string = "JSON"

//API Interface
var RecruitersRoutes = RouteCollection{
    Route {
    	"GetJobs",
    	"GET",
    	"/jobs",
    	GetJobs,
    },
    Route {
    	"NewJob",
    	"POST",
    	"/jobs",
    	NewJob,
    },
    Route {
    	"UpdateJob",
    	"PUT",
    	"/job/status",
    	UpdateJob,
    },
    Route {
    	"DeleteJob",
    	"DELETE",
    	"/jobs",
    	DeleteJob,
    },
    Route {
    	"RecruiterLogin",
    	"GET",
    	"/recruiters/login",
    	RecruiterLogin,
    },
    Route {
    	"NewRecruiter",
    	"POST",
    	"/recruiters/new",
    	NewRecruiter,
    },
    Route {
    	"GetUnapprovedResumes",
    	"GET",
    	"/resumes/unapproved",
    	GetUnapprovedResumes,
    },
    Route {
    	"NewResume",
    	"POST",
    	"/resumes",
    	NewResume,
    },
    Route {
    	"ApproveResume",
    	"PUT",
    	"/resumes/approve",
    	ApproveResume,
    },
    Route {
    	"DeleteResume",
    	"DELETE",
    	"/resumes/approve",
    	DeleteResume,
    },
    Route {
    	"GetUsers",
    	"GET",
    	"/users",
    	GetUsers,
   	},
    Route {
    	"GetUser",
    	"GET",
    	"/users/{netid}",
    	GetUser,
    },
    Route {
    	"UpdateUser",
    	"PUT",
    	"/users/{netid}",
    	UpdateUser,
    },
    Route {
    	"DeleteUser",
    	"DELETE",
    	"/users/{netid}",
    	DeleteUser,
    },
    Route {
    	"SearchUsers",
    	"GET",
    	"/users/search",
    	SearchUsers,
    },
}

//Route handler
func GetJobs(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func NewJob(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func RecruiterLogin(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func NewRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetUnapprovedResumes(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func NewResume(w http.ResponseWriter, r *http.Request) {
        fmt.Println(*r)
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ApproveResume(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteResume(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}
