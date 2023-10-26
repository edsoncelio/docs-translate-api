---
título: Pod
status: Concluído
categoria: conceito
tags: ["infraestrutura”, “fundamental”, “"]
---

## O que é

Em um ambiente [Kubernetes] (/kubernetes/), um pod atua como a unidade implantável mais básica.
Ele representa um componente essencial para a implantação e o gerenciamento de aplicativos em contêineres.
Cada pod contém uma única instância de aplicativo e pode conter um ou mais [contêineres] (/container/).
O Kubernetes gerencia pods como parte de uma implantação maior e pode escalar pods [verticalmente] (/vertical-scaling/) ou [horizontalmente] (/horizontal-scaling/) conforme necessário.

## Problema que ele resolve

Embora os contêineres geralmente atuem como unidades independentes que executam e controlam uma carga de trabalho específica, 
há casos em que os contêineres precisam interagir e ser controlados de forma fortemente acoplada. 

Se cada um desses contêineres estreitamente relacionados fosse gerenciado individualmente, isso resultaria em tarefas de gerenciamento redundantes.
Por exemplo, o operador precisaria determinar repetidamente a localização dos contêineres relacionados para garantir que eles permaneçam juntos.
E embora os ciclos de vida desses contêineres relacionados precisem ser sincronizados, eles só podem ser gerenciados individualmente. 


## Como isso ajuda

Os pods agrupam contêineres estreitamente conectados em uma única unidade, simplificando significativamente as operações de contêineres.
Por exemplo, contêineres auxiliares geralmente são usados junto com o contêiner principal para adicionar funcionalidades adicionais ou configurar configurações globais. 
Os exemplos incluem contêineres que injetam e aplicam configurações básicas ao contêiner principal, 
_sidecar_ (contêineres) que lidam com o roteamento de tráfego de rede para o contêiner principal (consulte [service mesh] (/service-mesh/)), 
ou contêineres coletando troncos em conjunto com cada contêiner.

A alocação de memória e CPU pode ser definida no nível do pod, permitindo que os contêineres internos compartilhem recursos de forma flexível ou por contêiner.
