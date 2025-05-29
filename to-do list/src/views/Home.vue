<script setup>

import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

import api from '@/services/api'

import IconWrapper from '../components/icons/IconWrapper.vue';
import TableEscolha from '../components/TableEscolha.vue';
import Modal from '../components/Modal.vue';
import TitleCheckListModal from '../components/TitleCheckListModal.vue';

import Card from '@/components/Card.vue';

const isModalOpen = ref(false);
const isTitleModalOpen = ref(false);
const selectedView = ref('');

const router = useRouter();
const notas = ref([]);

onMounted(async () => {
    const { data } = await api.get('/notas');
    notas.value = data;
})

function abrirNota(nota) {
    router.push({ name: 'note', params: { id: nota.id } });
}

function handleEscolha(view) {
    if (view === 'checklist') 
    {
        selectedView.value = 'checklist';
        isModalOpen.value = false;
        isTitleModalOpen.value = true;
    } 
    else 
    {
        router.push('note');
    }
}

function confirmTitleAndNavigate(title) {
    router.push({ path: 'checklist', query: { title }})
    isTitleModalOpen.value = false;
}

</script>

<template> 

    <div class="page-wrapper">
        <header>
            <h1 id="title">To-Do List</h1>
        </header>

        <IconWrapper
            id="icon-add"
            @click="isModalOpen = true"
        />

        <Modal :isOpen="isModalOpen" @close="isModalOpen = false">
            <TableEscolha @escolha="handleEscolha"/>
        </Modal>

        <TitleCheckListModal 
            :isOpen="isTitleModalOpen"
            @close="isTitleModalOpen = false"
            @confirm="confirmTitleAndNavigate"
        />

        <Card
            v-for="nota in notas"
            :key="nota.id"
            :card="nota"
            @abrirCard="abrirNota"
        />
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
