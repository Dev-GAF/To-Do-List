<script setup>
import IconCheck from '../components/icons/IconCheck.vue'
</script>


<template>
    <div class="container">
        <div>
            <input v-model="title" placeholder="Título" class="note-title" />
        </div>
        <div>
            <textarea v-model="content" placeholder="Conteúdo" class="note-content" />
        </div>

        <div class="div-iconcheck">
            <IconCheck />
        </div>
    </div>
</template>


<style scoped>

input,
textarea {
    color: white;
}

input {
    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}

textarea {
    font-family: Arial, Helvetica, sans-serif;
}

input::placeholder,
textarea::placeholder {
  font-family: sans-serif; 
}

.note-title {
    width: 100vw;
    height: 6vh;
    margin-top: 1vh;
    padding: 10px;
    font-size: 30px;
}

.note-content {
    margin-top: 1vh;
    height: 83.2vh;
    width: 100vw;
    font-size: 18px;
    padding: 5px;
}

.note-title, .note-content{
    background: none;
    border: none;
    outline: 0;
}

.div-iconcheck {
    width: 100vw;
    background: rgba(0, 0, 0, 0.541); /* Só afeta a div */
    height: 8vh;
}

</style>