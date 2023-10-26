---
título: Máquina virtual
status: Concluído
categoria: Tecnologia
tags: ["fundamental”, “infraestrutura”, “"]
---

## O que é

Uma máquina virtual (VM) é um computador e seu sistema operacional 
que não está vinculado a uma peça específica de hardware. 
As VMs dependem da [virtualização] (/virtualization/) para dividir um único computador físico em vários computadores virtuais. 
Essa separação permite que organizações e provedores de infraestrutura 
crie e destrua facilmente VMs sem afetar o hardware subjacente.

## Problema que ele resolve

Quando uma máquina [bare metal] (/bare-metal-machine/) está vinculada a um único sistema operacional, 
o quão bem os recursos da máquina podem ser usados é um pouco limitado. 
Além disso, quando um sistema operacional está vinculado a uma única máquina física, 
sua disponibilidade está diretamente vinculada a esse hardware. 
Se a máquina física estiver offline devido a falhas de manutenção ou hardware, o sistema operacional também estará.

## Como isso ajuda

Ao remover a relação direta entre um sistema operacional e uma única máquina física, 
você resolve vários problemas de máquinas de metal puro: 
tempo de provisionamento, utilização de hardware e resiliência.

Sem nenhum hardware novo a ser comprado, instalado ou configurado para suportá-lo, 
o tempo de provisionamento de um novo computador foi drasticamente aprimorado. 
As VMs permitem que você use melhor seus recursos de hardware físico existentes 
colocando várias máquinas virtuais em uma única máquina física. 
Não vinculadas a uma máquina física específica, as VMs também são mais resilientes do que as máquinas físicas. 
Quando uma máquina física precisa ficar off-line, 
as VMs em execução nela podem ser movidas para outra máquina com pouco ou nenhum tempo de inatividade
