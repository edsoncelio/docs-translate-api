---
Título: Serverless
Status: Concluído
Categoria: Tecnologia
tags: ["arquitetura”, “”, “"]
---

## O que é

Serverless é um modelo de desenvolvimento nativo em nuvem que permite aos desenvolvedores 
crie e execute aplicativos sem precisar gerenciar servidores. 
Embora os servidores ainda existam dentro do paradigma sem servidor, eles estão [abstraídos] (/abstraction/) longe do processo de desenvolvimento de aplicativos.
Um provedor de nuvem lida com o trabalho rotineiro de provisionar, manter e [escalar] (/escalabilidade/) da infraestrutura do servidor. 
Os desenvolvedores podem facilmente empacotar seu código em [containers] (/container/) para implantação.
Depois de implantados, os aplicativos sem servidor respondem à demanda e aumentam e diminuem automaticamente a escala conforme necessário. 
As ofertas sem servidor de provedores de nuvem pública geralmente são medidas sob demanda por meio de um modelo de execução orientado por eventos. 
Consequentemente, quando uma função sem servidor está em um estado inativo, não há custos associados.

## Problema que ele resolve

Sob um modelo padrão [Infraestrutura como serviço (IaaS)] (/infraestrutura-como-serviço/) [computação em nuvem] (/computação-nuvem/), 
os usuários pré-compram unidades de capacidade, o que significa que você paga a um provedor de nuvem pública por componentes de servidor sempre ativos para executar seus aplicativos. 
É responsabilidade do usuário aumentar a capacidade do servidor em períodos de alta demanda e 
para reduzir quando essa capacidade não for mais necessária. 
A infraestrutura de nuvem necessária para operar um aplicativo permanece ativa mesmo quando o aplicativo não está em uso.

## Como isso ajuda

Em contraste com as abordagens tradicionais, a arquitetura sem servidor lança aplicativos somente quando eles são necessários. 
Quando um evento aciona a execução do código do aplicativo, o provedor de nuvem pública aloca dinamicamente recursos para esse código. 
O usuário para de pagar quando o código termina de ser executado. 
Além dos benefícios de custo e eficiência, 
O serverless libera os desenvolvedores de tarefas rotineiras e domésticas associadas ao dimensionamento de aplicativos e ao provisionamento de servidores. 
Com tarefas rotineiras e sem servidor, como gerenciar o sistema operacional e o sistema de arquivos, patches de segurança, 
balanceamento de carga, gerenciamento de capacidade, escalabilidade, registro e monitoramento são todos transferidos para um provedor de serviços em nuvem.
