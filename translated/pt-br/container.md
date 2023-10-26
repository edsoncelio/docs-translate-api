---
título: Contêineres
status: Concluído
categoria: tecnologia
tags: ["aplicativo”, “fundamental”, “"]
---

## O que é

Um contêiner é um processo em execução com restrições de recursos e capacidades gerenciadas pelo sistema operacional de um computador. 
Os arquivos disponíveis para o processo de contêiner são empacotados como uma imagem de contêiner. 
Os contêineres correm adjacentes uns aos outros na mesma máquina, 
mas normalmente o sistema operacional impede que os processos separados do contêiner interajam entre si.

## Problema que ele aborda

Antes que os contêineres estivessem disponíveis, máquinas separadas eram necessárias para executar os aplicativos. 
Cada máquina exigiria seu próprio sistema operacional, que ocupa CPU, memória e espaço em disco, 
tudo para que um aplicativo individual funcione. 
Além disso, a manutenção, o upgrade e a inicialização de um sistema operacional são outra fonte significativa de trabalho. 

## Como isso ajuda

Os contêineres compartilham o mesmo sistema operacional e seus recursos de máquina, 
espalhando a sobrecarga de recursos do sistema operacional e criando um uso eficiente da máquina física. 
Esse recurso só é possível porque os contêineres normalmente não conseguem interagir uns com os outros. 
Isso permite que muitos outros aplicativos sejam executados na mesma máquina física.

No entanto, existem limitações. 
Como os contêineres compartilham o mesmo sistema operacional, os processos podem ser considerados menos seguros do que as alternativas. 
Os contêineres também exigem limites nos recursos compartilhados. 
Para garantir recursos, os administradores devem restringir e limitar o uso da memória e da CPU para que outros aplicativos não tenham um desempenho ruim.
