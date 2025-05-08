import './assets/main.css' // CSS principal

import { createApp } from 'vue' // Cria a aplicação
import App from './App.vue'// Componente principal
import router from './router'

const app = createApp(App) // Criando a aplicação utilizando esse App como meu componente principal

app.use(router) // Vai utilizar o roteamento

app.mount('#app') // Montando a aplicação dentro de uma div (no index.html)
