### Inicializar modulo Go apontando para o repo do Blocks
go mod init github.com/hajimehoshi/ebiten/examples/blocks

### Instalar o modulo wasmserve
go get github.com/hajimehoshi/wasmserve@latest

### Servir o jogo na web apontando para o repo do Blocks
wasmserve -http localhost:8082 -tags example  github.com/hajimehoshi/ebiten/examples/blocks