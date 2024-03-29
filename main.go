package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	n0 := NewNode()
	c0 := NewClient()
	_ = c0.Connect(n0)

	subscribers := []Subscriber{}

	for i := 0; i < 1; i++ {
		c := NewClient()
		err := c.Connect(n0)
		if err != nil {
			log.Fatal("error!", err)
		}

		sub, err := c.Subscribe("/test", func(msg *Message) error {
			return nil
		})
		if err != nil {
			log.Fatal("error subscribing", err)
		}
		subscribers = append(subscribers, sub)
	}

	defer func() {
		for _, sub := range subscribers {
			sub.Close()
		}
	}()

	// pool := pond.New(10, 1000)

	done := make(chan struct{})
	go func() {
		// defer pool.StopAndWait()
		i := 0
		for {
			select {
			case sig := <-interrupt:
				slog.Info("received signal interrupt", "signal", sig)
				done <- struct{}{}
				return
			default:
				if i < 100 {
					start := time.Now()
					for i := 0; i < 5_000_000; i++ {
						// pool.Submit(func() {
						err := c0.Publish("/test", &Message{
							SenderId: c0.Id(),
							Topic:    "/test",
							Data:     []byte("hello from c0"),
						})

						if err != nil {
							log.Fatalf("error: %s", err.Error())
						}
						// })

					}

					dur := time.Since(start)

					fmt.Println("took", dur)
					i++
				}

			}
		}
	}()

	<-done

}
