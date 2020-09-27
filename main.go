package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svc BookService
    svc = NewService(logger)

    var sva AuthorService
    sva = NewAuthorService(logger)


    // svc = loggingMiddleware{logger, svc}
    // svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(svc),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(svc),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(svc),
        decodeUpdateBookRequest,
        encodeResponse,
    )
    http.Handle("/", r)
    http.Handle("/book", CreateBookHandler)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")


    //<<<----------------------------------------------------------------->>>
    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(sva),
        decodeCreateAuthorRequest,
        encodeAuthorResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(sva),
        decodeGetAuthorByIdRequest,
        encodeAuthorResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(sva),
        decodeDeleteAuthorRequest,
        encodeAuthorResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(sva),
        decodeUpdateAuthorRequest,
        encodeAuthorResponse,
    )
    http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")

    // http.Handle("/metrics", promhttp.Handler())
    //logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
	//logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	logger.Log("msg", "HTTP", "addr", ":"+"8080")
    logger.Log("err", http.ListenAndServe(":"+"8080", nil))
}