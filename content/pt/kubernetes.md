---
título: Kubernetes
Status: concluído
Categoria: Tecnologia
tags: ["infraestrutura”, “infraestrutura”, “"]
---

## O que é

O Kubernetes, geralmente abreviado como K8s, é um orquestrador de contêineres de código aberto. 
Ele automatiza o ciclo de vida de aplicativos em contêineres em infraestruturas modernas, desenvolvido como um “sistema operacional de data center” que distribuía aplicativos em um [sistema distribuído] (/distributed-systems/).

O Kubernetes agenda [contêineres] (/container/) em [nós] (/nodes/) em um [cluster] (/cluster/), agrupando recursos de infraestrutura, como balanceadores de carga, armazenamento persistente, convenientes para executar aplicativos em contêineres.

Automação e extensões do Kubernetes, usuários descentralizados para implantar declarações de aplicativos (veja abaixo) de forma reproduzível. 
O Kubernetes é extensível por meio de sua [API] (/application-programming-interface/). O extenso Kubernetes se expande para seus recursos de automação de acordo com suas necessidades.

## Problema que ele resolve

A automação da infraestrutura e o gerenciamento declarativo de configurações são conceitos importantes há muito tempo, mas se tornaram conceitos mais importantes à medida que a [computação em nuvem] (/cloud-computing/) já se comprovou. 
Como a demanda por recursos de computação, profissionais e organizações precisam fornecer mais recursos operacionais com engenheiros competentes, novas tecnologias e métodos de trabalho são necessários para atender a essa demanda. 

## Como isso aconteceu

Semelhante às ferramentas tradicionais de [infraestrutura como código] (/infrastructure-as-code/), o Kubernetes implementa com automação, mas tem a vantagem de trabalhar com contêineres. 
Os contêineres são mais fáceis de configurar do que as máquinas [virtuais] (/virtual-machine/) ou físicas. 

Declarações, declarações do Kubernetes Works, o que significa que, em vez de os operadores instruírem a máquina sobre como fazer algo, eles descrevem — explicam como arquivos de manifesto (por exemplo, YAML) — como a infraestrutura deve ser. 
O Kubernetes então cuida do “como”. 
Isso faz com que o Kubernetes seja compatível com a infraestrutura como código.

O Kubernetes também [cura por si mesmo] (/autocura/). 
O ESTADO REAL DO CLUSTER SEMPRE CORRESPONDERÁ AO ESTADO DESEJADO PELO OPERADOR.
Se o Kubernetes quebrar um desvio do que está acontecendo nos arquivos de manifesto, um controlador do Kubernetes entra em ação e o corrige. 
Embora a infraestrutura usada pelo Kubernetes possa estar mudando, o Kubernetes se adapta e se adapta às mudanças e mudanças que correspondam ao estado desejado.