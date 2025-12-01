package handlers

import (
	"log"
	"net/http"
	"sync" // Pacote para garantir segurança em operações concorrentes (mutexes)

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" // Biblioteca padrão para WebSockets
)

// upgrader: Configuração para fazer o HTTP Request se tornar uma conexão WebSocket.
var upgrader = websocket.Upgrader{
	// CheckOrigin: Permite conexões de qualquer origem em ambiente de desenvolvimento.
	// Em produção, isso deve ser restrito ao domínio do seu aplicativo por segurança.
	CheckOrigin: func(r *http.Request) bool { return true },
}

// rooms: Mapa principal que armazena todas as conexões ativas.
// Estrutura: [nome_do_setor (string)] -> [conexões_ativas (*websocket.Conn)]
var rooms = make(map[string]map[*websocket.Conn]bool)

// roomsMu: Mutex para proteger o mapa 'rooms'.
// Garante que apenas uma goroutine por vez acesse ou modifique o mapa 'rooms' (segurança concorrente).
var roomsMu sync.Mutex

// --- Funções de Gerenciamento de Conexão ---

// addConn: Adiciona uma nova conexão WebSocket ao mapa de salas (rooms).
func addConn(setor string, conn *websocket.Conn) {
	roomsMu.Lock()         // BLOQUEIA o acesso ao mapa
	defer roomsMu.Unlock() // Libera o acesso quando a função terminar

	// Inicializa o mapa de conexões para o setor se ele ainda não existir
	if rooms[setor] == nil {
		rooms[setor] = make(map[*websocket.Conn]bool)
	}
	// Adiciona a nova conexão ao setor
	rooms[setor][conn] = true
}

// removeConn: Remove uma conexão WebSocket do mapa de salas e limpa o setor se estiver vazio.
func removeConn(setor string, conn *websocket.Conn) {
	roomsMu.Lock()
	defer roomsMu.Unlock()

	if rooms[setor] != nil {
		// Remove a conexão específica do setor
		delete(rooms[setor], conn)

		// Se o setor ficar vazio após a remoção, remove o setor inteiro do mapa 'rooms'
		if len(rooms[setor]) == 0 {
			delete(rooms, setor)
		}
	}
}

// --- Funções de Transmissão ---

// broadcastToSetor: Envia uma mensagem para TODAS as conexões ativas dentro de um setor específico.
func broadcastToSetor(setor string, msg []byte) {
	roomsMu.Lock()
	conns := rooms[setor]
	roomsMu.Unlock() // Libera o Mutex o mais rápido possível

	// Itera sobre todas as conexões do setor
	for conn := range conns {
		// Tenta escrever a mensagem (como Texto) na conexão
		err := conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			log.Println("Erro ao enviar msg, removendo conexão:", err)
			conn.Close() // Fecha a conexão com erro
			// Remove a conexão do mapa de forma segura
			removeConn(setor, conn)
		}
	}
}

// --- Handler Principal ---

// ChatWS: Handler que faz o upgrade do HTTP para WebSocket e gerencia o loop de mensagens.
// Esta função é chamada quando a rota WebSocket é acessada (Ex: /ws/chat/:setor).
func ChatWS(c *gin.Context) {
	// 1. Obtém o nome do setor a partir do parâmetro da URL
	setor := c.Param("setor")

	// 2. Faz o upgrade do HTTP para WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade:", err)
		return
	}

	// 3. Adiciona a nova conexão ao setor e loga
	addConn(setor, conn)
	log.Println("Novo usuário no setor:", setor)

	// 4. Loop infinito para ler mensagens
	for {
		// Lê o tipo de mensagem (mt) e o conteúdo (msg) da conexão
		_, msg, err := conn.ReadMessage()

		if err != nil {
			// Se houver erro de leitura, a conexão foi encerrada pelo cliente (fechamento ou erro de rede)
			log.Println("Conexão encerrada:", err)
			removeConn(setor, conn) // Remove a conexão do mapa
			conn.Close()            // Fecha a conexão formalmente
			break                   // Sai do loop 'for'
		}

		// 5. Transmite a mensagem recebida para todos os outros usuários do mesmo setor
		broadcastToSetor(setor, msg)
	}
}
