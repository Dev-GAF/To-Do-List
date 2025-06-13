<script setup>

const props = defineProps({
    card: Object
});

const emit = defineEmits(['abrirCard']);

function abrir() {
    emit('abrirCard', props.card);
}

</script>

<template> 
    <div class="card" @click="abrir">
        <div class="div-titulo">
            <h2>{{ card.titulo }}</h2>
        </div>
    </div>
</template>

<style scoped>

.card {
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #13ddbb;
    height: 150px;
    width: 150px;
    margin: 10px;
    border-radius: 10px;
    cursor: pointer;
    color: white;
    transition: background 0.3s;
    padding: 10px; 
    text-align: center; 
}

.card:hover {
    background-color: #4df5d6;
}

h2 {
    margin: 0;
    font-size: 1.3rem;
    word-wrap: break-word; 
}

.div-titulo {
    width: 100%;
}


</style>
