# Hello World

[Estudo sobre GO utilizando TDD](https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/ola-mundo)


## Anotações:

* Funções podem ser declaradas dentro de outra função desde que a mesma seja atribuida a uma variavel.
* a Função `Helper()` de `testing` marca a função que a recebe como auxiliar, em caso de falha ela imprime a função que a chamou. O não uso dela na stack do debug mostrará a linha onde falhou o teste
* Funções privadas comecam com letra minúscula enquanto as publicas comecam com maiúscula
* Quando o valores das propriedades da função forem do mesmo tipo pode-se fazer a tipagem no ultimo valor que será aplicado a todos. 
* O retorna da função pode ser nomeada ou nao nomeada. O retorno nomeado é geralmente utilizado quando o contexto nao está claro. Em caso se retorno não nomeado basta declará o tipo do dado retornado