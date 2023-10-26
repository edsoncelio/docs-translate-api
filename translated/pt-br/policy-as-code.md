---
título: Política como Código (PaC)
status: Concluído
categoria: conceito
tags: ["metodologia”, “”, “"]
rascunho: 
---

## O que é

Política como código é a prática de armazenar a definição de políticas como um ou mais arquivos em formato legível e processável por máquina. 
Isso substitui o modelo tradicional em que as políticas são documentadas de forma legível por humanos em documentos separados.

## Problema que ele resolve

A construção de aplicativos e infraestruturas geralmente é limitada por muitas políticas definidas por uma organização, 
por exemplo, políticas de segurança que proíbem armazenar segredos no código-fonte, executar um contêiner com permissões de superusuário, 
ou armazenar alguns dados fora de uma região geográfica específica.
É muito trabalhoso e propenso a erros que desenvolvedores e revisores verifiquem manualmente os aplicativos e a infraestrutura em relação às políticas documentadas. 
Os processos manuais não podem atender aos requisitos de capacidade de resposta e escala dos aplicativos nativos da nuvem.

## Como isso ajuda

Descrever políticas por meio de código permite a repetibilidade e reduz os erros (ao contrário de quando feito manualmente). 
Outra vantagem do Policy as Code é que o código pode ser gerenciado por um sistema de controle de versão como o Git.
O Git cria um histórico de registro de alterações que é particularmente útil quando algo não funciona conforme o esperado.
Ele permite que o usuário determine quem fez a alteração e volte para uma versão anterior.  
