---
título: Multilocação
status: Concluído
categoria: Propriedade
tags: ["arquitetura”, “propriedade”, “"]
---

## O que é

A multilocação (ou multilocação) se refere a uma única instalação de software que atende a vários inquilinos.
Um inquilino é um usuário, aplicativo ou grupo de usuários/aplicativos que utilizam o software para operar em seu próprio conjunto de dados.
Esses inquilinos não compartilham dados (a menos que sejam explicitamente instruídos pelo proprietário) e podem nem mesmo estar cientes uns dos outros.

Um inquilino pode ser tão pequeno quanto um usuário independente com um único ID de login — pense em produtividade pessoal
software — ou tão grande quanto uma corporação inteira com milhares de IDs de login, cada uma com seus próprios privilégios
mas inter-relacionados de várias maneiras. Exemplos de software multilocatário incluem Google Mail, Google Docs,
Microsoft Office 365, Salesforce CRM e Dropbox, entre muitos outros que são categorizados como totalmente
ou software parcialmente multilocatário.

## Problema que ele resolve

Sem multilocação, cada inquilino precisaria de uma instalação de software dedicada.
Isso aumenta a utilização de recursos e os esforços de manutenção e, em última análise, os custos de software.

## Como isso ajuda

O software multilocatário fornece a cada inquilino um ambiente segregado (dados de trabalho, configurações, lista de credenciais etc.),
atendendo simultaneamente vários inquilinos. Do ponto de vista do inquilino, cada um tem sua instalação de software dedicada,
embora, na realidade, todos estejam compartilhando um. Isso é obtido executando o software em um servidor e permitindo
inquilinos devem se conectar a ele pela rede por meio de uma interface e/ou uma [API] (/application-programming-interface/)
(consulte também [Arquitetura cliente-servidor] (/client-server-architecture/)).
Com o software multilocatário, os inquilinos compartilham os recursos de uma instalação sem afetar uns aos outros ou apenas
de formas predefinidas e controladas. A economia de recursos resultante do lado do fornecedor de software pode ser repassada
para os inquilinos, reduzindo significativamente o custo do software para os usuários (novamente, pense em editores de e-mail ou documentos baseados na web).

## Termos relacionados

Multilocação não é sinônimo de SaaS,
embora seja muito comum que o SaaS seja multilocatário e até mesmo tenha a multilocação como um de seus principais benefícios.
