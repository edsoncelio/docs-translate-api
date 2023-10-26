---
título: Kubernetes
status: Concluído
categoria: tecnologia
tags: ["infraestrutura”, “fundamental”, “"]
---

## O que é

O Kubernetes, geralmente abreviado como K8s, é um orquestrador de contêineres de código aberto. 
Ele automatiza o ciclo de vida de aplicativos em contêineres em infraestruturas modernas, funcionando como um “sistema operacional de data center” que gerencia aplicativos em um [sistema distribuído] (/distributed-systems/).

O Kubernetes agenda [contêineres] (/container/) em [nós] (/nodes/) em um [cluster] (/cluster/), agrupando vários recursos de infraestrutura, como balanceador de carga, armazenamento persistente etc. para executar aplicativos em contêineres.

O Kubernetes permite automação e extensibilidade, permitindo que os usuários implantem aplicativos de forma declarativa (veja abaixo) de forma reproduzível. 
O Kubernetes é extensível por meio de sua [API] (/application-programming-interface/), permitindo que profissionais experientes do Kubernetes aproveitem seus recursos de automação de acordo com suas necessidades.

## Problema que ele aborda

A automação da infraestrutura e o gerenciamento declarativo de configurações são conceitos importantes há muito tempo, mas se tornaram mais urgentes à medida que a [computação em nuvem] (/cloud-computing/) ganhou popularidade. 
À medida que a demanda por recursos computacionais aumenta e as organizações precisam fornecer mais recursos operacionais com menos engenheiros, novas tecnologias e métodos de trabalho são necessários para atender a essa demanda. 

## Como isso ajuda

Semelhante às ferramentas tradicionais de [infraestrutura como código] (/infrastructure-as-code/), o Kubernetes ajuda na automação, mas tem a vantagem de trabalhar com contêineres. 
Os contêineres são mais resistentes a variações de configuração do que as máquinas [virtuais] (/virtual-machine/) ou físicas. 

Além disso, o Kubernetes funciona de forma declarativa, o que significa que, em vez de os operadores instruírem a máquina sobre como fazer algo, eles descrevem — geralmente como arquivos de manifesto (por exemplo, YAML) — como a infraestrutura deve ser. 
O Kubernetes então cuida do “como”. 
Isso faz com que o Kubernetes seja extremamente compatível com a infraestrutura como código.

O Kubernetes também [se cura] (/autocura/). 
O estado real do cluster sempre corresponderá ao estado desejado pelo operador.
Se o Kubernetes detectar um desvio do que está descrito nos arquivos de manifesto, um controlador do Kubernetes entra em ação e o corrige. 
Embora a infraestrutura usada pelo Kubernetes possa estar mudando continuamente, o Kubernetes se adapta constante e automaticamente às mudanças e garante que ela corresponda ao estado desejado.
