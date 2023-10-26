---
título: Service Proxy
status: Concluído
categoria: tecnologia
etiquetas: ["networking”, “”, “"]
---

## O que é

Um proxy de serviço intercepta o tráfego de ou para um determinado [serviço] (/service/), 
aplica alguma lógica a ele e, em seguida, encaminha esse tráfego para outro serviço. 
Ele atua essencialmente como um “intermediário” que coleta informações sobre o tráfego de rede e/ou aplica regras a ele.

## Problema que ele resolve

Para acompanhar a comunicação entre serviços (também conhecida como tráfego de rede) e 
potencialmente transformá-lo ou redirecioná-lo, precisamos coletar dados. 
Tradicionalmente, o código que permite a coleta de dados e o gerenciamento do tráfego de rede era incorporado em cada aplicativo.

## Como isso ajuda

Um proxy de serviço nos permite “externalizar” essa funcionalidade. 
Ele não precisa mais estar nos aplicativos. 
Em vez disso, agora está incorporado na camada da plataforma (onde seus aplicativos são executados).

Atuando como guardiões entre os serviços, os proxies fornecem informações sobre o tipo de comunicação que está acontecendo. 
Com base em suas percepções, eles determinam para onde enviar uma solicitação específica ou até mesmo negá-la totalmente.

Os proxies coletam dados críticos, gerenciam o roteamento (distribuindo o tráfego uniformemente entre os serviços ou redirecionando se alguns serviços falharem), 
criptografar conexões e armazenar conteúdo em cache (reduzindo o consumo de recursos).
