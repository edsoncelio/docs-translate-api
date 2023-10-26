---
título: eBPF
status: Concluído
categoria: arquitetura
---

## O que é

O eBPF, ou filtro de pacotes de Berkeley estendido, é uma tecnologia que permite que pequenos programas ou scripts em sandbox sejam executados no espaço do kernel de um sistema Linux sem precisar alterar o código-fonte do kernel ou carregar módulos do kernel Linux.

Um sistema Linux tem dois espaços: o kernel e o espaço do usuário. 
O kernel representa o núcleo do sistema operacional e é a única parte 
com acesso ilimitado ao hardware. 

Os aplicativos residem no espaço do usuário e, quando precisam de permissões maiores, 
eles enviam uma solicitação para o kernel.
Para aplicativos que exigem mais flexibilidade, como hardware direto 
acesso, o kernel pode ser estendido por meio do que é conhecido como “Linux” 
abordagem dos módulos do kernel. Essa abordagem estende a funcionalidade padrão do kernel,
 permitindo aos aplicativos um acesso mais profundo aos componentes subjacentes. 
 No entanto, essa abordagem também apresenta riscos de segurança, tornando o eBPF uma alternativa atraente.

## Problema que ele resolve
Normalmente, os aplicativos são executados no espaço do usuário e, se o aplicativo exigir alguns privilégios do kernel (por exemplo, para acessar algum hardware), 
ele o solicita do kernel por meio da chamada “chamada do sistema”. 
Na maioria dos casos, essa abordagem funciona muito bem. No entanto, há casos em que os desenvolvedores precisam de mais flexibilidade para acesso de baixo nível ao sistema.
Os recursos de observabilidade, segurança e rede são bons exemplos.
Para conseguir isso, podemos usar módulos do kernel Linux, estendendo a base do kernel sem modificar o código-fonte do kernel. 
Embora haja benefícios em usar os módulos do kernel Linux, ele também apresenta riscos de segurança. 
Como eles operam dentro do espaço do kernel, os módulos do kernel Linux podem travar o kernel e, quando o kernel falha, o mesmo acontece com toda a máquina.
Além disso, os módulos do kernel têm privilégios elevados e acesso direto aos recursos do sistema. E se não estiverem devidamente protegidos, os atacantes podem explorá-los.

## Como isso ajuda
O eBPF fornece um ambiente mais controlado e contido para executar programas definidos pelo usuário do que os módulos do kernel Linux.
Ele é executado em um ambiente de sandbox dentro do kernel, fornecendo isolamento e mitigando riscos. 
Se uma vulnerabilidade ou falha for explorada em um programa eBPF, seu impacto geralmente é limitado ao ambiente de sandbox.
Além disso, antes que um programa eBPF possa começar a ser executado no kernel, ele precisa passar por algumas verificações. 
O componente verificador verifica o programa eBPF quanto a possíveis violações de segurança, 
como acesso à memória fora dos limites, loops infinitos e funções não autorizadas do kernel.
Dessa forma, ele garante que o programa não entre em um loop infinito e cause uma falha no kernel.
Esses controles de segurança tornam o eBPF uma opção mais segura para executar aplicativos no kernel Linux do que os módulos do kernel Linux.
