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
	"net/http"

	"github.com/acm-uiuc/groot/config"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const RecruiterURL string = "http://localhost:4567"

const RecruiterToken string = config.RecruiterToken

//Service Data Type
const RecruiterFormat string = "JSON"

//API Interface
var RecruitersRoutes = RouteCollection{
	Route{
		"GetJobs",
		"GET",
		"/jobs",
		GetJobs,
	},
	Route{
		"CreateJob",
		"POST",
		"/jobs",
		CreateJob,
	},
	Route{
		"ApproveJob",
		"PUT",
		"/jobs/{job_id}/approve",
		ApproveJob,
	},
	Route{
		"DeleteJob",
		"DELETE",
		"/jobs/{job_id}",
		DeleteJob,
	},
	Route{
		"GetRecruiters",
		"GET",
		"/recruiters",
		GetRecruiters,
	},
	Route{
		"RecruiterLogin",
		"POST",
		"/recruiters/login",
		RecruiterLogin,
	},
	Route{
		"CreateRecruiter",
		"POST",
		"/recruiters",
		CreateRecruiter,
	},
	Route{
		"GetRecruiter",
		"GET",
		"/recruiters/{recruiter_id}",
		GetRecruiter,
	},
	Route{
		"RenewRecruiter",
		"PUT",
		"/recruiters/{recruiter_id}/renew",
		RenewRecruiter,
	},
	Route{
		"ResetAllRecruiters",
		"POST",
		"/recruiters/reset",
		ResetAllRecruiters,
	},
	Route{
		"UpdateRecruiter",
		"PUT",
		"/recruiters/{recruiter_id}",
		UpdateRecruiter,
	},
	Route{
		"ResetRecruiter",
		"POST",
		"/recruiters/reset_password",
		ResetRecruiter,
	},
	Route{
		"GetRecruiterInvite",
		"GET",
		"/recruiters/{recruiter_id}/invite",
		GetRecruiterInvite,
	},
	Route{
		"SendRecruiterInvite",
		"POST",
		"/recruiters/{recruiter_id}/invite",
		SendRecruiterInvite,
	},
	Route{
		"ResetRecruiterInvite",
		"POST",
		"/recruiters/invite",
		ResetRecruiterInvite,
	},
	Route{
		"DeleteRecruiter",
		"DELETE",
		"/recruiters/{recruiter_id}",
		DeleteRecruiter,
	},
	Route{
		"GetStudents",
		"GET",
		"/students",
		GetStudents,
	},
	Route{
		"CreateStudent",
		"POST",
		"/students",
		CreateStudent,
	},
	Route{
		"ApproveStudent",
		"PUT",
		"/students/{netid}/approve",
		ApproveStudent,
	},
	Route{
		"GetStudent",
		"GET",
		"/students/{netid}",
		GetStudent,
	},
	Route{
		"DeleteStudent",
		"DELETE",
		"/students/{netid}",
		DeleteStudent,
	},
}

//Route handler
func GetJobs(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ApproveJob(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetRecruiters(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func RecruiterLogin(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func CreateRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func RenewRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ResetAllRecruiters(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func UpdateRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ResetRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func SendRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ResetRecruiterInvite(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteRecruiter(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func ApproveStudent(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, RecruiterURL+r.URL.String(), RecruiterFormat, RecruiterToken, r)
}
