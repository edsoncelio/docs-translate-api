---
título: Firewall
status: Obsoleto
rascunho: verdadeiro
categoria: Tecnologia
etiquetas: ["”, “”, “"]
---

## O que é

Um firewall é um sistema que filtra o tráfego de rede com base em regras especificadas. 
Os firewalls podem ser hardware, software ou uma combinação dos dois.

## Problema que ele resolve

Por padrão, uma rede permitirá que qualquer pessoa entre e saia, desde que siga as regras de roteamento da rede. 
Por causa desse comportamento padrão, proteger uma rede é um desafio. 
Por exemplo, em um aplicativo bancário baseado em [microservices] (/microservices/), os serviços se comunicam entre si 
transmitindo dados financeiros altamente confidenciais por meio de sua rede. 
Um agente mal-intencionado pode se infiltrar na rede, interceptar a comunicação e causar danos se não houver um firewall instalado.
 
## Como isso ajuda

Um firewall examina o tráfego de rede usando regras predefinidas. 
Todo o tráfego é filtrado e qualquer tráfego proveniente de fontes suspeitas ou não confiáveis é bloqueado 
— somente o tráfego configurado para ser aceito entra. 
Os firewalls estabelecem uma barreira entre redes confiáveis internas protegidas e controladas. 
