---
título: Transmissão de eventos
status: Concluído
categoria: conceito
tags: ["metodologia”, “networking”, “"]
---

## O que é

O streaming de eventos é uma abordagem em que o software envia dados de eventos de um aplicativo para outro para comunicar continuamente o que eles estão fazendo.
Imagine um serviço transmitindo tudo o que faz para todos os outros serviços.
Cada atividade realizada por um serviço é chamada de evento, portanto, streaming de eventos.
Por exemplo, a NASDAQ recebe atualizações sobre preços de ações e commodities a cada segundo.
Se você tivesse um aplicativo que monitorasse um conjunto específico de estoques, gostaria de receber essas informações quase em tempo real.
Yahoo! O setor financeiro fornece uma [API] (/application-programming-interface/) que extrai da NASDAQ e envia (ou transmite) as informações (ou eventos) de seu aplicativo para qualquer aplicativo que o assine.
Os dados enviados, bem como as alterações nesses dados (preços das ações), são os eventos, enquanto o processo de entregá-los a um aplicativo é o streaming de eventos.

## Problema que ele resolve

Tradicionalmente, o Yahoo! O setor financeiro usaria solicitações TCP únicas.
Isso seria muito ineficiente, pois exigiria a criação de uma conexão para cada evento.
À medida que os dados se tornam mais em tempo real, escalar essa solução se torna ineficiente.
Abrir uma conexão uma vez e permitir que os eventos fluam é ideal para a coleta em tempo real.
A quantidade de dados gerados está crescendo exponencialmente e, com isso, o estado dos dados está em constante fluxo. Desenvolvedores e usuários precisam ser capazes de ver esses dados quase em tempo real.

## Como isso ajuda

O streaming de eventos permite que as alterações de dados sejam comunicadas da fonte ao receptor.
Em vez de esperar que os serviços solicitem informações, o serviço transmite continuamente todos os seus eventos (ou atividades).
Ela não está preocupada com o que acontece com as informações.
Ele apenas faz o que precisa fazer e o transmite, permanecendo completamente independente de qualquer outro serviço.
