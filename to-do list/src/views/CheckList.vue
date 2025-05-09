<script setup>

import IconWrapper from "../components/icons/IconWrapper.vue";
import Modal from "../components/Modal.vue"

import { ref } from 'vue'

const isModalOpen = ref(false);

function controlModal() {
    isModalOpen.value = !isModalOpen.value;
}

</script>


<template>
    <div class="container">
        <input v-model="title" placeholder="Título" class="note-title" />
        <div id="teste">
            
        </div>
        <div class="footer">
            <icon-wrapper 
                tamanho=55
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
                        <input type="text" name="task" id="task" placeholder="Digite sua tarefa aqui...">
                    </div>
                    <div class="div-button">
                        <button>Adicionar</button>
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
    margin-top: 30px;
}

#task::placeholder {
    color: rgba(255, 255, 255, 0.727);
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

</style>