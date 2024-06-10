package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

type ctxKey struct{}

/* Define a struct that satisfies HTTP Handler */
type TestHandler struct {
	ctx         context.Context // parent context
	jobDuration time.Duration
}

func (t *TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/* Create a channel to where job will ping when completed + create job */
	jobCh := make(chan struct{})
	go imaginaryJob(t.jobDuration, jobCh)

	/* Send depending on job completion or context cancellation */
	select {
	case <-t.ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Context timed out"))
	case <-jobCh:
		w.Write([]byte(fmt.Sprintf("Job of %s duration completed before context timeout!", t.jobDuration)))
	}
}

/* Executes an imaginary job */
func imaginaryJob(jobDuration time.Duration, ch chan struct{}) {
	time.Sleep(jobDuration)
	ch <- struct{}{}
}

func main() {
	/* A */
	// /* Context with value wrapped */
	// ctx := context.Background()
	// ctx = context.WithValue(ctx, ctxKey{}, "ctxkeyVal")
	// ContextValueFunc(ctx)

	/* B */
	/* Context with timeout signal */
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*3)
	defer cancelFunc()

	th := TestHandler{ctx: ctx, jobDuration: time.Second * 5}
	req := httptest.NewRequest("GET", "/context", io.NopCloser(bytes.NewBufferString("Context test request!")))
	recorder := httptest.NewRecorder()
	th.ServeHTTP(recorder, req)

	resp := recorder.Result()
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()

}

func ContextValueFunc(ctx context.Context) {
	fmt.Printf("\n Unwrapping the value of our context: %s", ctx.Value(ctxKey{}))
}
