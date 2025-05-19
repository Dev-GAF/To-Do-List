import './assets/main.css' // CSS principal

import { createApp } from 'vue' // Cria a aplicação
import App from './App.vue'// Componente principal
import router from './router'

import { createPinia } from 'pinia'

/*
    📌 Pinia é como um caderno de anotações dentro do meu app.

    -Todos os componentes podem ler e escrever nele.
    -Ele é temporário (a menos que eu use localStorage).
    -Serve para manter tudo sincronizado e organizado.
*/

const pinia = createPinia()
const app = createApp(App) // Criando a aplicação utilizando esse App como meu componente principal

app.use(router) // Vai utilizar o roteamento
app.use(pinia)

app.mount('#app') // Montando a aplicação dentro de uma div (no index.html)
