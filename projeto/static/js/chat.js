document.addEventListener("DOMContentLoaded", function () {
    // Garante que o script só execute após o HTML estar completamente carregado.

    // 1. Configuração e Variáveis
    // ---------------------------------------------------------------------

    // Variável que armazena o setor do usuário (Ex: "dev", "rh"). 
    // Assume-se que esta variável (window.SETOR) é injetada no template HTML (chat.html).
    const setor = window.SETOR; 

    const box = document.getElementById("chat-box");      // A área de exibição das mensagens
    const input = document.getElementById("msg");          // O campo de texto para digitar a mensagem
    const btn = document.getElementById("sendBtn");        // O botão de envio

    // 2. Conexão WebSocket
    // ---------------------------------------------------------------------
    
    // Estabelece a conexão WebSocket.
    // "ws://" é o protocolo WebSocket. window.location.host pega o endereço e porta (Ex: localhost:8080).
    // O endpoint "/ws/" + setor (Ex: "/ws/dev") deve corresponder à rota definida no Gin.
    const socket = new WebSocket("ws://" + window.location.host + "/ws/" + setor);

    // 3. Recebimento de Mensagens
    // ---------------------------------------------------------------------

    // Define o que acontece quando uma mensagem é recebida do servidor (websocket.onmessage).
    socket.onmessage = function (event) {
        const p = document.createElement("p"); // Cria um novo elemento <p> para a mensagem
        p.textContent = event.data;            // Define o texto da mensagem recebida
        box.appendChild(p);                    // Adiciona a mensagem ao chat-box
        
        // Mantém a caixa de chat rolando para a mensagem mais recente
        box.scrollTop = box.scrollHeight; 
    };

    // 4. Lógica de Envio
    // ---------------------------------------------------------------------

    // Função que envia o conteúdo do input via WebSocket.
    function sendMsg() {
        // Verifica se o campo está vazio (após remover espaços em branco)
        if (input.value.trim() === "") return; 
        
        socket.send(input.value); // Envia o texto da mensagem ao servidor
        input.value = "";         // Limpa o campo de entrada após o envio
    }

    // Adiciona um listener ao botão de envio para chamar a função sendMsg.
    btn.addEventListener("click", sendMsg);

    // Adiciona um listener para enviar a mensagem ao pressionar a tecla Enter (keypress).
    input.addEventListener("keypress", function (e) {
        if (e.key === "Enter") sendMsg();
    });
});