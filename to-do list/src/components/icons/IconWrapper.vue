<template>
  <div class="icon-wrapper">
    <svg 
      xmlns="http://www.w3.org/2000/svg" 
      width="50" 
      height="50" 
      viewBox="0 0 24 24" 
      fill="currentColor"
    >
      <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
    </svg>
  </div>
</template>

<style scoped>

.icon-wrapper {
  background-color: #1ec2ac;
  color: white;
  border-radius: 50%;
  padding: 12px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 8px rgba(0,0,0,0.2);
  cursor: pointer;
  transition: box-shadow 0.3s ease, transform 0.3s ease;
}

.icon-wrapper:hover {
  box-shadow: 0 10px 20px rgba(255, 255, 255, 0.4);
  transform: scale(1.05);
}

</style>
