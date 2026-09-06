package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG118 = []CodeSample{

	{[]string{`
package main

import (
	"context"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = ctx
	go func() {
		child, _ := context.WithTimeout(context.Background(), time.Second)
		_ = child
	}()
}
`}, 2, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func work(ctx context.Context) {
	child, _ := context.WithTimeout(ctx, time.Second)
	_ = child
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func run(ctx context.Context) {
	for {
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func complexInfinite(ctx context.Context, ch <-chan int) {
	_ = ctx
	for {
		select {
		case <-ch:
			time.Sleep(time.Millisecond)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				return
			case <-time.After(time.Millisecond):
			}
		}
	}(ctx)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func work(ctx context.Context) {
	child, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	_ = child
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func forwarded(ctx context.Context) {
	child, cancel := context.WithCancel(ctx)
	_ = child
	cancelCopy := cancel
	defer cancelCopy()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func bounded(ctx context.Context) {
	_ = ctx
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func worker(ctx context.Context, max int) {
	_ = ctx
	i := 0
	for {
		if i >= max {
			break
		}
		time.Sleep(time.Millisecond)
		i++
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func work(ctx context.Context) {
	child, _ := context.WithCancel(ctx)
	_ = child
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func work(ctx context.Context) {
	child, _ := context.WithDeadline(ctx, time.Now().Add(time.Second))
	_ = child
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = ctx
	go func() {
		bg := context.TODO()
		_ = bg
	}()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func handler(r *http.Request) {
	_ = r.Context()
	go func() {
		go func() {
			ctx := context.Background()
			_ = ctx
		}()
	}()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func worker(ctx context.Context) {
	_ = ctx
	go func() {
		newCtx := context.Background()
		_, _ = context.WithTimeout(newCtx, time.Second)
	}()
}
`}, 2, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func consume(ctx context.Context, ch <-chan int) {
	_ = ctx
	for val := range ch {
		_ = val
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func selectLoop(ctx context.Context, ch <-chan int) {
	_ = ctx
	for {
		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func multiContext(ctx context.Context) {
	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	_ = ctx1

	ctx2, _ := context.WithCancel(ctx)
	_ = ctx2
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func createContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		ctx := context.Background()
		_ = ctx
	}()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
	"time"
)

func pollAPI(ctx context.Context) {
	for {
		resp, _ := http.Get("https://api.example.com")
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
	"time"
)

func pollDB(ctx context.Context, db *sql.DB) {
	for {
		db.Query("SELECT 1")
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
	"time"
)

func watchFile(ctx context.Context) {
	for {
		os.ReadFile("config.txt")
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
	"time"
)

func safePoller(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			resp, _ := http.Get("https://api.example.com")
			if resp != nil {
				resp.Body.Close()
			}
		}
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func startWorker(ctx context.Context) {
	go func() {
		newCtx, cancel := context.WithTimeout(context.TODO(), time.Second)
		defer cancel()
		_ = newCtx
	}()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func leakyLoop(ctx context.Context) {
	for i := 0; i < 10; i++ {
		child, _ := context.WithTimeout(ctx, time.Second)
		_ = child
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func properLoop(ctx context.Context) {
	for i := 0; i < 10; i++ {
		child, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_ = child
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func storeCancel(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	_ = cancel
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func interfaceCancel(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	var fn func() = cancel
	defer fn()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func nestedContext(ctx context.Context) {
	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()

	ctx2, _ := context.WithCancel(ctx1)
	_ = ctx2
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func spawnWorkers(ctx context.Context) {
	for {
		go func() {
			time.Sleep(time.Millisecond)
		}()
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
	"time"
)

func deferredWrites(ctx context.Context) {
	for {
		defer func() {
			os.WriteFile("log.txt", []byte("data"), 0644)
		}()
		time.Sleep(time.Second)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"io"
	"time"
)

func readLoop(ctx context.Context, r io.Reader) {
	buf := make([]byte, 1024)
	for {
		r.Read(buf)
		time.Sleep(time.Millisecond)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func fetchWithBreak(ctx context.Context) error {
	client := &http.Client{}
	for i := 0; i < 5; i++ {
		req, _ := http.NewRequest("GET", "https://example.com", nil)
		_, err := client.Do(req)
		if err != nil {
			return err
		}
	}
	return nil
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Job struct {
	cancelFn context.CancelFunc
}

func NewJob(ctx context.Context) *Job {
	childCtx, cancel := context.WithCancel(ctx)
	job := &Job{cancelFn: cancel}
	_ = childCtx
	return job
}

func (j *Job) Close() {
	if j.cancelFn != nil {
		j.cancelFn()
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Worker struct {
	cancel context.CancelFunc
}

func NewWorker(ctx context.Context) *Worker {
	childCtx, cancel := context.WithCancel(ctx)
	w := &Worker{cancel: cancel}
	_ = childCtx
	return w
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Service struct {
	stopFn func()
}

func (s *Service) Start(ctx context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	s.stopFn = cancel
	_ = childCtx
}

func (s *Service) Stop() {
	if s.stopFn != nil {
		s.stopFn()
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func conditionalCancel(ctx context.Context, useTimeout bool) {
	var cancel context.CancelFunc
	if useTimeout {
		_, cancel = context.WithCancel(ctx)
	} else {
		_, cancel = context.WithCancel(ctx)
	}
	defer cancel()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func storeAndLoad(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	var holder func()
	holder = cancel
	defer holder()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func changeType(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	fn := (func())(cancel)
	defer fn()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func makeInterface(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	var iface interface{} = cancel
	defer iface.(func())()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Container struct {
	cleanup func()
}

func (c *Container) Setup(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	c.cleanup = cancel
}

func (c *Container) Teardown() {
	if c.cleanup != nil {
		c.cleanup()
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type TaskA struct {
	cancelFn func()
}

type TaskB struct {
	cancelFn func()
}

func NewTaskA(ctx context.Context) *TaskA {
	_, cancel := context.WithCancel(ctx)
	return &TaskA{cancelFn: cancel}
}

func (t *TaskB) Close() {
	if t.cancelFn != nil {
		t.cancelFn()
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type MultiField struct {
	name     string
	cancelFn context.CancelFunc
	data     []byte
}

func (m *MultiField) Init(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	m.cancelFn = cancel
}

func (m *MultiField) Cleanup() {
	if m.cancelFn != nil {
		m.cancelFn()
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func helper(fn func()) {
	defer fn()
}

func useHelper(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	helper(cancel)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func callAsValue(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	(func(f func()) { defer f() })(cancel)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func multiPhiEdges(ctx context.Context, a, b, c bool) {
	var cancel context.CancelFunc
	if a {
		_, cancel = context.WithCancel(ctx)
	} else if b {
		_, cancel = context.WithCancel(ctx)
	} else if c {
		_, cancel = context.WithCancel(ctx)
	} else {
		_, cancel = context.WithCancel(ctx)
	}
	defer cancel()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Outer struct {
	inner Inner
}

type Inner struct {
	cancel func()
}

func (o *Outer) Setup(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	o.inner.cancel = cancel
}

func (o *Outer) Teardown() {
	if o.inner.cancel != nil {
		o.inner.cancel()
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func pollWithInterface(ctx context.Context, client HTTPClient) {
	for {
		req, _ := http.NewRequest("GET", "https://example.com", nil)
		client.Do(req)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Sender interface {
	Send(interface{}) error
}

func sendLoop(ctx context.Context, s Sender) {
	for {
		s.Send("data")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Receiver interface {
	Recv() (interface{}, error)
}

func recvLoop(ctx context.Context, r Receiver) {
	for {
		r.Recv()
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

type Querier interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
}

func queryLoop(ctx context.Context, q Querier) {
	for {
		q.QueryContext(ctx, "SELECT 1")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

type Executor interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

func execLoop(ctx context.Context, e Executor) {
	for {
		e.ExecContext(ctx, "UPDATE foo SET bar = 1")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func roundTripLoop(ctx context.Context, rt http.RoundTripper) {
	for {
		req, _ := http.NewRequest("GET", "https://example.com", nil)
		rt.RoundTrip(req)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func headLoop(ctx context.Context) {
	for {
		http.Head("https://example.com")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func postLoop(ctx context.Context) {
	for {
		http.Post("https://example.com", "text/plain", nil)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
	"net/url"
)

func postFormLoop(ctx context.Context) {
	for {
		http.PostForm("https://example.com", url.Values{})
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

func beginLoop(ctx context.Context, db *sql.DB) {
	for {
		db.Begin()
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

func beginTxLoop(ctx context.Context, db *sql.DB) {
	for {
		db.BeginTx(ctx, nil)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
)

func openLoop(ctx context.Context) {
	for {
		os.Open("file.txt")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
)

func openFileLoop(ctx context.Context) {
	for {
		os.OpenFile("file.txt", os.O_RDONLY, 0644)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
)

func writeFileLoop(ctx context.Context) {
	for {
		os.WriteFile("file.txt", []byte("data"), 0644)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func withContext(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func deadlineNotCalled(ctx context.Context) {
	deadline := time.Now().Add(time.Hour)
	child, _ := context.WithDeadline(ctx, deadline)
	_ = child
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"net/http"
)

func useRequestContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	_ = child
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func withDoneCheck(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond):
		}
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func workerWithBackground(ctx context.Context) {
	go func() {
		bg := context.Background()
		time.Sleep(time.Second)
		_ = bg
	}()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func usesBackground() {
	ctx := context.Background()
	_ = ctx
}

func launchWorker(ctx context.Context) {
	go usesBackground()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func boundedSleep(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func loopWithBreak(ctx context.Context) {
	count := 0
	for {
		time.Sleep(time.Millisecond)
		count++
		if count > 100 {
			break
		}
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func emptyFunc(ctx context.Context) {
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func simpleHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func multipleGoroutines(ctx context.Context) {
	go func() {
		bg1 := context.Background()
		time.Sleep(time.Millisecond)
		_ = bg1
	}()
	go func() {
		bg2 := context.TODO()
		time.Sleep(time.Millisecond)
		_ = bg2
	}()
}
`}, 2, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func spawnWithBg(ctx context.Context) {
	bg := context.Background()
	go func(c context.Context) {
		_ = c
	}(bg)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func singleBlockLoop(ctx context.Context) {
	for i := 0; i < 5; i++ {
		_ = i
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type CancelFunc func()

func convertCancel(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	converted := CancelFunc(cancel)
	defer converted()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"io"
)

func readLoop(ctx context.Context, r io.Reader) {
	buf := make([]byte, 1024)
	for {
		r.Read(buf)
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"io"
)

func writeLoop(ctx context.Context, w io.Writer) {
	for {
		w.Write([]byte("data"))
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func noIssues(ctx context.Context) {
	_ = ctx
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

func queryInLoop(ctx context.Context, db *sql.DB) {
	for {
		rows, _ := db.Query("SELECT * FROM users")
		if rows != nil {
			rows.Close()
		}
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
)

func execInLoop(ctx context.Context, db *sql.DB) {
	for {
		db.Exec("UPDATE users SET active = 1")
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func worker(ctx context.Context) {
	defer time.Sleep(time.Second)
	// work...
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

type Job struct {
	cancel context.CancelFunc
}

func (j *Job) Start(ctx context.Context) {
	childCtx, cancel := context.WithTimeout(ctx, time.Second)
	j.cancel = cancel
	_ = childCtx
}

func (j *Job) Stop() {
	if j.cancel != nil {
		j.cancel()
	}
}

func run(ctx context.Context) {
	job := &Job{}
	job.Start(ctx)
	job.Stop()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

type Task struct {
	cancelFn context.CancelFunc
}

func (t *Task) Execute(ctx context.Context) {
	childCtx, cancel := context.WithTimeout(ctx, time.Second)
	t.cancelFn = cancel
	_ = childCtx
}

func run(ctx context.Context) {
	task := &Task{}
	task.Execute(ctx)
	// Never calls task.cancelFn()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func multipleViolations(ctx context.Context) {
	_, cancel1 := context.WithTimeout(ctx, time.Second)
	_, cancel2 := context.WithTimeout(ctx, time.Second)
	_, cancel3 := context.WithTimeout(ctx, time.Second)
	_, _, _ = cancel1, cancel2, cancel3
}
`}, 3, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"database/sql"
	"fmt"
)

type Env struct {
	DB       *sql.DB
	Shutdown func()
}

func withContext(ctx context.Context, env *Env) error {
	db, closeFn, err := initDatabase(ctx)
	if err != nil {
		return fmt.Errorf("creating database: %w", err)
	}

	prev := env.Shutdown
	env.Shutdown = func() {
		prev()
		closeFn()
	}

	env.DB = db
	return nil
}

func initDatabase(ctx context.Context) (*sql.DB, func(), error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	_ = ctx

	db, err := sql.Open("sqlite", "testing")
	if err != nil {
		return nil, cancelFunc, fmt.Errorf("opening database: %%w", err)
	}
	return db, cancelFunc, nil
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	go func() {
		cancel()
	}()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
)

type Foo struct {
	Cancel func()
}

func NewFoo() Foo {
	_, cancel := context.WithCancel(context.Background())
	foo := Foo{Cancel: cancel}
	return foo
}

func main() {
	foo := NewFoo()
	foo.Cancel()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type State struct {
	done context.CancelFunc
}

func manage(ctx context.Context) {
	s := &State{}
	_, s.done = context.WithCancel(ctx)
	defer s.done()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

type Runner struct {
	stop context.CancelFunc
}

func launch(ctx context.Context) {
	r := &Runner{}
	_, r.stop = context.WithCancel(ctx)
	defer func() {
		r.stop()
	}()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func shutdown() {
	cancel()
}

func main() {
	shutdown()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func handleSignal() {
	cancel()
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		handleSignal()
	}()
	select {}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func main() {
	// Never calls cancel()
	select {}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func cleanup() {
	defer cancel()
}

func main() {
	defer cleanup()
	select {}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func background() {
	go func() {
		cancel()
	}()
}

func main() {
	background()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel1, cancel2 context.CancelFunc

func init() {
	ctx1, c1 := context.WithCancel(context.Background())
	cancel1 = c1
	_ = ctx1

	ctx2, c2 := context.WithCancel(context.Background())
	cancel2 = c2
	_ = ctx2
}

func shutdown() {
	cancel1()
}

func main() {
	shutdown()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func setup() {
	defer func() {
		cancel()
	}()
}

func main() {
	setup()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

type App struct{}

func init() {
	ctx, c := context.WithCancel(context.Background())
	cancel = c
	_ = ctx
}

func (a *App) Stop() {
	cancel()
}

func main() {
	app := &App{}
	defer app.Stop()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	_, cancel = context.WithCancel(context.Background())
}

func invoke(fn func()) {
	fn()
}

func execute() {
	invoke(cancel)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

var cancel context.CancelFunc

func init() {
	_, cancel = context.WithCancel(context.Background())
}

func setup() {
	defer func() {
		defer cancel()
	}()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"context"
	"time"
)

func main() {
	defers := make([]func(), 0, 1)
	defer func() {
		for _, fn := range defers {
			fn()
		}
	}()

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defers = append(defers, cancel)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func cleanups(ctx context.Context) []func() {
	var fns []func()
	_, cancel := context.WithCancel(ctx)
	fns = append(fns, cancel)
	return fns
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func register(ctx context.Context, key string) map[string]context.CancelFunc {
	m := make(map[string]context.CancelFunc)
	_, cancel := context.WithCancel(ctx)
	m[key] = cancel
	return m
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "context"

func arrayStore(ctx context.Context) [1]context.CancelFunc {
	var arr [1]context.CancelFunc
	_, cancel := context.WithCancel(ctx)
	arr[0] = cancel
	return arr
}
`}, 0, gosec.NewConfig()},
}
