<script setup>

import IconCheck from '../components/icons/IconCheck.vue';
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
// import { useTarefaStore } from '@/store/tarefaStore';
import api from '@/services/api';

const title = ref('');
const content = ref('');

// const tarefaStore = useTarefaStore();
const router = useRouter();
const route = useRoute();

const id = route.params.id;;

onMounted(async () => {
    if (id) 
    {
        const { data } = await api.get(`/notas/${id}`);
        title.value = data.titulo;
        content.value = data.conteudo;
    }
});

async function salvarNota() 
{
    if (title.value.trim() || content.value.trim()) 
    {
        try 
        {
            if (id) 
            {
                await api.put(`/notas/${id}`, {
                    titulo: title.value,
                    conteudo: content.value
                });
            } 
            else 
            {
                await api.post('/notas', {
                    titulo: title.value,
                    conteudo: content.value
                });
            }

            router.push({ name: 'home' });
        } 
        catch (error) 
        {
            console.error('Erro ao salvar nota', error);
        }
    }
}

</script>


<template>
    <div class="container">
        <input v-model="title" placeholder="Título" class="note-title" />
        <textarea v-model="content" placeholder="Conteúdo" class="note-content"></textarea>
        <div class="footer">
            <IconCheck @click="salvarNota"/>
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
    justify-content: flex-end;
}

.icon-check {
    margin: 7px 15px;
}

</style>