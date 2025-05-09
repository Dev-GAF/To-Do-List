<script setup>

import IconWrapper from "../components/icons/IconWrapper.vue";

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
            />
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
    background: none;
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

</style>