<script setup>

import { useRouter } from 'vue-router'

import { ref } from 'vue'

import IconWrapper from '../components/icons/IconWrapper.vue'
import TableEscolha from '../components/TableEscolha.vue'
import Modal from '../components/Modal.vue'

const isModalOpen = ref(false)
const router = useRouter()

function handleEscolha(view) {
  router.push(view === 'note' ? '/note' : '/checklist')
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
    </main>
  </div>
  
</template>


<style scoped>

header {
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
}

#title {
  text-align: center;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
  font-size: 30px;
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

</style>
