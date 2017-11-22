package realize

import (
	"os"
	"testing"
	"time"
	"strings"
	"log"
	"bytes"
)

func TestRealize_Stop(t *testing.T) {
	r := Realize{}
	r.exit = make(chan os.Signal, 2)
	r.Stop()
	_, ok := <-r.exit
	if ok != false {
		t.Error("Unexpected error", "channel should be closed")
	}
}

func TestRealize_Start(t *testing.T) {
	r := Realize{}
	go func(){
		time.Sleep(100)
		close(r.exit)
		_, ok := <-r.exit
		if ok != false {
			t.Error("Unexpected error", "channel should be closed")
		}
	}()
	r.Start()
}

func TestRealize_Prefix(t *testing.T) {
	r := Realize{}
	input := "test"
	result := r.Prefix(input)
	if len(result) <= 0 && !strings.Contains(result,input){
		t.Error("Unexpected error")
	}
}

func TestSettings_Read(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	w := LogWriter{}
	input := ""
	int, err := w.Write([]byte(input))
	if err != nil || int > 0{
		t.Error("Unexpected error", err, "string lenght should be 0 instead",int)
	}
}