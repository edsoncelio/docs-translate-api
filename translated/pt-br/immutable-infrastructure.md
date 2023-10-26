---
título: Infraestrutura imutável
status: Concluído
categoria: propriedade
tags: ["infraestrutura”, “propriedade”, “"]
---

Infraestrutura imutável se refere à infraestrutura de computadores 
([máquinas virtuais] (/virtual-machine/), [contêineres] (/container/), dispositivos de rede) 
que não pode ser alterado depois de implantado. 
Isso pode ser aplicado por um processo automatizado que substitui alterações não autorizadas ou 
por meio de um sistema que não permite mudanças em primeiro lugar. 
Os contêineres são um bom exemplo de infraestrutura imutável 
porque alterações persistentes nos contêineres só podem ser feitas por 
criando uma nova versão do contêiner ou recriando o contêiner existente a partir de sua imagem.

Ao impedir ou identificar alterações não autorizadas, 
infraestruturas imutáveis facilitam a identificação e a mitigação dos riscos de segurança. 
Operar esse sistema se torna muito mais simples 
porque os administradores podem fazer suposições sobre isso. 
Afinal, eles sabem que ninguém cometeu erros ou mudanças que eles esqueceram de comunicar. 
A infraestrutura imutável anda de mãos dadas com [infraestrutura como código] (/infrastructure-as-code/) 
onde toda a automação necessária para criar a infraestrutura é armazenada no controle de versão (por exemplo, Git). 
Essa combinação de imutabilidade e controle de versão significa que 
há um registro de auditoria durável de cada alteração autorizada em um sistema.
