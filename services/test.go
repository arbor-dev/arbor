package services

import (
  "net/http"
  "github.com/acm-uiuc/groot/proxy"
)

var TestURL string = "http://jsonplaceholder.typicode.com"

var TestRoutes = RouteCollection {
    Route{
        "getAllPosts",
        "GET",
        "/posts",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.GET(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "getPost",
        "GET",
        "/posts/{postId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.GET(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "getPostComments",
        "GET",
        "/post/{postId}/comments",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.GET(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "getComments",
        "GET",
        "/comments?postId={postId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.GET(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "getUser",
        "GET",
        "/post?userId={userId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.GET(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "postPost",
        "POST",
        "/posts",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.POST(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "updatePost",
        "PUT",
        "/posts/{postId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.PUT(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "patchPost",
        "PATCH",
        "/posts/{postId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.PATCH(w, TestURL+r.URL.String(), r)
        },
    },
    Route{
        "deletePost",
        "DELETE",
        "/posts/{postId}",
        func (w http.ResponseWriter, r *http.Request) {
            proxy.DELETE(w, TestURL+r.URL.String(), r)
        },
    },
}
