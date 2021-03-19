### Inicializar modulo Go apontando para o repo do 2048
go mod init github.com/hajimehoshi/ebiten/examples/2048

### Instalar o modulo wasmserve
go get github.com/hajimehoshi/wasmserve@latest

### Servir o jogo na web apontando para o repo do 2048
wasmserve -http localhost:8083 -tags example  github.com/hajimehoshi/ebiten/examples/2048