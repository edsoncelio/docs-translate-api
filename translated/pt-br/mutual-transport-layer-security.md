---
título: Mutual Transport Layer Security (mTLS)
status: Concluído
categoria: Conceito
tags: ["segurança”, “rede”, “"]
---

## O que é

O TLS mútuo (mTLS) é uma técnica usada para autenticar e criptografar mensagens enviadas entre dois [serviços] (/service/). 
O TLS mútuo é o protocolo padrão [Transport Layer Security] (/transport-layer-security/) (TLS), mas, 
em vez de validar a identidade de apenas uma conexão, os dois lados são validados.

## Problema que ele resolve

[Microserviços] (/microserviços/) se comunicam por meio de uma rede e, 
assim como sua rede wifi, a comunicação em trânsito por essa rede pode ser invadida. 
O mTLS garante que nenhuma parte não autorizada possa ouvir ou se passar por solicitações legítimas.

## Como isso ajuda

O mTLS garante que o tráfego seja seguro e confiável nas duas direções entre um cliente e um servidor, 
fornecendo uma camada adicional de segurança para usuários que fazem login em uma rede ou em aplicativos. 
Ele também verifica as conexões com dispositivos clientes que não seguem um processo de login, como dispositivos da Internet das Coisas (IoT). 
Ataques como ataques diretos, ataques de falsificação, preenchimento de credenciais, ataques de força bruta etc. podem ser evitados pelo mTLS.
