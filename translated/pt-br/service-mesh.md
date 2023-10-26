---
título: Service Mesh
status: Concluído
categoria: tecnologia
etiquetas: ["networking”, “”, “"]
---

## O que é

Em um mundo de [microsserviços] (/microserviços/), os aplicativos são divididos em vários [serviços] menores (/service/) que se comunicam por meio de uma rede. 
Assim como sua rede wifi, as redes de computadores são intrinsecamente não confiáveis, hackeáveis e geralmente lentas. 
As malhas de serviços abordam esse novo conjunto de desafios gerenciando o tráfego (ou seja, a comunicação) entre serviços e 
adicionando [confiabilidade] (/confiability/), [observabilidade] (/observabilidade/) e recursos de segurança de forma uniforme em todos os serviços.

## Problema que ele resolve

Depois de migrar para uma arquitetura de microsserviços, os engenheiros agora estão lidando com centenas, 
possivelmente até milhares de serviços individuais, todos precisando se comunicar. 
Isso significa que muito tráfego está indo e voltando pela rede. 
Além disso, aplicativos individuais podem precisar criptografar as comunicações para atender aos requisitos regulatórios, 
forneça métricas comuns às equipes de operações ou forneça informações detalhadas sobre o tráfego para ajudar a diagnosticar problemas. 
Se incorporado aos aplicativos individuais, 
cada um desses recursos causará atrito entre as equipes e retardará o desenvolvimento de novos recursos.

## Como isso ajuda

As malhas de serviço adicionam recursos de confiabilidade, observabilidade e segurança 
uniformemente em todos os serviços em um cluster sem exigir alterações no código. 
Antes das malhas de serviços, essa funcionalidade precisava ser codificada em cada serviço, 
tornando-se uma fonte potencial de bugs e dívidas técnicas.
