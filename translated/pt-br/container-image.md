---
título: Imagem do contêiner
status: Feedback apreciado
categoria: conceito
etiquetas: ["”, “”, “"]
---

## O que é

Uma imagem de contêiner é um arquivo estático imutável que contém as dependências para a criação de um [contêiner] (/container/). 
Essas dependências podem incluir um único arquivo binário executável, bibliotecas do sistema, 
ferramentas do sistema, variáveis de ambiente e outras configurações de plataforma necessárias. 
As imagens de contêiner resultam da [conteinerização] de um aplicativo (/containerization/) e normalmente são armazenadas em registros de contêineres, 
onde eles podem ser baixados e executados como um processo isolado usando uma Container Runtime Interface (CRI). 
Uma estrutura de imagem de contêiner deve seguir o esquema padrão definido pela Open Container Initiative (OCI).

## Problema que ele resolve 

Tradicionalmente, os servidores de aplicativos são configurados por ambiente e os aplicativos são implantados neles. 
Qualquer configuração incorreta entre ambientes é problemática e geralmente causa inatividade ou falhas nas implantações. 
O ambiente de um aplicativo precisa ser repetível e bem definido; 
caso contrário, a chance de bugs relacionados ao meio ambiente aumenta. 
Quando os ambientes de aplicativos são configurados de forma inadequada ou imprecisa, 
O dimensionamento [horizontal] (/horizontal-scaling/) e [vertical] (/vertical-scaling/) de aplicativos se torna um desafio. 

## Como isso ajuda

As imagens de contêiner agrupam um aplicativo com qualquer uma de suas dependências de tempo de execução, como um servidor de aplicativos. 
Isso fornece consistência em todos os ambientes, incluindo a máquina do desenvolvedor. 
As imagens de contêiner podem ser usadas para instanciar quantos contêineres forem necessários, permitindo maior [escalabilidade] (/escalabilidade/). 
