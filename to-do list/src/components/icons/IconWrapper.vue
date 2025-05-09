<script setup>

defineProps({
  tamanho: {
    type: Number,  // Tipo 
    default: 50    // Valor padrão
  },
  corFundo: {
    type: String,  
    default: '#1ec2ac' 
  },
  corIcone: {
    type: String,
    default: 'white' 
  }
});

</script>

<template>
  <div 
    class="icon-wrapper"
    :style="{
      backgroundColor: corFundo,
      color: corIcone,
      width: `${tamanho}px`,
      height: `${tamanho}px`,
    }"
  >
    <svg 
      :width="tamanho * 0.6" 
      :height="tamanho * 0.6"
      :fill="corIcone"
      viewBox="0 0 24 24"
    >
      <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2"/>
    </svg>
  </div>
</template>

<style scoped>

.icon-wrapper {
  display: flex; 
  justify-content: center; 
  align-items: center;
}

.icon-wrapper:hover {
  cursor: pointer;
  box-shadow: 0 10px 20px rgba(255, 255, 255, 0.4);
  transform: scale(1.05);
}

</style>
