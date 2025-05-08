<script setup>

const props = defineProps({
    isOpen: Boolean
})

const emit = defineEmits(['close'])

function close() {
    emit('close')
}

</script>

<template> 
  <div v-if="isOpen" class="modal-overlay" @click.self="close">
    <div class="modal-content">
        <slot></slot>

    </div>
  </div>
</template>

<style scoped>

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5); /* Preto semi-transparente */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000; /* Garante que fique acima de tudo */
}

/* Conteúdo do modal */
.modal-content {
  padding: 50px;
  border-radius: 10px;
  max-width: 80%; /* Largura máxima */
  max-height: 80vh; /* Altura máxima */
  overflow-y: auto; /* Rolagem se o conteúdo for muito grande */
  position: relative;
}

</style>
