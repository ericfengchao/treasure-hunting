package player_service

type PlayerService interface {
	Start(chan<- struct{})
	StartHeartbeat()
}
