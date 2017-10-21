srv := &http.Server{
	Addr: ":3333",
}
go func() {
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}()

stop := make(chan os.Signal, 1)
signal.Notify(stop, syscall.SIGTERM)
<-stop

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := srv.Shutdown(ctx); err != nil {
	log.Println("server shutdown error:", err)
	return
}
log.Println("server shutdown")