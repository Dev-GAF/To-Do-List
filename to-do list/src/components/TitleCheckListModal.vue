<script setup>

import { ref } from 'vue';
import Modal from './Modal.vue';

const props = defineProps({
    isOpen: Boolean
});

const emit = defineEmits(['close', 'confirm']);

const titleInput = ref('');

function confirmTitle() {
    const title = titleInput.value.trim();

    if (title) 
    {
        emit('confirm', title);
        titleInput.value = '';
    }
    else 
    {
        alert('Digite um título para continuar.');
    }
}

function handleClose() {
    emit('close');
}

</script>

<template> 

    <Modal :isOpen="isOpen" @close="handleClose">
        <div class="modal-content">
            <h3>Digite o título da checklist:</h3>
            <input v-model="titleInput" placeholder="Título..."/>
            <button @click="confirmTitle">Confirmar</button>
        </div>
    </Modal>

</template>


<style scoped>

h3 {
    font-size: 25px;
}

.modal-content {
    display: flex;
    flex-direction: column;
    gap: 20px;
    background: #1c1d1d;
    padding: 40px;
    border-radius: 20px;
    color: white;
}

.modal-content input {
    padding: 10px;
    font-size: 16px;
}

.modal-content button {
    padding: 10px;
    font-size: 16px;
    cursor: pointer;
}

</style>
