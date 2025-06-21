<script setup>
import Trash from '../components/icons/IconTrash.vue';

const props = defineProps({
    card: Object
});

const emit = defineEmits(['abrirCard']);

function abrir() {
    emit('abrirCard', props.card);
}

function deletar() {
  emit('deletarCard', props.card);
}

</script>

<template> 
    <div class="card-container">
        <div class="card" @click="abrir">
            <div class="div-titulo">
                <h2>{{ card.titulo }}</h2>
            </div>
        </div>

        <Trash
            @click.stop="deletar"
        />
    </div>
    
</template>

<style scoped>

.card-container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.card {
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #13ddbb;
    height: 150px;
    width: 150px;
    margin: 10px;
    border-radius: 16px;
    cursor: pointer;
    color: white;
    transition: all 0.3s ease;
    padding: 10px; 
    text-align: center; 
}

.card:hover {
    background-color: #4df5d6;
    transform: translateY(-4px);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

h2 {
    margin: 0;
    font-size: 1.3rem;
    word-wrap: break-word; 
}

.div-titulo {
    width: 100%;
}

.icon-trash {
    cursor: pointer;
    color: rgb(225, 50, 50);
}

.icon-trash:hover {
    color: rgb(255, 80, 80);
}


</style>
