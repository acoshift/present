srv := &http.Server{Addr: ":8080"}

go func() {
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}()

stop := make(chan os.Signal, 1)
signal.Notify(stop, syscall.SIGTERM)
<-stop

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
if err = srv.Shutdown(ctx); err != nil {
	log.Println("server shutdown error;", err)
	return
}
