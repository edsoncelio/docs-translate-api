---
título: Função como serviço (FaaS)
status: Concluído
categoria: Tecnologia
tags: ["infraestrutura”, “”, “"]
---

## O que é

Função como serviço (FaaS) é um tipo de [serverless] (/serverless/) [computação em nuvem] (/cloud-computing/) [serviço] (/service/) 
que permite executar código em resposta a eventos 
sem manter a infraestrutura complexa 
normalmente associado à criação e lançamento de aplicativos [microsserviços] (/microserviços/). 
Com o FaaS, os usuários gerenciam somente funções e dados, enquanto o provedor de nuvem gerencia o aplicativo. 
Isso permite que os desenvolvedores obtenham as funções de que precisam sem pagar pelos serviços quando o código não está em execução. 

## Problema que ele resolve

Em um cenário local tradicional, uma empresa gerencia e mantém seu próprio data center. 
A empresa deve investir em servidores, armazenamento, software e outras tecnologias 
e potencialmente contrate uma equipe de TI ou prestadores de serviços para comprar, gerenciar e atualizar todos os equipamentos e licenças. 
O data center precisa ser construído para atender aos picos de demanda, mesmo quando as cargas de trabalho diminuem e esses recursos ficam inativos. 
Por outro lado, se a empresa crescer rapidamente, o departamento de TI poderá ter dificuldade em acompanhar o ritmo. 
Sob um modelo padrão de computação em nuvem [Infraestrutura como serviço (IaaS)] (/infraestrutura como serviço/), 
os usuários compram unidades de capacidade com antecedência, o que significa que você paga a um provedor de nuvem pública por componentes de servidor sempre ativos para executar seus aplicativos. 
É responsabilidade do usuário aumentar a capacidade do servidor em períodos de alta demanda 
e reduza a escala quando essa capacidade não for mais necessária. 
A infraestrutura de nuvem necessária para executar um aplicativo está ativa mesmo quando o aplicativo não está sendo usado.

## Como isso ajuda

O FaaS oferece aos desenvolvedores uma [abstração] (/abstraction/) para executar aplicativos web em resposta a eventos sem gerenciar servidores. 
Por exemplo, o upload de um arquivo pode acionar um código personalizado que transcodifica o arquivo em vários formatos. 
A infraestrutura FaaS escalará automaticamente o código para uso pesado, 
e o desenvolvedor não precisa gastar tempo ou recursos criando o código para [escalabilidade] (/escalabilidade/). 
O faturamento é baseado apenas no tempo de computação, o que significa que as empresas não precisam pagar quando as funções não estão em uso.
