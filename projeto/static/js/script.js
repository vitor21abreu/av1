const socket = new WebSocket("ws://localhost:8080/ws")
// Estabelece a conexão WebSocket.
// "ws://" é o protocolo (o equivalente a http://).
// A conexão é feita com o servidor Go (Gin) na porta 8080, no endpoint "/ws".

socket.onmessage = (Event) => {
    // 1. Manipulador de Evento: Executado sempre que uma mensagem é recebida do servidor.
    
    const chat = document.getElementById("chat");
    // Seleciona a área de exibição das mensagens no HTML.
    
    chat.innerHTML += `<p>${Event.data}</p>`
    // Adiciona o conteúdo da mensagem (Event.data) como um novo parágrafo (<p>)
    // ao final da área de chat.
};

function enviar() {
    // 2. Função de Envio: Chamada ao clicar no botão "Enviar" (ou pressionar Enter).
    
    const input = document.getElementById("msg");
    // Seleciona o campo de texto onde o usuário digitou a mensagem.
    
    socket.send(input.value);
    // Envia o conteúdo do campo (input.value) para o servidor através da conexão WebSocket.
    
    input.value = "";
    // Limpa o campo de entrada para que o usuário possa digitar a próxima mensagem.
}