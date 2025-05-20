import { defineStore } from "pinia";

export const useTarefaStore = defineStore('tarefa', {
    state: () => ({
        tarefas: [], // [{ id, titulo, conteudo, tipo, done }]
        checklistTituloTemp: ''
    }),
    actions: {
        adicionarTarefa(titulo, conteudo = '', tipo = 'nota') {
            this.tarefas.push({
                id: Date.now(),
                titulo,
                conteudo,
                tipo,
                done: false
            });
        },
        alternarFeito(id) {
            const tarefa = this.tarefas.find(t => t.id === id);
            if (tarefa) tarefa.done = !tarefa.done;
        },
        removerTarefa(id) {
            this.tarefas = this.tarefas.filter(t => t.id !== id);
        }
    }
    //persist: true
})