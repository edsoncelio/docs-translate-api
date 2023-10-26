---
título: Service Discovery
status: Concluído
categoria: conceito
etiquetas: ["networking”, “”, “"]
---

## O que é

A descoberta de serviços é o processo de encontrar instâncias individuais que compõem um serviço. 
Uma ferramenta de descoberta de serviços acompanha os vários nós ou endpoints que compõem um serviço. 

## Problema que ele aborda

As arquiteturas nativas da nuvem são dinâmicas e fluidas, o que significa que estão em constante mudança. 
Um aplicativo [conteinerizado] (/containerization/) provavelmente acabará iniciando e parando várias vezes durante sua vida útil. 
Cada vez que isso acontecer, ele terá um novo endereço e 
qualquer aplicativo que queira encontrá-lo precisa de uma ferramenta para fornecer as novas informações de localização. 

## Como isso ajuda

A descoberta de serviços acompanha os aplicativos na rede para que eles possam se encontrar quando necessário. 
Ele fornece um local comum para encontrar e potencialmente identificar serviços individuais. 
Os mecanismos de descoberta de serviços são ferramentas semelhantes a bancos de dados que armazenam informações sobre quais serviços existem e como localizá-los.
