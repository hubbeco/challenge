# Desenvolvedor de Sistemas :computer:

> Candidato a contratação deve criar um banco de dados **chave/valor em C**, implementado um parser de comandos e o armazenamento.

## Requisitos do Banco de Dados :floppy_disk:

O banco de dados deve implementar **apenas 4 comandos**:

| Comando | Descrição |
|---------|-----------|
| `set <chave> <valor>` | Armazena o valor associado à chave. Se a chave existir, substitui o valor. |
| `get <chave>` | Retorna o valor associado à chave. |
| `del <chave>` | Remove a chave e seu valor do banco de dados. |
| `has <chave>` | Retorna se a chave existe ou não. |

### Interface

Expor os 3 tipos de interfaces.

* Protocolo de rede próprio.
* Linha de comando simulada. "REPL"
* Funções C/C++ públicas:
    ```c
    void db_set(const char* key, const char* value);
    void db_del(const char* key);
    bool db_has(const char* key);
    bool db_get(const char* key, char* value_out, size_t max_len);
    ```
### Exemplo de Uso
Via linha de comando

```
> set user Alice
1
> get user
Alice
> has user
1
> del user
1
> has user
0
```

### Regras Técnicas

 * Deve ser implementado em C.
 * Suporte a pelo menos 100 entradas simultâneas.
 * Implementação modular e código limpo.
 * Gerenciamento de memória seguro, sem vazamentos.

### Diferenciais :sparkles:

 * Testes unitários
 * Suporte a crosscompiling
 * Utilização de makefile ou cmake
 * Utilização de doxygen para documentação

---

_encaminhe o link de seu repositório no github para o rh._
_o tempo esperado para o desafio é de uma semana e meia!_

:raising_hand_man: Boa sorte!
