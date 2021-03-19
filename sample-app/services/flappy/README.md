### Inicializar modulo Go apontando para o repo do Flappy
go mod init github.com/hajimehoshi/ebiten/examples/flappy

### Instalar o modulo wasmserve
go get github.com/hajimehoshi/wasmserve@latest

### Servir o jogo na web apontando para o repo do Flappy
wasmserve -http localhost:8081 -tags example  github.com/hajimehoshi/ebiten/examples/flappy