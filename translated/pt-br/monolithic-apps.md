---
título: Monolithic Apps
status: Concluído
categoria: conceito
tags: ["arquitetura”, “fundamental”, “"]
---

## O que é

Um aplicativo monolítico contém todas as funcionalidades em um único programa implantável. 
Geralmente, esse é o lugar mais simples e fácil para começar a criar um aplicativo. 
No entanto, quando a complexidade do aplicativo aumenta, os monólitos podem se tornar difíceis de manter. 
Com mais desenvolvedores trabalhando na mesma base de código, 
a probabilidade de mudanças conflitantes e a necessidade de comunicação interpessoal entre desenvolvedores aumentam.

## Problema que ele resolve

Transformar um aplicativo em [microservices] (/microservices/) aumenta sua sobrecarga operacional 
— há mais coisas para testar, implantar e continuar funcionando. 
No início do ciclo de vida de um produto, pode ser vantajoso adiar essa complexidade e criar um aplicativo monolítico 
até que o produto seja determinado como bem-sucedido.

## Como isso ajuda

Um monólito bem projetado pode defender os princípios do lean sendo a maneira mais simples de colocar um aplicativo em funcionamento. 
Quando o valor comercial do aplicativo monolítico é bem-sucedido, ele pode ser decomposto em microsserviços. 
Criar um aplicativo baseado em microsserviços antes que ele se mostre valioso pode ser um esforço prematuro de engenharia. 
Se o aplicativo não gerar nenhum valor, esse esforço será desperdiçado.
