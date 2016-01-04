package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/api/v1/",
        Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/api/v1/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/api/v1/todos/{todoId}",
        TodoShow,
    },
    Route{
        "TodoCreate",
        "POST",
        "/api/v1/todos",
        TodoCreate,
    },
}
