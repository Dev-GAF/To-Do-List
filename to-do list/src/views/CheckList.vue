<script setup>
import IconWrapper from "../components/icons/IconWrapper.vue";
import Modal from "../components/Modal.vue";

import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useTarefaStore } from '@/stores/tarefaStore';

const title = ref('');
const tasks = ref([]);
const newTask = ref('');
const isModalOpen = ref(false);

const tarefaStore = useTarefaStore();
const router = useRouter();

function controlModal() 
{
    isModalOpen.value = !isModalOpen.value;
}

function addTask() 
{
    if (newTask.value.trim() !== '') {
        tasks.value.push({ text: newTask.value, done: false });
        newTask.value = '';
        isModalOpen.value = false;
    }
}

function salvarChecklist() 
{
    if (!title.value.trim()) {
        alert('Digite um título antes de salvar.');
        return;
    }

    if (tasks.value.length === 0) {
        alert('Adicione pelo menos uma tarefa ao checklist.');
        return;
    }

    const conteudo = tasks.value.map(t => `- [${t.done ? 'x' : ' '}] ${t.text}`).join('\n');

    tarefaStore.adicionarTarefa(title.value, conteudo, 'checklist');
    router.push('/');
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
                tamanho=55
                corFundo="none"
                corIcone="White"
                @click="controlModal"
            />

            <button @click="salvarChecklist">Salvar Checklist</button>

            <Modal :isOpen="isModalOpen" @close="isModalOpen = false">
                <div class="modal-content">
                    <div>
                        <h3>Digite sua Tarefa:</h3>
                    </div>
                    <div>
                        <input type="text" name="task" id="task" placeholder="Digite sua tarefa aqui..." v-model="newTask">
                    </div>
                    <div class="div-button">
                        <button @click="addTask">Adicionar</button>
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

button {
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

button:hover {
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