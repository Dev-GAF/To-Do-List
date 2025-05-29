<script setup>

import IconWrapper from "../components/icons/IconWrapper.vue";
import Modal from "../components/Modal.vue";

import { ref } from 'vue';
import { useRoute } from 'vue-router';
// import { useTarefaStore } from '@/store/tarefaStore';

import api from '@/services/api'

const title = ref('');
const tasks = ref([]);
const newTask = ref('');
const isModalOpen = ref(false);

// const tarefaStore = useTarefaStore();
const route = useRoute();

if (route.query.title) {
  title.value = String(route.query.title);
}

function controlModal() {
    isModalOpen.value = !isModalOpen.value;
}

async function addTask() {
    if (!title.value.trim()) {
        alert('Digite um título antes de adicionar uma tarefa.');
        return;
    }

    if (newTask.value.trim() !== '') {
        // const task = { text: newTask.value, done: false };
        // tasks.value.push(task);

        // // Também salva na Pinia
        // const conteudo = `- [ ] ${task.text}`;
        // tarefaStore.adicionarTarefa(title.value, conteudo, 'checklist');

        // newTask.value = '';
        // isModalOpen.value = false;

        const task = {
            id: Date.now(),
            titulo: newTask.value,
            feito: false
        }

        try 
        {
            const { data: checklists } = await api.get('/checklists');
            const existente = checklists.find(c => c.titulo == title.value);

            if (existente) 
            {
                existente.items.push(task);
                await api.put(`/checklists/${existente.id}`, existente)
                tasks.value = [...existente.items];
            } 
            else 
            {
                const novoChecklist = {
                    titulo: title.value,
                    items: [task]
                };
                const { data } = await api.post('/checklists', novoChecklist);
                tasks.value = [...data.items];
            }

            newTask.value = '';
            isModalOpen.value = false;
        } catch (error) {
            console.error('Erro ao salvar checklist:', error);
        }
    }
}

</script>


<template>
    <div class="container">
        <input v-model="title" placeholder="Título" class="note-title" />
        <div id="teste">
            <ul class="task-list">
                <li v-for="(task, index) in tasks" :key="index">
                    <label class="custom-checkbox">
                        <input type="checkbox" v-model="task.done" />
                        <span class="checkmark"></span>
                        <span :class="{ done: task.done }">{{ task.text }}</span>
                    </label>
                </li>
            </ul>
        </div>
        <div class="footer">
            <icon-wrapper 
                :tamanho="55"
                corFundo="none"
                corIcone="White"
                @click="controlModal"
            />

            <Modal :isOpen="isModalOpen" @close="isModalOpen = false">
                <div class="modal-content">
                    <div>
                        <h3>Digite sua Tarefa:</h3>
                    </div>
                    <div>
                        <input type="text" name="task" id="task" placeholder="Digite sua tarefa aqui..." v-model="newTask">
                    </div>
                    <div class="div-button">
                        <button @click="addTask" class="button-addTask">Adicionar</button>
                    </div>
                </div>
            </Modal>
        </div>
    </div>
</template>


<style scoped>

input {
    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}

textarea {
    font-family: Arial, Helvetica, sans-serif;
}

.note-title::placeholder,
.note-content::placeholder {
    font-family: sans-serif;
    color: rgba(255, 255, 255, 0.5);
}

.container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: 100vw;
}

.note-title {
    flex: 0 0 auto;
    padding: 10px;
    font-size: 30px;
    background: none;
    border: none;
    outline: none;
    color: white;
}

.note-content {
    flex: 1;
    padding: 10px;
    font-size: 18px;
    background: transparent;
    border: none;
    outline: none;
    resize: none;
    color: white;
}

.footer {
    flex: 0 0 auto;
    background: rgba(0, 0, 0, 0.541);
    display: flex;
    justify-content: center;
}

.footer ::v-deep(.icon-wrapper) {
    border-radius: 50%;
    box-shadow: none;
    transform: none;
}

#teste {
    flex: 1;
}

.modal-content {
    background-color: rgb(28, 29, 29);
    border-radius: 30px;
    padding: 50px;
}

.modal-content h3 {
    font-size: 80px;
    font-family: Arial, Helvetica, sans-serif;
}

#task {
    width: 100%;
    height: 3vh;
    background-color: transparent;
    border: none;
    border-bottom: 2px solid rgb(104, 101, 101);
    outline: none;
    resize: none;
    color: white;
    font-family: Arial, Helvetica, sans-serif;
    font-size: 15px;
    margin-top: 30px;
}

#task::placeholder {
    color: rgba(255, 255, 255, 0.727);
    font-family: sans-serif;
    padding-left: 2px;
}

.div-button {
    display: flex;
    justify-content: center;
    align-items: center;
}

.button-addTask {
    margin-top: 30px;
    padding: 12px 24px;
    font-size: 16px;
    color: white;
    background: linear-gradient(135deg, #00d4ff, #0077ff);
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: background 0.3s, transform 0.2s;
    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}

.button-addTask:hover {
    background: linear-gradient(135deg, #00aaff, #0055cc);
    transform: scale(1.05);
}

.task-list {
    list-style-type: none;
    padding: 10px;
    color: white;
    font-size: 18px;
}

.task-list li {
    margin-bottom: 10px;
    display: flex;
    align-items: center;
}

.task-list input[type="checkbox"] {
    margin-right: 10px;
    accent-color: #00d4ff;
}

.done {
    text-decoration: line-through;
    opacity: 0.6;
}



</style>