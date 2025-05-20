<script setup>

import { useRouter } from 'vue-router';
import { ref } from 'vue';

import IconWrapper from '../components/icons/IconWrapper.vue';
import TableEscolha from '../components/TableEscolha.vue';
import Modal from '../components/Modal.vue';

const isModalOpen = ref(false);
const isTitleModalOpen = ref(false);
const selectedView = ref('');
const titleInput = ref('');

const router = useRouter();

function handleEscolha(view) {
    if (view === 'checklist') 
    {
        selectedView.value = 'checklist';
        isModalOpen.value = false;
        isTitleModalOpen.value = true;
    } 
    else 
    {
        router.push('/note');
    }
}

function confirmTitleAndNavigate() {
    const title = titleInput.value.trim();

    if (title) {
        router.push({ path: '/checklist', query: { title }})
        isTitleModalOpen.value = false;
        titleInput.value = ''
    }
    else {
        alert('Digite um título para continuar.')
    }
}

</script>

<template> 

    <div class="container">
        <header>
            <h1 id="title">To-Do List</h1>
        </header>

        <main>
            <IconWrapper
                id="icon-add"
                @click="isModalOpen = true"
            />

            <Modal :isOpen="isModalOpen" @close="isModalOpen = false">
                <TableEscolha @escolha="handleEscolha"/>
            </Modal>

            <Modal :isOpen="isTitleModalOpen" @close="isModalOpen = false">
                <div class="modal-content">
                    <h3>Digite o título da checklist:</h3>
                    <input v-model="titleInput" placeholder="Título..."/>
                    <button @click="confirmTitleAndNavigate">Confirmar</button>
                </div>
            </Modal>
        </main>
    </div>

</template>


<style scoped>

header {
    width: 100vw;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 25px;
}

#title {
    text-align: center;
    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    font-size: 35px;
}

#icon-add {
    position: fixed;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    border-radius: 50%;
    padding: 0;
    cursor: pointer;
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
