package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	Source string
	Text string
}

func main(){
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //& graceful app shutdown
	defer stop() 

	logs := make(chan LogEntry) //? what does that channel represent and why even I created it
	var wg sync.WaitGroup //? why does I used both, the channel and sync package 
	sources := []string{"Auth-Service", "Payment-Gateway", "Inventory-API"} //& In a real app, these could be file paths
	fmt.Println("--- Starting Log Watcher (Press Ctrl+C to stop) ---")

	for _, s := range sources {
		wg.Add(1)
		go watchSource(ctx, &wg, s, logs) //? why I had passed the wg pointer right here
	}

	go func(){ //? why does that go rotien separated and what doee it be used for
		wg.Wait()
		close(logs)
	}()

	for entry := range logs { //? how can I loop on a channel 
		filterAndPrint(entry)
	}
}

func watchSource(ctx context.Context, wg *sync.WaitGroup, name string, out chan <- LogEntry){
	defer wg.Done()
	ticker := time.NewTicker(800 * time.Millisecond) //? what is the ticker and what is its usage here 
	defer ticker.Stop()

	for {
		select { //? why I use the selcet here not just a simple if conditions
		case <-ctx.Done(): //? what does ctx.Done() mean and why I stope the function here
			return
		case <-ticker.C: //? what does ticker.C mean and why I execute the code after it
			msg := fmt.Sprintf("Event detected at %s", time.Now().Format("15:04:05")) //? why Sprintf here
			if time.Now().UnixNano()%5 == 0 { //? what does unix time mean and and what does that line check even why that condition here
				msg = "ERROR: Database connection timeout!"
			}
			out <- LogEntry{Source: name, Text: msg}
		}
	}
}

func filterAndPrint(e LogEntry){
	if strings.Contains(e.Text, "ERROR") {
		fmt.Printf("\033[31m[%s] %s\033[0m\n", e.Source, e.Text)
	} else {
		fmt.Printf("[%s] %s\n", e.Source, e.Text)
	}
}