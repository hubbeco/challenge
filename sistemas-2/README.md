# 🖥️ Desenvolvedor de Sistemas 2

> 🎯 **Candidato a contratação deve criar uma game engine em C**, implementando desde a API em **Lua** até a camada de renderização gráfica **OpenGL 2**.

## 🎮 Requisitos da Engine

A engine deve possuir **2 funções de desenho** e **2 callbacks**.

### 📐 Funções de desenho

| Função               | Descrição                             |
| :------------------- | :------------------------------------ |
| `rect(x, y, w, h)`   | Desenha um retângulo sólido branco 🟦 |
| `png(x, y, src)`     | Desenha uma imagem 🖼️                 |

### 🔁 Callbacks

| Callback                                 | Descrição                                      |
| :--------------------------------------- | :--------------------------------------------- |
| `function tick() end`                    | Chamada cerca de **60 vezes por segundo** ⏱️    |
| `function key(name, pressed) end`        | Chamada quando uma tecla é pressionada/solta ⌨️ |

---

## 🛠️ Como executar o projeto

Já existe um [`CMakeLists.txt`](CMakeLists.txt) que vai baixar automaticamente as dependências:  
_(`glad`, `glfw`, `spng`, `libz`, `lua`, `klib`)_

Você deve copiá-lo, e criar seu código fonte na pasta `src/` para seu projeto e executar o comando de preparação:

```
cmake -Bbuild -H.
```

Após isso, você deve compilar o binário.

```
make -C build
```

E executar um dos jogos de exemplo

```
./build/bin/engine pong.lua
```

### Regras Técnicas

 * Deve rodar em 640x480.
 * Utilizar OpenGL 2 moderno.
 * O Código deve estar otimizado e limpo.
 * A Engine deve suportar rodar no linux.
 * Todos os jogos de exemplo devem funcionar.
 * Gerenciamento de memória seguro, sem vazamentos.

### Diferenciais :sparkles:

 * Adicionar testes unitários.
 * Adicionar flags extras como `--fps` e outras.
 * Adicionar suporte a gamepad (controle de videogame).
 * Utilização de doxygen para documentação.

---

_encaminhe o link de seu repositório no github para o rh._
_o tempo esperado para o desafio é de uma semana!_

:raising_hand_man: Boa sorte!
